# mks-go: Go SDK for Managed Kubernetes Service
[![Go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/selectel/mks-go/v2)
![Build Status](https://github.com/selectel/mks-go/actions/workflows/v2-unit-tests.yml/badge.svg?branch=sdk-v2)

Package mks-go provides Go SDK to work with the Selectel Managed Kubernetes Service.

## Documentation

The Go library documentation is available at [go.dev](https://pkg.go.dev/github.com/selectel/mks-go/v2/).

## What this library is capable of

You can use this library to work with the following objects of the Selectel Managed Kubernetes Service:

* [cluster](https://pkg.go.dev/github.com/selectel/mks-go/v2/pkg/cluster)
* [nodegroup](https://pkg.go.dev/github.com/selectel/mks-go/v2/pkg/nodegroup)
* [node](https://pkg.go.dev/github.com/selectel/mks-go/v2/pkg/node)
* [task](https://pkg.go.dev/github.com/selectel/mks-go/v2/pkg/task)
* [kubeversion](https://pkg.go.dev/github.com/selectel/mks-go/v2/pkg/kubeversion)
* [kubeoptions](https://pkg.go.dev/github.com/selectel/mks-go/v2/pkg/kubeoptions)
* [registries](https://pkg.go.dev/github.com/selectel/mks-go/v2/pkg/registries)

## Getting started

### Installation

You can install needed `mks-go` packages via `go get` command:

```bash
go get github.com/selectel/mks-go/v2/pkg/cluster github.com/selectel/mks-go/v2/pkg/task
```

### Authentication

To work with the Selectel Managed Kubernetes Service API you first need to:

* Create a Selectel account: [registration page](https://my.selectel.ru/registration).
* Create a project in Selectel Cloud Platform [projects](https://my.selectel.ru/iam/projects).
* Retrieve a token for your project via API or [go-selvpcclient](https://github.com/selectel/go-selvpcclient).

### Endpoints

* Selectel Managed Kubernetes Service currently has the following API endpoints: [URLs](https://docs.selectel.ru/en/api/urls/#managed-kubernetes)

> [!NOTE]
> mks-go/v2 designed to work with Managed Kubernetes API v2

### Usage example

```go
package main

import (
	"context"
	"fmt"
	"log"

	mks "github.com/selectel/mks-go/v2/pkg"
	"github.com/selectel/mks-go/v2/pkg/cluster"
	"github.com/selectel/mks-go/v2/pkg/kubeversion"
	"github.com/selectel/mks-go/v2/pkg/mksclient"
	"github.com/selectel/mks-go/v2/pkg/nodegroup"
	"github.com/selectel/mks-go/v2/pkg/task"
)

func main() {
	// Token to work with Selectel Cloud project.
	token := "gAAAAABeVNzu-..."

	// MKS endpoint to work with.
	endpoint := "https://ru-3.mks.selcloud.ru/v2"

	// Initialize the MKS V2 client.
	mksClient, err := mks.NewMKSClientV2(token, endpoint)
	if err != nil {
		log.Fatal(err)
	}

	// Prepare empty context.
	ctx := context.Background()

	// Get supported Kubernetes versions.
	kubeVersions, err := kubeversion.List(ctx, mksClient)
	if err != nil {
		log.Fatal(err)
	}
	if len(kubeVersions) == 0 {
		log.Fatal("There are no available Kubernetes versions")
	}

	// Use the first version in list.
	kubeVersion := kubeVersions[0]

	// Build final options for a new cluster.
	createOpts := &mksclient.ClusterCreateStruct{
		Name:        "test-cluster",
		KubeVersion: *kubeVersion.Version,
		Pool:        "ru-3",
		Basic:       true,
		NetworkType: "standard",
	}

	// Create a cluster.
	newCluster, err := cluster.Create(ctx, mksClient, createOpts)
	if err != nil {
		log.Fatal(err)
	}

	// Print cluster fields.
	fmt.Printf("Created cluster: %+v\n", newCluster)

	// Nodegroup with nodes based on network volumes for root partition.
	firstNodegroup := mksclient.NodegroupCreateStruct{
		Count:   3,
		Segment: "ru-3a",
		CloudNodegroupConfig: &mksclient.CloudNodegroupConfig{
			Cpus:       1,
			RamMb:      2048,
			VolumeGb:   50,
			VolumeType: "fast.ru-3a",
		},
	}

	// Nodegroup with nodes based on local volumes for root partition.
	secondNodegroup := mksclient.NodegroupCreateStruct{
		Count:   2,
		Segment: "ru-3a",
		CloudNodegroupConfig: &mksclient.CloudNodegroupConfig{
			Cpus:        2,
			RamMb:       4096,
			VolumeGb:    20,
			LocalVolume: true,
		},
	}

	// Create node groups for the new cluster.
	err = nodegroup.Create(ctx, mksClient, newCluster.Id, []mksclient.NodegroupCreateStruct{firstNodegroup, secondNodegroup})
	if err != nil {
		log.Fatal(err)
	}

	// Get cluster tasks.
	tasks, err := task.List(ctx, mksClient, newCluster.Id, 10, 0)
	if err != nil {
		log.Fatal(err)
	}

	// Print cluster tasks.
	for _, t := range tasks {
		fmt.Printf("Cluster task: %+v\n", t)
	}
}
```
