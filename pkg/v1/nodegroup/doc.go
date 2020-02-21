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

Example of creating a new cluster nodegroup

  createOpts := &nodegroup.CreateOpts{
    Count:            1,
    CPUs:             1,
    RAMMB:            2048,
    VolumeGB:         10,
    VolumeType:       "fast.ru-3a",
    KeypairName:      "ssh-key",
    AvailabilityZone: "ru-3a",
  }
  _, err := nodegroup.Create(ctx, mksClient, clusterID, createOpts)
  if err != nil {
    log.Fatal(err)
  }
*/
package nodegroup
