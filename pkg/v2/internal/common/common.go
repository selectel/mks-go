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
	NameConflict                  = "conflict"
	NameEmptyFeatureGates         = "empty feature gates"
	NameEmptyAdmissionControllers = "empty admission controllers"
	NameEmptyKubeVersions         = "empty kube_versions"
	NameEmptyNodegroups           = "empty nodegroups"
	NameEmptyRegistries           = "empty registries"
)

// Error message constants.
const (
	MsgInternalError      = "internal server error"
	MsgClusterNotFound    = "cluster not found"
	MsgNodeNotFound       = "node not found"
	MsgBadRequest         = "bad request"
	MsgConflict           = "conflict"
	MsgNodegroupNotFound  = "nodegroup not found"
	MsgRegistriesNotFound = "registries not found"
	MsgRegistryNotFound   = "registry not found"
)

// Ptr returns a pointer to the given value.
func Ptr[T any](value T) *T {
	return &value
}
