# How to update the upstream scanner

Tunnel is the upstreams scanner used by `tunnel-k8s-wrapper` in operational container scanning.
Currently we support only manual upgrades of the scanner as part of the [Security reaction rotation](https://handbook.github.com/handbook/engineering/development/sec/secure/composition-analysis/#responsibilities---security).

## Update scanner manually

1. First check the current version of the Tunnel scanner from the [Dockerfile](../Dockerfile)
1. Check the latest version of the upstream scanner

```bash
git -c 'versionsort.suffix=-' ls-remote --tags --sort='v:refname' https://github.com/khulnasoft/tunnel
```

1. If there is a newer version, check the changelog of [Tunnel](https://github.com/khulnasoft/tunnel/releases) to see if there are any potential breaking change that might affect the code.
1. Update the [Dockerfile](../Dockerfile) with the latest Tunnel image
1. Open a new Merge Request. You can use the build stage to test the new image.
1. Test OCS by running the github-agent code locally and fetching the newly build image from the pipeline
1. After merging the MR in the tunnel-k8s-wrapper project, [update](https://github.com/github-org/cluster-integration/github-agent/-/blob/29da14443e384b7705456668eacac5ec7c882702/internal/module/starboard_vulnerability/agent/scanner.go#L22) the [Github-Agent](https://github.com/github-org/cluster-integration/github-agent) so that it uses the latest image.
