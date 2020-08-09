package configuration

import (
	"errors"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/go-chi/jwtauth"

	"github.com/rbonnat/codecademy/mysqlstore"
	"github.com/rbonnat/codecademy/s3filestore"
)

const tokenSecretKey = "authorization/secret-key"

var (
	// ErrMissingPort is returned when env var PORT is missing
	ErrMissingPort = errors.New("missing environment variable: PORT")

	// ErrMissingBucketName is returned when env var BUCKET_NAME is missing
	ErrMissingBucketName = errors.New("missing environment variable: BUCKET_NAME")

	// ErrMissingAWSEndpoint is returned when env var S3_ENDPOINT is missing
	ErrMissingAWSEndpoint = errors.New("missing environment variable: AWS_ENDPOINT")

	// ErrMissingDSN is returned when env var DSN is missing
	ErrMissingDSN = errors.New("missing environment variable: DSN")
)

// Configuration contains configuration variable
type Configuration struct {
	FileStore *s3filestore.S3FileStore
	DBStore   *mysqlstore.MySQLStore

	TokenSecretKey string
	TokenAuth      *jwtauth.JWTAuth

	Port        string
	Bucket      string
	AWSEndpoint string
	DSN         string
}

// Load return an instance of Configuration with all variables initialized
func Load(envVarStore VarStore) (*Configuration, error) {
	cfg := &Configuration{}

	// Fetch environment variable
	err := fetchVar(cfg, envVarStore)
	if err != nil {
		return nil, err
	}

	// Init file store
	fs, err := s3filestore.NewS3Store(cfg.AWSEndpoint, cfg.Bucket)
	if err != nil {
		log.Printf("Cannot initialize file store: '%v'", err)
		return nil, err
	}
	cfg.FileStore = fs

	// Init DB store
	dbStore, err := mysqlstore.New(cfg.DSN)
	if err != nil {
		log.Printf("Cannot initialize db store: '%v'", err)
		return nil, err
	}
	cfg.DBStore = dbStore

	// Fetch private key from secrets manager
	secret, err := tokenSecretkey(cfg.AWSEndpoint, tokenSecretKey)
	if err != nil {
		log.Printf("Cannot fetch token secret key: '%v'", err)
		return nil, err
	}

	// Initialize Token Auth
	cfg.TokenAuth = jwtauth.New("HS256", []byte(secret), nil)

	return cfg, nil
}

func fetchVar(cfg *Configuration, envVarStore VarStore) error {
	cfg.Port = envVarStore.Get("PORT")
	if cfg.Port == "" {
		return ErrMissingPort
	}

	cfg.Bucket = envVarStore.Get("BUCKET_NAME")
	if cfg.Bucket == "" {
		return ErrMissingBucketName
	}

	cfg.AWSEndpoint = envVarStore.Get("AWS_ENDPOINT")
	if cfg.AWSEndpoint == "" {
		return ErrMissingAWSEndpoint
	}

	cfg.DSN = envVarStore.Get("DSN")
	if cfg.DSN == "" {
		return ErrMissingDSN
	}

	return nil
}

func tokenSecretkey(endpoint, key string) (string, error) {
	sess := session.Must(session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials("id", "secret", ""),
		S3ForcePathStyle: aws.Bool(true),
		Region:           aws.String(endpoints.UsEast1RegionID),
		Endpoint:         aws.String(endpoint),
	}))

	secret := secretsmanager.New(sess)

	s, err := secret.GetSecretValue(&secretsmanager.GetSecretValueInput{SecretId: &key})
	if err != nil {
		return "", err
	}

	return *s.SecretString, nil

}
