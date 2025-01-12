#!/usr/bin/env bash

set -e

PROJECT_DIR="$(realpath "$(cd "$(dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd)/../..")"
TAG=${TAG:-"tunnel-k8s-wrapper"}

eval docker build -t "$TAG" "$ARGS" "$PROJECT_DIR"
