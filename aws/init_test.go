package aws

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestInitAWS(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatalf("Error loading .env file")
	}

	err = InitAWS()
	if err != nil {
		t.Errorf("InitAWS() error = %v", err)
	}
}
