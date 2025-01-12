#!/usr/bin/env bash

set -e

PROJECT_DIR="$(realpath "$(cd "$(dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd)/../..")"

# shellcheck disable=SC1090 # (dont follow, will be checked separately)
source "$PROJECT_DIR/lib/scripts/script-utils.sh"

# build_release_json_payload builds a payload that can be used to create a GitHub release via the API.
# arguments: [Version] [Changelog description] [Project URL]
# Line lines are stripped from the changelog, otherwise the GitHub release will be malformatted.
build_release_json_payload() {
  local version="$1"
  local changelog_description="$2"
  local project_url="$3"

  verify_has_value "$version" "Aborting, version has not been supplied to ${FUNCNAME[0]}"
  verify_has_value "$changelog_description" "Aborting, changelog description has not been supplied to ${FUNCNAME[0]}"
  verify_has_value "$project_url" "Aborting, project URL has not been supplied to ${FUNCNAME[0]}"

  local description=""
  description="$description ##### Changes\n"
  description="$description $changelog_description"

  local unsafe_release_data="{\"tag_name\":\"$version\",\"description\":\"$description\"}"

  # use node to help replace new lines with \n
  local release_data
  local json_type
  release_data=$(echo "$unsafe_release_data" |
    node -e "console.log((require('fs')).readFileSync(process.stdin.fd, 'utf-8').trim().replace(/\n/g, '\\\n'));")
  json_type=$(echo "$release_data" | jq type | sed "s/\"//g")

  if [[ "$json_type" != "object" ]]; then
    error "Aborting, extracted release data type '$json_type' is not a JSON object. Release data: $release_data"
  fi

  local extracted_tag_name
  local extracted_description
  extracted_tag_name=$(echo "$release_data" | jq ".tag_name" | sed "s/\"//g")
  extracted_description=$(echo "$release_data" | jq ".description" | sed "s/\"//g")

  verify_has_value "$extracted_tag_name" "Aborting, unable to determine the tag name from the release data $release_data"
  verify_has_value "$extracted_description" "Aborting, unable to determine the description from the release data $release_data"

  echo "$release_data"
}

# verify_version_not_released ensures that there is not already a release for the version attempting to be released.
# arguments: [GitHub API token] [CI project ID] [Version]
verify_version_not_released() {
  local github_token="$1"
  local project_id="$2"
  local version="$3"

  verify_has_value "$github_token" "Aborting, GitHub CI Token has not been supplied to ${FUNCNAME[0]}"
  verify_has_value "$project_id" "Aborting, CI Project ID has not been supplied to ${FUNCNAME[0]}"
  verify_has_value "$version" "Aborting, version has not been supplied to ${FUNCNAME[0]}"

  local version_url="https://github.com/api/v4/projects/$project_id/repository/tags/$version"

  if curl --silent --fail --show-error --header "private-token:$github_token" "$version_url"; then
    error "Aborting, tag $version already exists. If this is not expected, please remove the tag and try again."
  fi
}

# tag_git_commit uses the GitHub API to create a lightweight Git tag.
# arguments: [GitHub API token] [CI Project ID] [Tag name] [Commit SHA]
tag_git_commit() {
  local github_token="$1"
  local project_id="$2"
  local tag_name="$3"
  local commit_sha="$4"

  verify_has_value "$github_token" "Aborting, GitHub CI Token has not been supplied to ${FUNCNAME[0]}"
  verify_has_value "$project_id" "Aborting, CI Project ID has not been supplied to ${FUNCNAME[0]}"
  verify_has_value "$tag_name" "Aborting, tag name has not been supplied to ${FUNCNAME[0]}"
  verify_has_value "$commit_sha" "Aborting, commit SHA has not been supplied to ${FUNCNAME[0]}"

  local tag_url="https://github.com/api/v4/projects/$project_id/repository/tags?tag_name=$tag_name&ref=$commit_sha"

  curl --silent --fail --show-error --request POST --header "PRIVATE-TOKEN:$github_token" "$tag_url"
}

# create_github_release uses the GitHub API to create a GitHub release.
# The Git tag should be created before this function is run.
# arguments: [GitHub API token] [CI Project ID] [Release payload]
create_github_release() {
  local github_token="$1"
  local project_id="$2"
  local payload="$3"

  verify_has_value "$github_token" "Aborting, GitHub CI Token has not been supplied to ${FUNCNAME[0]}"
  verify_has_value "$project_id" "Aborting, CI Project ID has not been supplied to ${FUNCNAME[0]}"
  verify_has_value "$payload" "Aborting, release payload has not been supplied to ${FUNCNAME[0]}"

  local release_url="https://github.com/api/v4/projects/$project_id/releases"

  curl --silent --fail --show-error --request POST \
    --header "PRIVATE-TOKEN:$github_token" \
    --header 'Content-Type:application/json' \
    --data "$payload" \
    "$release_url"
}

tag_and_release(){
  local github_token="$1"
  local project_id="$2"
  local tag_name="$3"
  local commit_sha="$4"
  local release_data="$5"

  # Creates a git tag for the release commit
  # tag has follows semver: vx.y.z
  log "Tagging Git SHA $commit_sha with $tag_name"
  tag_git_commit "$github_token" "$project_id" "$tag_name" "$commit_sha"

  # Creates the actual release
  log "Creating GitHub release from Git tag $tag_name"
  create_github_release "$github_token" "$project_id" "$release_data"
}