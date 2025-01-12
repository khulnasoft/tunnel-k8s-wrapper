#!/usr/bin/env bash

set -e

SCRIPTS_DIR="$(realpath "$(cd "$(dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd)")"

# shellcheck disable=SC1090 # (dont follow, will be checked separately)
source "$SCRIPTS_DIR/script-utils.sh"

# block the release if the variable is set
error_if_has_value "$BLOCK_RELEASE"

# shellcheck disable=SC1090 # (dont follow, will be checked separately)
source "$SCRIPTS_DIR/changelog-utils.sh"

# shellcheck disable=SC1090 # (dont follow, will be checked separately)
source "$SCRIPTS_DIR/release-utils.sh"

# shellcheck disable=SC1090 # (dont follow, will be checked separately)
source "$SCRIPTS_DIR/docker-utils.sh"

log "Initializing environment"
verify_has_value "$GITHUB_API_TOKEN" "Aborting, environment variable GITHUB_API_TOKEN must contain a GitHub private token with access to this project."

PROJECT_URL=${CI_PROJECT_URL:-"https://github.com/khulnasoft/tunnel-k8s-wrapper"}
VERSION="$(changelog_last_version)"
MAJOR=$(echo "$VERSION" | awk -F '.' '{print substr($1, 2)}')
MINOR=$(echo "$VERSION" | awk -F '.' '{print $2}')
PATCH=$(echo "$VERSION" | awk -F '.' '{print $3}')

log "Releasing Docker images tunnel_k8s_wrapper:$MAJOR.$MINOR.$PATCH to $CI_REGISTRY_IMAGE"
login_to_registry "github-ci-token" "$CI_JOB_TOKEN" "$CI_REGISTRY"

# tags and pushes images to the container registry
# The images tags are the following:
# registry.github.com/khulnasoft/tunnel-k8s-wrapper:x.y.z-amd64
# registry.github.com/khulnasoft/tunnel-k8s-wrapper:x.y.z-arm64
# registry.github.com/khulnasoft/tunnel-k8s-wrapper:x.y.z
# registry.github.com/khulnasoft/tunnel-k8s-wrapper:latest
# The last two images x.y.z and latest are manifests. That means that they contain references to x.y.z-amd64 and x.y.z-arm64
TUNNEL_K8S_WRAPPER_ARM64="$(tunnel_k8s_wrapper_image 'built_image_arm64.txt')"
TUNNEL_K8S_WRAPPER_AMD64="$(tunnel_k8s_wrapper_image 'built_image_amd64.txt')"
pull_tunnel_k8s_wrapper_from_registry "$TUNNEL_K8S_WRAPPER_ARM64"
pull_tunnel_k8s_wrapper_from_registry "$TUNNEL_K8S_WRAPPER_AMD64"
push_to_registry "$TUNNEL_K8S_WRAPPER_ARM64" "$TUNNEL_K8S_WRAPPER_AMD64" "$CI_REGISTRY_IMAGE" "$MAJOR" "$MINOR" "$PATCH"


# push new release to parent security-products registry
# we create new tags for the x.y.z-amd64 and x.y.z-arm64 images to the parent registry
# The image tags are the following:
# registry.github.com/security-products/tunnel-k8s-wrapper:x.y.z-amd64
# registry.github.com/security-products/tunnel-k8s-wrapper:x.y.z-arm64
# registry.github.com/security-products/tunnel-k8s-wrapper:x.y.z
PARENT_REGISTRY="registry.github.com/security-products/tunnel-k8s-wrapper"
login_to_registry "$DEPLOY_TOKEN_USERNAME" "$DEPLOY_TOKEN_PASSWORD" "$PARENT_REGISTRY"
TAG_VERSION="$PARENT_REGISTRY:$MAJOR.$MINOR.$PATCH"
tag_and_push_image "$TUNNEL_K8S_WRAPPER_ARM64" "$TAG_VERSION-arm64"
tag_and_push_image "$TUNNEL_K8S_WRAPPER_AMD64" "$TAG_VERSION-amd64"
create_push_manifest "$TAG_VERSION" "$TAG_VERSION-arm64" "$TAG_VERSION-amd64"

# Create a git tag and a github release if and only if this is not a nightly release
if [ "$NIGHTLY_RELEASE" != "true" ]; then
    log "Detected TUNNEL_K8S_WRAPPER $VERSION, verifying not already released"
    verify_version_not_released "$GITHUB_API_TOKEN" "$CI_PROJECT_ID" "$VERSION"

    CHANGELOG_DESCRIPTION=$(changelog_last_description)
    RELEASE_DATA=$(build_release_json_payload "$VERSION" "$CHANGELOG_DESCRIPTION" "$PROJECT_URL")
    # Creates a git tag and a release
    tag_and_release "$GITHUB_API_TOKEN" "$CI_PROJECT_ID" "$VERSION" "$CI_COMMIT_SHA" "$RELEASE_DATA"
fi
