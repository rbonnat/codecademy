package secretstore

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

type SecretStore struct {
	client *secretsmanager.SecretsManager
}

func NewSecretStore(endpoint string) *SecretStore {
	sess := session.Must(session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials("id", "secret", ""),
		S3ForcePathStyle: aws.Bool(true),
		Region:           aws.String(endpoints.UsEast1RegionID),
		Endpoint:         aws.String(endpoint),
	}))

	secret := secretsmanager.New(sess)

	return &SecretStore{client: secret}
}

func (s *SecretStore) Get(key string) (string, error) {
	secret, err := s.client.GetSecretValue(&secretsmanager.GetSecretValueInput{SecretId: &key})
	if err != nil {
		return "", err
	}

	return *secret.SecretString, nil
}
