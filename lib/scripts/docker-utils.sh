#!/usr/bin/env bash

set -e

PROJECT_DIR="$(realpath "$(cd "$(dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd)/../..")"

# shellcheck disable=SC1090 # (dont follow, will be checked separately)
source "$PROJECT_DIR/lib/scripts/script-utils.sh"

# path_to_tunnel_k8s_wrapper_image_name checks presence of and returns the path to the file containing the built tunnel_k8s_wrapper image
path_to_tunnel_k8s_wrapper_image_name() {
  local image_file_name="$1"
  verify_has_value "$image_file_name" "Aborting, image file name has not been supplied to ${FUNCNAME[0]}"

  local image_file
  image_file=$(realpath "$PROJECT_DIR/$image_file_name")

  if ! [[ -f $image_file ]]; then
    error "Aborting, unable to determine previously built tunnel_k8s_wrapper image as file containing name doesn't exist: '$image_file'"
  fi

  echo "$image_file"
}

# tunnel_k8s_wrapper_image returns the image name of the built tunnel_k8s_wrapper image
tunnel_k8s_wrapper_image() {
  local image_file_name="$1"
  verify_has_value "$image_file_name" "Aborting, image file name has not been supplied to ${FUNCNAME[0]}"

  cat "$(path_to_tunnel_k8s_wrapper_image_name "$image_file_name")"
}

# tag_and_push_image tags a docker image
# arguments: [from] [to]
# The from image must be available locally
# You must be logged into the registry hosting the to image
tag_and_push_image() {
  local from="$1"
  local to="$2"

  verify_has_value "$from" "Aborting, unable to tag image as the from location has not been supplied to ${FUNCNAME[0]}"
  verify_has_value "$to" "Aborting, unable to tag image as the to location has not been supplied to ${FUNCNAME[0]}"

  docker tag "$from" "$to"
  docker push "$to"
}

# create_push_manifest creates a manifest for
# binding two different arch images in one tag
# arguments: [tag] [arm64] [amd64]
# Tag is the new tag that you want to create for the manifest
# arm64 is the linux/arm64 image to ammend in the manifest
# amd64 is the linux/amd64 image to ammend in the manifest
create_push_manifest() {
  local tag="$1"
  local arm64="$2"
  local amd64="$3"

  verify_has_value "$tag" "Aborting, tag has not been supplied to ${FUNCNAME[0]}"
  verify_has_value "$arm64" "Aborting, arm64 has not been supplied to ${FUNCNAME[0]}"
  verify_has_value "$amd64" "Aborting, amd64 has not been supplied to ${FUNCNAME[0]}"

  docker manifest create ${tag} --amend ${arm64} --amend ${amd64}
  docker manifest push ${tag}
}

# push_to_registry pushes the Docker image to a remote registry.
# Applies multiple tags to the same image.
# arguments: [tunnel-k8s-wrapper-arm64] [tunnel-k8s-wrapper-amd64] [CI Registry Image]  [Major Version] [Minor Version] [Patch Version]
push_to_registry() {
  local arm64="$1"
  local amd64="$2"
  local ci_registry_image="$3"
  local major_version="$4"
  local minor_version="$5"
  local patch_version="$6"

  verify_has_value "$arm64" "Aborting, tunnel-k8s-wrapper-arm64 image has not been supplied to ${FUNCNAME[0]}"
  verify_has_value "$amd64" "Aborting, tunnel-k8s-wrapper-amd64 image has not been supplied to ${FUNCNAME[0]}"
  verify_has_value "$ci_registry_image" "Aborting, ci registry image has not been supplied to ${FUNCNAME[0]}"
  verify_has_value "$major_version" "Aborting, major version has not been supplied to ${FUNCNAME[0]}"
  verify_has_value "$minor_version" "Aborting, minor version has not been supplied to ${FUNCNAME[0]}"
  verify_has_value "$patch_version" "Aborting, patch version has not been supplied to ${FUNCNAME[0]}"

  tag_version="$ci_registry_image:$major_version.$minor_version.$patch_version"
  tag_and_push_image "$arm64" "$tag_version-arm64"
  tag_and_push_image "$amd64" "$tag_version-amd64"
  create_push_manifest "$tag_version" "$tag_version-arm64" "$tag_version-amd64"
  tag_and_push_image "$arm64" "$ci_registry_image:latest"
  tag_and_push_image "$amd64" "$ci_registry_image:latest"
  create_push_manifest "$ci_registry_image:latest" "$tag_version-arm64" "$tag_version-amd64"
}

# pulls an image from the registry logs in to a remote registry
pull_tunnel_k8s_wrapper_from_registry() {
  local tunnel_k8s_wrapper="$1"

  verify_has_value "$tunnel_k8s_wrapper" "Aborting, tunnel_k8s_wrapper image has not been supplied to ${FUNCNAME[0]}"
  docker pull "$tunnel_k8s_wrapper"
}

# login_to_registry logs in to a remote registry
# arguments: [Registry Username] [Registry Password] [CI Registry]
login_to_registry() {
  local registry_username="$1"
  local registry_password="$2"
  local ci_registry="$3"

  verify_has_value "$registry_username" "Aborting, ci registry username has not been supplied to ${FUNCNAME[0]}"
  verify_has_value "$registry_password" "Aborting, ci registry password has not been supplied to ${FUNCNAME[0]}"
  verify_has_value "$ci_registry" "Aborting, ci registry has not been supplied to ${FUNCNAME[0]}"

  docker login -u "$registry_username" -p "$registry_password" "$ci_registry"
}
