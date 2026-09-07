/*
Package cluster provides the ability to retrieve and manage Kubernetes clusters
through the MKS V2 API.

Example of getting a single cluster referenced by its id

	mksCluster, err := cluster.Get(ctx, client, clusterID)
	if err != nil {
	  log.Fatal(err)
	}
	fmt.Printf("%+v\n", mksCluster)
*/
package cluster
