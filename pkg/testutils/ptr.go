package testutils

// BoolToPtr can be used to convert boolean value to boolean pointer.
func BoolToPtr(v bool) *bool {
	return &v
}

// IntToPtr can be used to convert integer value to integer pointer.
func IntToPtr(v int) *int {
	return &v
}
