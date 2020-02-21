/*
Package nodegroup provides the ability to retrieve and manage cluster nodegroups
through the MKS V1 API.

Example of getting a single cluster nodegroup referenced by its id

  clusterNodegroup, _, err := nodegroup.Get(ctx, mksClient, clusterID, nodegroupID)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%+v\n", clusterNodegroup)

Example of getting all cluster nodegroups

  clusterNodegroups, _, err := nodegroup.List(ctx, mksClient, clusterID)
  if err != nil {
    log.Fatal(err)
  }
  for _, clusterNodegroup := range clusterNodegroups {
    fmt.Printf("%+v\n", clusterNodegroup)
  }
*/
package nodegroup
