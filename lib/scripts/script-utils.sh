#!/usr/bin/env bash

set -e

# error prints an error message and exits. builds a payload that can be used to create a GitHub release via the API.
# arguments: [Message...]
error() {
  printf "\n"
  printf "%s" "$*" >>/dev/stderr
  printf "\n"
  exit 1
}

# log prints a message in a heading format.
# arguments: [Message...]
log() {
  printf "\n\n######### %s #########\n" "$*" >>/dev/stdout
}

# verify_has_value errors with given message if the variable is not set.
# arguments: [Variable] [Message]
verify_has_value() {
  variable="$1"
  message="$2"

  if [[ -z $variable ]]; then
    error "$message"
  fi
}

# error_if_has_value errors with value of the first argument if the first argument has a non-empty value.
# arguments: [Variable]
error_if_has_value() {
  variable="$1"

  if [[ -n $variable ]]; then
    error "$variable"
  fi
}
