package aws

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

func InitAWS() error {
	sess, err := session.NewSession()
	if err != nil {
		return err
	}

	_, err = credentials.NewEnvCredentials().Get()
	if err != nil {
		return err
	}

	svc := sts.New(sess)
	_, err = svc.GetCallerIdentity(&sts.GetCallerIdentityInput{})
	if err != nil {
		return err
	}

	return nil
}
