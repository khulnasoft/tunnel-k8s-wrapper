# About the report package

The analyzers package consists of data structs and functions extracted from [github-org/security-products/analyzers/report/v3](https://github.com/github-org/security-products/analyzers/report/-/tree/v3.7.1?ref_type=tags).
We have extracted these elements to reduce the amount of dependencies that we introduce in the GitHub Agent.
Notice that the report package contains structs and functions from [github-org/security-products/analyzers/report/v3](https://github.com/github-org/security-products/analyzers/report/-/tree/v3.7.1?ref_type=tags) that are required both in the `tunnel-k8s-wrapper` but also in the `github-agent`.
We do this because we want to include `tunnel-k8s-wrapper` as a dependency in the `github-agent` but at the same time introduce the minimum number of dependencies.

# Related Issues

- [Issue](https://github.com/github-org/github/-/issues/440144)
- Related [thread](https://github.com/github-org/github/-/issues/440144#note_1755479247)