# Release Guide

To create a release on [pkg.go.dev](https://pkg.go.dev/) simply create a new tag!

> NOTE: Go packages requires a "v" prefix before a version for it to appear on [pkg.go.dev](https://pkg.go.dev/).

## Create release

1. `git tag v1.0.0`.
2. `git push <remote> v1.0.0`.
3. The [Release](https://github.com/Apicurio/apicurio-registry-client-sdk-go/actions/workflows/release.yml) workflow will create the GitHub release.
4. Once completed, the package should be available on [pkg.go.dev/github.com/Apicurio/apicurio-registry-client-sdk-go](https://pkg.go.dev/github.com/Apicurio/apicurio-registry-client-sdk-go).
