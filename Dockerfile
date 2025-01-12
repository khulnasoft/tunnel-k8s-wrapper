# syntax=docker/dockerfile:1

FROM golang:1.23-bookworm AS builder

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

# Build the binary.
RUN go build -v -o tunnelK8Swrapper cmd/tunnel/main.go

FROM alpine:3.19 AS tunnel-download

ARG TARGETOS
ARG TARGETARCH

COPY ./tunnel/versions/TUNNEL_VERSION .
COPY ./lib/scripts/setup.sh .

RUN ./setup.sh

# This version supports scanning of private images
FROM alpine:3.19

RUN <<-EOF
  # gcompat contains libresolv.so.2 which is required by the go binary.
  apk update && apk upgrade && apk --no-cache add gcompat wget ca-certificates
  # create github user and give write access to app dir
  addgroup --gid 1001 github
  adduser -S github -G github
EOF

USER github
WORKDIR /home/github
ENV PATH /home/github:${PATH}

COPY --from=builder --chown=github:github --chmod=700 /app/tunnelK8Swrapper /app/tunnelK8Swrapper
COPY --from=tunnel-download --chown=github:github --chmod=700 /home/github/opt/tunnel/tunnel /home/github/tunnel

CMD ["/app/tunnelK8Swrapper", "scan"]
ENTRYPOINT [""]
