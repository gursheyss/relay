package server

import (
	"testing"
)

func TestSetupServer(t *testing.T) {
	_, err := SetupServer()
	if err != nil {
		t.Errorf("SetupServer() error = %v", err)
	}
}
