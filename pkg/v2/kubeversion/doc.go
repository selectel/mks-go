/*
Package kubeversion provides the ability to retrieve all supported Kubernetes
versions through the MKS V2 API.

Example of getting all supported Kubernetes versions

	kubeVersions, err := kubeversion.List(ctx, client)
	if err != nil {
	  log.Fatal(err)
	}
	for _, version := range kubeVersions {
	  fmt.Printf("%+v\n", version)
	}
*/
package kubeversion
