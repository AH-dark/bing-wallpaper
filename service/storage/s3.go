package storage

import (
	"github.com/AH-dark/bing-wallpaper/pkg/conf"
	"github.com/AH-dark/bing-wallpaper/pkg/util"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
	"net/url"
	"path/filepath"
)

type S3Impl struct {
	client   *s3.S3
	uploader *s3manager.Uploader
}

func NewS3() *S3Impl {
	sess := session.Must(session.NewSession())

	client := s3.New(sess, &aws.Config{
		Credentials: credentials.NewStaticCredentials(conf.StorageConfig.AccessID, conf.StorageConfig.AccessKey, ""),
		Endpoint:    &conf.StorageConfig.Endpoint,
		Region:      &conf.StorageConfig.Region,
		DisableSSL:  &conf.StorageConfig.SSL,
	})

	uploader := s3manager.NewUploader(sess, func(u *s3manager.Uploader) {
		u.S3 = client
	})

	return &S3Impl{
		client:   client,
		uploader: uploader,
	}
}

func (s *S3Impl) Upload(name string, file io.Reader) (*url.URL, error) {
	key := filepath.Join(conf.StorageConfig.BasePath, name)
	key = util.FormSlash(key)

	_, err := s.uploader.Upload(&s3manager.UploadInput{
		Key:    &key,
		Body:   file,
		Bucket: &conf.StorageConfig.Bucket,
	})
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(conf.StorageConfig.BaseUrl)
	if err != nil {
		return nil, err
	}

	u, err = u.Parse(key)
	if err != nil {
		return nil, err
	}

	return u, nil
}
