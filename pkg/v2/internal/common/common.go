// Package common provides helper utilities.
package common

// Test case name constants.
const (
	NameSuccess                   = "success"
	NameInternalError             = "internal server error"
	NameNotFound                  = "not found"
	NameUnknownStatus             = "unknown status"
	NameHTTPError                 = "http error"
	NameBadRequest                = "bad request"
	NameEmptyFeatureGates         = "empty feature gates"
	NameEmptyAdmissionControllers = "empty admission controllers"
	NameEmptyKubeVersions         = "empty kube_versions"
	NameEmptyRegistries           = "empty registries"
)

// Error message constants.
const (
	MsgInternalError      = "internal server error"
	MsgClusterNotFound    = "cluster not found"
	MsgRegistriesNotFound = "registries not found"
	MsgRegistryNotFound   = "registry not found"
	MsgBadRequest         = "bad request"
)

// Ptr returns a pointer to the given value.
func Ptr[T any](value T) *T {
	return &value
}
