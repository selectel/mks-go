/*
Package node provides the ability to retrieve and manage Kubernetes nodes
of a cluster nodegroup through the MKS V2 API.

Example of getting a single node of a cluster nodegroup by its id

	singleNode, err := node.Get(ctx, client, clusterID, nodegroupID, nodeID)
	if err != nil {
	  log.Fatal(err)
	}
	fmt.Printf("%+v\n", singleNode)
*/
package node
