package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		name        string
		port        string
		bucketName  string
		AWSEndpoint string
		dsn         string
		expected    *Configuration
		err         error
	}{
		{
			name:        "Success",
			port:        "8080",
			bucketName:  "codecademy",
			AWSEndpoint: "http://localstack-codecademy:4566",
			dsn:         "root:@tcp(mysql:3306)/codecademy",
			expected: &Configuration{
				Port:        "8080",
				Bucket:      "codecademy",
				AWSEndpoint: "http://localstack-codecademy:4566",
				DSN:         "root:@tcp(mysql:3306)/codecademy",
			},
		},
		{
			name:        "Fail missing Port",
			bucketName:  "codecademy",
			AWSEndpoint: "http://localstack-codecademy:4566",
			dsn:         "root:@tcp(mysql:3306)/codecademy",
			err:         ErrMissingPort,
		},
		{
			name:        "Fail missing bucket name",
			port:        "8080",
			AWSEndpoint: "http://localstack-codecademy:4566",
			dsn:         "root:@tcp(mysql:3306)/codecademy",
			err:         ErrMissingBucketName,
		},
		{
			name:       "Fail missing s3 endpoint",
			port:       "8080",
			bucketName: "codecademy",
			dsn:        "root:@tcp(mysql:3306)/codecademy",
			err:        ErrMissingAWSEndpoint,
		},
		{
			name:        "Fail missing dsn",
			port:        "8080",
			bucketName:  "codecademy",
			AWSEndpoint: "http://localstack-codecademy:4566",
			err:         ErrMissingDSN,
		},
	}

	for _, test := range tests {
		mockStore := MockVarStore{}
		mockStore.On("Get", "PORT").Return(test.port)
		mockStore.On("Get", "BUCKET_NAME").Return(test.bucketName)
		mockStore.On("Get", "S3_ENDPOINT").Return(test.AWSEndpoint)
		mockStore.On("Get", "DSN").Return(test.dsn)
		cfg, err := Load(&mockStore)
		assert.Equal(t, test.expected, cfg)
		assert.Equal(t, test.err, err)
	}
}
