package s3filestore

import (
	"bytes"
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// S3FileStore follow Store interface using S3
type S3FileStore struct {
	client     *s3.S3
	downloader *s3manager.Downloader
	bucketName string
}

const (
	// ACL stands for Access-Control-List and manage access to bucket
	ACL = "public-read"
)

var (
	// ErrInitAWSSession is returned when AWS cannot be initialized
	ErrInitAWSSession = errors.New("error while initializing AWS session")
)

// NewS3Store returns a instance of S3Store
func NewS3Store(endpoint, bucketName string) (*S3FileStore, error) {
	sess := session.Must(session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials("id", "secret", ""),
		S3ForcePathStyle: aws.Bool(true),
		Region:           aws.String(endpoints.UsEast1RegionID),
		Endpoint:         aws.String(endpoint),
	}))

	downloader := s3manager.NewDownloader(sess)

	store := &S3FileStore{
		client:     s3.New(sess),
		bucketName: bucketName,
		downloader: downloader,
	}

	return store, nil
}

// Get a file
func (s *S3FileStore) Get(name string) ([]byte, error) {
	buf := aws.NewWriteAtBuffer([]byte{})
	_, err := s.downloader.Download(buf,
		&s3.GetObjectInput{
			Bucket: aws.String(s.bucketName),
			Key:    aws.String(name),
		})
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Delete a file
func (s *S3FileStore) Delete(name string) error {
	_, err := s.client.DeleteObject(
		&s3.DeleteObjectInput{
			Bucket: aws.String(s.bucketName),
			Key:    aws.String(name),
		})

	if err != nil {
		return err
	}

	err = s.client.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(name),
	})
	if err != nil {
		return err
	}

	return nil
}

// Update a file
func (s *S3FileStore) Update(pic []byte, name string) error {
	_, err := s.Insert(pic, name)

	return err
}

// Insert a file
func (s *S3FileStore) Insert(pic []byte, name string) (int, error) {
	reader := bytes.NewReader(pic)

	poi := s3.PutObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(name),
		ACL:    aws.String(ACL),
		Body:   reader,
	}

	_, err := s.client.PutObject(&poi)

	return 0, err
}
