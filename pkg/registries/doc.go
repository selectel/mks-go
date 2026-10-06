/*
Package registries provides the ability to retrieve and manage cluster registries
through the MKS V2 API.

Example of getting cluster registries

	clusterRegistriesIntegration, err := registries.Get(ctx, client, clusterID)
	if err != nil {
	  log.Fatal(err)
	}
	for _, registry := range clusterRegistriesIntegration.Registries {
	  fmt.Printf("%+v\n", registry)
	}
*/
package registries
