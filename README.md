# mks-go: Go SDK for Managed Kubernetes Service
[![Go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/selectel/mks-go/)
[![Go Report Card](https://goreportcard.com/badge/github.com/selectel/mks-go)](https://goreportcard.com/report/github.com/selectel/mks-go)
[![Build Status](https://travis-ci.org/selectel/mks-go.svg?branch=master)](https://travis-ci.org/selectel/mks-go)
[![Coverage Status](https://coveralls.io/repos/github/selectel/mks-go/badge.svg?branch=master)](https://coveralls.io/github/selectel/mks-go?branch=master)

Package mks-go provides Go SDK to work with the Selectel Managed Kubernetes Service.

## Documentation

The Go library documentation is available at [go.dev](https://pkg.go.dev/github.com/selectel/mks-go/).

## What this library is capable of

You can use this library to work with the following objects of the Selectel Managed Kubernetes Service:

* [cluster](https://pkg.go.dev/github.com/selectel/mks-go/pkg/v1/cluster)
* [nodegroup](https://pkg.go.dev/github.com/selectel/mks-go/pkg/v1/nodegroup)
* [node](https://pkg.go.dev/github.com/selectel/mks-go/pkg/v1/node)
* [task](https://pkg.go.dev/github.com/selectel/mks-go/pkg/v1/task)
* [kubeversion](https://pkg.go.dev/github.com/selectel/mks-go/pkg/v1/kubeversion)

## Getting started

### Installation

You can install needed `mks-go` packages via `go get` command:

```bash
go get github.com/selectel/mks-go/pkg/v1/cluster github.com/selectel/mks-go/pkg/v1/task
```

### Authentication

To work with the Selectel Managed Kubernetes Service API you first need to:

* Create a Selectel account: [registration page](https://my.selectel.ru/registration).
* Create a project in Selectel Cloud Platform [projects](https://my.selectel.ru/vpc/projects).
* Retrieve a token for your project via API or [go-selvpcclient](https://github.com/selectel/go-selvpcclient).

### Usage example

In progress.