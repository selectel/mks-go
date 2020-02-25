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

### Endpoints

Selectel Managed Kubernetes Service currently has the following API endpoints:

| URL                             | Region |
|---------------------------------|--------|
| https://ru-1.mks.selcloud.ru/v1 | ru-1   |
| https://ru-2.mks.selcloud.ru/v1 | ru-2   |
| https://ru-3.mks.selcloud.ru/v1 | ru-3   |
| https://ru-7.mks.selcloud.ru/v1 | ru-7   |
| https://ru-8.mks.selcloud.ru/v1 | ru-8   |

You can also retrieve all available API endpoints from the Identity catalog.

### Usage example

```go
package main

import (
	"context"
	"fmt"
	"log"

	v1 "github.com/selectel/mks-go/pkg/v1"
	"github.com/selectel/mks-go/pkg/v1/cluster"
	"github.com/selectel/mks-go/pkg/v1/kubeversion"
	"github.com/selectel/mks-go/pkg/v1/nodegroup"
	"github.com/selectel/mks-go/pkg/v1/task"
)

func main() {
	// Token to work with Selectel Cloud project.
	token := "gAAAAABeVNzu-..."

	// MKS endpoint to work with.
	endpoint := "https://ru-3.mks.selcloud.ru/v1"

	// Initialize the MKS V1 client.
	mksClient := v1.NewMKSClientV1(token, endpoint)

	// Prepare empty context.
	ctx := context.Background()

	// Get supported Kubernetes versions.
	kubeVersions, _, err := kubeversion.List(ctx, mksClient)
	if err != nil {
		log.Fatal(err)
	}
	if len(kubeVersions) == 0 {
		log.Fatal("There are no available Kubernetes versions")
	}

	// Use the first version in list.
	kubeVersion := kubeVersions[0]

	// Nodegroup with nodes based on network volumes for root partition.
	firstNodegroup := &nodegroup.CreateOpts{
		Count:            3,
		CPUs:             1,
		RAMMB:            2048,
		VolumeGB:         50,
		VolumeType:       "fast.ru-3a",
		AvailabilityZone: "ru-3a",
	}

	// Nodegroup with nodes based on local volumes for root partition.
	secondNodegroup := &nodegroup.CreateOpts{
		Count:            2,
		CPUs:             2,
		RAMMB:            4096,
		VolumeGB:         20,
		LocalVolume:      true,
		AvailabilityZone: "ru-3a",
	}

	// Build final options for a new cluster.
	createOpts := &cluster.CreateOpts{
		Name:        "test-cluster",
		KubeVersion: kubeVersion.Version,
		Region:      "ru-3",
		Nodegroups: []*nodegroup.CreateOpts{
			firstNodegroup,
			secondNodegroup,
		},
	}

	// Create a cluster.
	newCluster, _, err := cluster.Create(ctx, mksClient, createOpts)
	if err != nil {
		log.Fatal(err)
	}

	// Print cluster fields.
	fmt.Printf("Created cluster: %+v\n", newCluster)

	// Get cluster tasks.
	tasks, _, err := task.List(ctx, mksClient, newCluster.ID)
	if err != nil {
		log.Fatal(err)
	}

	// Print cluster tasks.
	for _, newClusterTask := range tasks {
		fmt.Printf("Cluster task: %+v\n", newClusterTask)
	}
}
```