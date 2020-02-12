package testutils

import (
	"net/http"
	"net/http/httptest"
)

// TestEnv represents a testing environment for all resources.
type TestEnv struct {
	Mux    *http.ServeMux
	Server *httptest.Server
}

// SetupTestEnv prepares the new testing environment.
func SetupTestEnv() *TestEnv {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	return &TestEnv{
		Mux:    mux,
		Server: server,
	}
}

// TearDownTestEnv releases the testing environment.
func (testEnv *TestEnv) TearDownTestEnv() {
	testEnv.Server.Close()
	testEnv.Server = nil
	testEnv.Mux = nil
}
