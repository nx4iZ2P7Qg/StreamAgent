package minio

import (
	"StreamAgent/internal/cus/config"
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/tal-tech/go-zero/core/logx"
)

func getClient() *minio.Client {
	endpoint := config.C.Minio.Endpoint
	accessKeyID := config.C.Minio.AccessKeyID
	secretAccessKey := config.C.Minio.SecretAccessKey

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		logx.Infof("create minio client failed, %v\n", err)
	}

	//logx.Infof("%#v\n", minioClient)

	return minioClient
}

func CreateBucket(bucketName string, location string) error {
	ctx := context.Background()
	minioClient := getClient()

	err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			logx.Infof("bucket exists %s\n", bucketName)
			return nil
		} else {
			logx.Infof("bucket unexpected status %v\n", err)
			return err
		}
	} else {
		logx.Infof("bucket created %s\n", bucketName)
		return nil
	}
}

func Upload(bucketName string, objectName string, filePath string) error {

	err := CreateBucket(bucketName, "")
	if err != nil {
		logx.Errorf("create bucket failed, %v\n", err)
		return err
	}

	ctx := context.Background()
	info, err := getClient().FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{})
	if err != nil {
		logx.Errorf("upload failed, %v\n", err)
		return err
	}

	logx.Infof("object uploaded %s of size %d\n", objectName, info.Size)
	return nil
}
