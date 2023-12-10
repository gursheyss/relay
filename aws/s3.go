package aws

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func ConnectToAWS(ctx context.Context) (*s3.Client, error) {
	awsRegion := os.Getenv("AWS_REGION")

	if awsRegion == "" {
		awsRegion = "us-east-1"
	}

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(awsRegion),
	)

	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg)
	log.Println("Successfully connected to AWS")
	return client, nil
}
