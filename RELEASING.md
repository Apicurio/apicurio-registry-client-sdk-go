# Release Guide

To create a release on [pkg.go.dev](https://pkg.go.dev/) simply create a new tag!

## Create release

1. `git tag 1.0.0`.
2. `git push <remote> 1.0.0`.
3. The [Release](https://github.com/Apicurio/apicurio-registry-client-sdk-go/actions/workflows/release.yml) workflow will create the GitHub release.
4. Once completed, the package should be available on [pkg.go.dev/github.com/Apicurio/apicurio-registry-client-sdk-go](https://pkg.go.dev/github.com/Apicurio/apicurio-registry-client-sdk-go).
