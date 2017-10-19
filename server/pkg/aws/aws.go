package aws

import (
	"github.com/aws/aws-sdk-go/aws/credentials"

	"github.com/aws/aws-sdk-go/aws"
)

// Adapter wraps around AWS
type Adapter struct {
	Config *aws.Config
}

// New creates a new adapter forAWS
func New(region, secretKey, accessKey, token string) Adapter {
	creds := credentials.NewStaticCredentials(accessKey, secretKey, token)

	return Adapter{
		Config: aws.NewConfig().WithRegion(region).WithCredentials(creds),
	}
}
