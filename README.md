# CDK for Terraform Go Example

[![Test & Build](https://github.com/ionos-cloud/event-gateway/actions/workflows/main.yml/badge.svg)](https://github.com/ionos-cloud/event-gateway/actions/workflows/main.yml)
[![Taylor Swift](https://img.shields.io/badge/secured%20by-taylor%20swift-brightgreen.svg)](https://twitter.com/SwiftOnSecurity)
[![Volkswagen](https://auchenberg.github.io/volkswagen/volkswargen_ci.svg?v=1)](https://github.com/auchenberg/volkswagen)

[Cloud Development Kit for Terraform](https://www.terraform.io/cdktf) (CDKTF) allows you to use familiar programming languages to define and provision infrastructure. This is a simple example of how to use CDKTF to create a [Kubernetes](https://kubernetes.io/) cluster on the IONOS Cloud using the [Go](https://go.dev/) programming language.

:warning: What is deployed is a [Managed Kubernetes Cluster](https://cloud.ionos.de/managed/kubernetes) on the IONOS Cloud :cloud:. Please be aware that this is not for free.

## Prerequisites

* [IONOS Account](https://cloud.ionos.com/)
* [Terraform](https://www.terraform.io/)
* [CDK for Terrform](https://learn.hashicorp.com/tutorials/terraform/cdktf-install?in=terraform/cdktf#install-cdktf)
* [Object Storage Bucket](https://cloud.ionos.com/storage/object-storage)

Optional:

* [`ionosctl`](https://github.com/ionos-cloud/ionosctl)

## Deploy

You either have to set the `IONOS_USERNAME` and `IONOS_PASSWORD` or `IONOS_TOKEN` environment variables.

> `ionosctl token create` and `ionosctl token get --token-id` gives you a token.

Furtheremore, you have to set `ACCESS_KEY`, `SECRET_KEY` and `BUCKET_NAME` environment variables for the Terraform backend.

> This can be used in combination with a GitHub Action to deploy the cluster.

> The bucket should have versioning enabled.

```bash
cdktf deploy
```

## Tests

Writing tests for CDKTF is easy with Go.

```bash
go test ./...
```

If you are finished experimenting you can do a simple `cdktf destroy` to clean up.

## License

[MIT](/LICENSE)
