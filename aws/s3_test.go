package aws

import (
	"context"
	"os"
	"testing"
)

func TestConnectToAWS(t *testing.T) {
	ctx := context.TODO()

	os.Setenv("AWS_REGION", "us-east-1")

	client, err := ConnectToAWS(ctx)
	if err != nil {
		t.Errorf("ConnectToAWS() error = %v", err)
	}

	_ = client
}
