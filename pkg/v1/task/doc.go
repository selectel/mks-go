/*
Package task provides the ability to retrieve cluster tasks through the MKS V1 API.

Example of getting a single cluster task referenced by its id

  clusterTask, _, err := task.Get(ctx, mksClient, clusterID, taskID)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%+v\n", clusterTask)
*/
package task
