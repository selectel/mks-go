/*
Package task provides the ability to retrieve cluster tasks through the MKS V1 API.

Example of getting a single cluster task referenced by its id

  clusterTask, _, err := task.Get(ctx, mksClient, clusterID, taskID)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%+v\n", clusterTask)

Example of getting all cluster tasks

  clusterTasks, _, err := task.List(ctx, mksClient, clusterID)
  if err != nil {
    log.Fatal(err)
  }
  for _, clusterTask := range clusterTasks {
    fmt.Printf("%+v\n", clusterTask)
  }
*/
package task
