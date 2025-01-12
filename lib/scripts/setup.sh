#!/bin/sh

set -eu

if [ "$TARGETOS" != "linux" ]; then
	echo "Unsupported operating system: $TARGETOS"
	exit 1
fi

get_archive_name() {
	TUNNEL_VERSION="$1"

	case "$TARGETARCH" in
	"arm64")
		printf "%s" "tunnel_${TUNNEL_VERSION}_Linux-ARM64.tar.gz"
		;;
	"amd64")
		printf "%s" "tunnel_${TUNNEL_VERSION}_Linux-64bit.tar.gz"
		;;
	*)
		echo "Unsupported architecture: $TARGETARCH"
		exit 1
		;;
	esac
}

TUNNEL_VERSION="$(cat TUNNEL_VERSION)"
TUNNEL_ARCHIVE_NAME="$(get_archive_name "${TUNNEL_VERSION}")"
TUNNEL_ARCHIVE_LOC="https://github.com/khulnasoft/tunnel/releases/download/v${TUNNEL_VERSION}/${TUNNEL_ARCHIVE_NAME}"
TUNNEL_CHECKSUMS_NAME="tunnel_${TUNNEL_VERSION}_checksums.txt"
TUNNEL_CHECKSUMS_LOC="https://github.com/khulnasoft/tunnel/releases/download/v${TUNNEL_VERSION}/${TUNNEL_CHECKSUMS_NAME}"

echo "Creating temp directory"
mkdir -p /home/gitlab/opt/tunnel

echo "Downloading checksums from ${TUNNEL_CHECKSUMS_LOC}"
wget "$TUNNEL_CHECKSUMS_LOC"

echo "Downloading binary from ${TUNNEL_ARCHIVE_LOC}"
wget "$TUNNEL_ARCHIVE_LOC"

grep "$TUNNEL_ARCHIVE_NAME" "$TUNNEL_CHECKSUMS_NAME" | sha256sum -c -

echo "Installing Tunnel ${TUNNEL_VERSION}"
tar -zxvf "${TUNNEL_ARCHIVE_NAME}" -C /home/gitlab/opt/tunnel
