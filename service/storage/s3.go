package storage

import (
	"github.com/AH-dark/bing-wallpaper/model"
	"github.com/AH-dark/logger"
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

	settings, err := model.GetSettings([]string{
		"storage_s3_endpoint",
		"storage_s3_bucket",
		"storage_s3_region",
		"storage_s3_access_key",
		"storage_s3_secret_key",
		"storage_s3_use_ssl",
	})
	if err != nil {
		logger.Log().Errorf("get s3 settings error: %s", err)
		return nil
	}

	client := s3.New(sess, &aws.Config{
		Credentials: credentials.NewStaticCredentials(settings["storage_s3_access_key"], settings["storage_s3_secret_key"], ""),
		Endpoint:    aws.String(settings["storage_s3_endpoint"]),
		Region:      aws.String(settings["storage_s3_region"]),
		DisableSSL:  aws.Bool(settings["storage_s3_use_ssl"] == "false"),
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
	base, err := model.GetSettingVal("storage_s3_path")
	if err != nil {
		return nil, err
	}

	key := filepath.Join(base, name)

	_, err = s.uploader.Upload(&s3manager.UploadInput{
		Key:  &key,
		Body: file,
	})
	if err != nil {
		return nil, err
	}

	base, err = model.GetSettingVal("storage_s3_base_url")
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(base)
	if err != nil {
		return nil, err
	}

	u, err = u.Parse(key)
	if err != nil {
		return nil, err
	}

	return u, nil
}
