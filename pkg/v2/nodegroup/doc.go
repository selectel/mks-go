/*
Package nodegroup provides the ability to retrieve and manage cluster nodegroups
through the MKS V2 API.

Example of getting a single cluster nodegroup referenced by its id

	clusterNodegroup, err := nodegroup.Get(ctx, client, clusterID, nodegroupID)
	if err != nil {
	  log.Fatal(err)
	}
	fmt.Printf("%+v\n", clusterNodegroup)
*/
package nodegroup
