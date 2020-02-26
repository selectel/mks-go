/*
Package node provides the ability to retrieve and manage Kubernetes nodes
of a cluster nodegroup through the MKS V1 API.

Example of getting a single node of a cluster nodegroup by its id

  singleNode, _, err := node.Get(ctx, mksClient, clusterID, nodegroupID, nodeID)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%+v\n", singleNode)
*/
package node
