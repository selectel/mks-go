/*
Package cluster provides the ability to retrieve and manage Kubernetes clusters
through the MKS V1 API.

Example of getting a single cluster referenced by its id

  mksCluster, _, err := cluster.Get(ctx, mksClient, id)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%+v\n", mksCluster)
*/
package cluster
