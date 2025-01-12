# Tunnel K8S wrapper changelog

## v0.8.2
- Bump Tunnel from 0.56.2 to 0.58.1 (!61)

## v0.8.1
- Bump go version from 1.21 to 1.23 (!60)

## v0.8.0
- Use GitHub tunnel java db mirror as fallback if default db mirror fails (!53)

## v0.7.0
- Remove usage of explicit CLI flags in Dockerfile entrypoint
- Rename `WORKLOAD_KINDS` environment variable to `WORKLOADS`

## v0.6.2
- Bump Tunnel image from 0.54.0 to 0.56.2 (!55,!57)

## v0.6.1
- Upgrade protobuf from v1.31.0 to v1.33.0 (!56)

## v0.6.0
- Introduce new CLI flag (`--report-max-size`) to configure the maximum size of a Tunnel Report

## v0.5.0
- Introduce new CLI flag (`--timeout`) to replace existing `--timeout-minutes` flag. Flag is responsible for setting Tunnel Timeout as well as Job Timeout. (!50)

## v0.4.0
- Define image for language vulnerabilities (!49)

## v0.3.3
- Update Tunnel to v0.54.0 (!44)

## v0.3.2
- Return correct Identifier URL (!40)

## v0.3.1
- Update Tunnel to v0.52.2 (!33)
- Disable node-collector which is enabled by default (!33)

## v0.3.0
- Add alpine as a base image (!30)

## v0.2.15
- Dockerfile is using a non root user (!21)

## v0.2.14
- Fix configmap namespaces to support any github agent namespace (!24)

## v0.2.13
- Add linux/arm64 image to the release (!23)

## v0.2.12
- Bump Tunnel image from 0.48.0 to 0.49.1 (!20)

## v0.2.11
- Bump Golang from 1.20 to 1.21 (!19)

## v0.2.10
- Replace logrus with uber (!18)

## v0.2.9
- Vendor analyzers report dependency (!16)

## v0.2.8
- Update packages (!17)

## v0.2.7
- Fix prototool package path (!15)

## v0.2.6
- Add agent ID in a label to all Configmaps (!14)

## v0.2.5
- Updates the module name to the new domain (!13)

## v0.2.4
- Introduce license, contributing and notice files (!9)

## v0.2.3
- Release image to registry.github.com/security-products/tunnel-k8s-wrapper  (!12)

## v0.2.2
- Build one image per branch and not one per commit (!10)

## v0.2.1
- Export exit codes constants and use Tunnel image v0.48.0 (!7)

## v0.2.0
- Store vulnerabilities in chained configmaps (!6)

## v0.1.1
- Introduce unit tests (!2)

## v0.1.0
- Introduce Tunnel K8S wrapper code (!1)
