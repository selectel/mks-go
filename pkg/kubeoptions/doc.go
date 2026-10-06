/*
Package kubeoptions provides the ability to retrieve all available Kubernetes
feature gates and admission controllers through the MKS V2 API.

Example of getting available feature gates by Kubernetes version:

	availableFG, err := kubeoptions.ListFeatureGates(ctx, client)
	if err != nil {
	  log.Fatal(err)
	}
	for _, fgList := range availableFG {
	  fmt.Printf("%s: %v\n", *fgList.KubeVersionMinor, *fgList.Names)
	}
*/
package kubeoptions
