# CDK for Terraform Go Example

[Cloud Development Kit for Terraform](https://www.terraform.io/cdktf) (CDKTF) allows you to use familiar programming languages to define and provision infrastructure. This is a simple example of how to use CDKTF to create a [Kubernetes](https://kubernetes.io/) cluster on the IONOS Cloud using the [Go](https://go.dev/) programming language.

:warning: What is deployed is a [Managed Kubernetes Cluster](https://cloud.ionos.de/managed/kubernetes) on the IONOS Cloud :cloud:. Please be aware that this is not for free.

## Prerequisites

* [IONOS Account](https://cloud.ionos.com/)
* [Terraform](https://www.terraform.io/)
* [CDK for Terrform](https://learn.hashicorp.com/tutorials/terraform/cdktf-install?in=terraform/cdktf#install-cdktf)

Optional:

* [`ionosctl`](https://github.com/ionos-cloud/ionosctl)

## Deploy

You either have to set the `IONOS_USERNAME` and `IONOS_PASSWORD` or `IONOS_TOKEN` environment variables.

> `ionosctl token create` and `ionosctl token get --token-id` gives you a token.

```bash
cdk deploy
```

## Tests

Writing tests for CDKTF is easy with Go.

```bash
go test ./...
```

If you are finished experimenting you can do a simple `cdk destroy` to clean up.
