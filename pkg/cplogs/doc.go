/*
Package cplogs provides the ability to retrieve and manage control plane logs
through the MKS V2 API.

Example of getting control plane logging settings

	loggingSettings, err := cplogs.Get(ctx, client, clusterID)
	if err != nil {
	  log.Fatal(err)
	}
	fmt.Printf("%+v\n", loggingSettings)
*/
package cplogs
