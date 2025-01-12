#!/usr/bin/env bash

set -e

PROJECT_DIR="$(realpath "$(cd "$(dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd)/../..")"
VERSION_PATTERN='^## \(v[0-9]*\.[0-9]*\.[0-9]*\)$'

# shellcheck disable=SC1090 # (dont follow, will be checked separately)
source "$PROJECT_DIR/lib/scripts/script-utils.sh"

# path_to_changelog checks presence of and returns the path to the CHANGELOG
path_to_changelog() {
  local changelog
  changelog=$(realpath "$PROJECT_DIR/CHANGELOG.md")

  verify_has_value "$changelog" "Unable to find CHANGELOG file with path '$changelog'."

  echo "$changelog"
}

# changelog_last_version checks and returns the version
# of the most recent changelog entry (first to appear in the file).
# It fails if the version is not Semver-compliant or is a pre-release.
changelog_last_version() {
  # find the first matching version, e.g. ## v1.6.0
  local version
  version=$(sed -n "s/$VERSION_PATTERN/\\1/p" "$(path_to_changelog)" | sed -n '1,1p;1q')

  verify_has_value "$version" "Aborting, unable to determine the latest version in the changelog file." \
    "Expected line with version to have format like '## v1.3.6'."

  # find the first line that starts with ## (should also be the first matching version)
  local most_recent_version
  most_recent_version=$(grep -m 1 '^##.*$' "$(path_to_changelog)" | sed 's/## //')

  if [[ "$most_recent_version" != "$version" ]]; then
    error "The most recent version in the changelog $most_recent_version does not conform to the expected format." \
      "Expected version line to have format like '## v1.3.6'."
  fi

  echo "$version"
}

# changelog_last_description checks and returns the description
# of the most recent changelog entry (first to appear in the file).
# It fails if the version is not Semver-compliant or is a pre-release.
changelog_last_description() {
  # extract the latest version description from the CHANGELOG
  local changelog_description_start
  local changelog_description_end
  local changelog_description
  changelog_description_start=$(sed -n "/$VERSION_PATTERN/=" "$(path_to_changelog)" | sed -n '1,1p;1q' | awk '{print $0 + 1}')
  changelog_description_end=$(sed -n "/$VERSION_PATTERN/=" "$(path_to_changelog)" | sed -n '2,2p;2q' | awk '{print $0 - 2}')

  if [[ "$changelog_description_end" == "" ]]; then
    changelog_description_end=$(wc -l "$(path_to_changelog)" | awk '{print $1}')
  fi

  changelog_description=$(sed -n "${changelog_description_start},${changelog_description_end}p;${changelog_description_end}q" "$(path_to_changelog)")

  echo "$changelog_description"
}
