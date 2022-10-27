package storage

import (
	"fmt"
	"github.com/AH-dark/bing-wallpaper/pkg/conf"
	"github.com/AH-dark/bing-wallpaper/pkg/util"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
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
		Key:                &key,
		Body:               file,
		Bucket:             &conf.StorageConfig.Bucket,
		ACL:                aws.String(conf.StorageConfig.ACL),
		ContentType:        aws.String("image/jpeg"),
		ContentDisposition: aws.String("inline"),
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

func (s *S3Impl) Test() error {
	fileKey := filepath.Join(conf.StorageConfig.BasePath, fmt.Sprintf("test-%s", util.RandString(8)))

	_, err := s.uploader.Upload(&s3manager.UploadInput{
		Key:                &fileKey,
		Body:               io.NopCloser(strings.NewReader("test")),
		Bucket:             &conf.StorageConfig.Bucket,
		ACL:                aws.String(conf.StorageConfig.ACL),
		ContentType:        aws.String("text/plain"),
		ContentDisposition: aws.String("inline"),
	})
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Get(fmt.Sprintf("%s/%s", util.RemoveSlash(conf.StorageConfig.BaseUrl), fileKey))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	d, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	} else if string(d) != "test" {
		return fmt.Errorf("unexpected body: %s", string(d))
	}

	_, err = s.client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: &conf.StorageConfig.Bucket,
		Key:    &fileKey,
	})
	if err != nil {
		return err
	}

	return nil
}
