package testing

import "github.com/selectel/mks-go/pkg/v1/kubeversion"

// testListKubeVersionsResponseRaw represents a raw response from the List request.
const testListKubeVersionsResponseRaw = `
{
    "kube_versions": [
        {
            "is_default": true,
            "version": "1.15.10"
        },
        {
            "is_default": false,
            "version": "1.15.7"
        }
    ]
}
`

// expectedListKubeVersionsResponse represents an unmarshalled testListKubeVersionsResponseRaw.
var expectedListKubeVersionsResponse = []*kubeversion.View{
	{
		Version:   "1.15.10",
		IsDefault: true,
	},
	{
		Version:   "1.15.7",
		IsDefault: false,
	},
}

// testManyKubeVersionsInvalidResponseRaw represents a raw invalid response with several Kubernetes versions.
const testManyKubeVersionsInvalidResponseRaw = `
{
    "kube_versions": [
        {
            "version": "1.15.10",
        }
    ]
}
`

// testErrGenericResponseRaw represents a raw response with an error in the generic format.
const testErrGenericResponseRaw = `{"error":{"message":"bad gateway"}}`
