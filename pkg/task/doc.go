/*
Package task provides the ability to retrieve cluster tasks through the MKS V2 API.

Example of getting a single cluster task referenced by its id

	clusterTask, err := task.Get(ctx, client, clusterID, taskID)
	if err != nil {
	  log.Fatal(err)
	}
	fmt.Printf("%+v\n", clusterTask)
*/
package task
