package storage

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var client *minio.Client;
const BUCKET_NAME string = "document-order-data"

func InitializeMinioClient() {
	accessKey := os.Getenv("MINIO_ACCESS_KEY");
	secretKey := os.Getenv("MINIO_SECRET_KEY");
	address := os.Getenv("MINIO_ADDRESS")

	retries := 5;
	timeout := 2;

	for range retries {
		minioClient, err := minio.New(address, &minio.Options{
			Creds: credentials.NewStaticV4(accessKey, secretKey, ""),
			Secure: false,
		})
		if err == nil {
			log.Println("Connected successfully to minio client")
			client = minioClient
			return
		} else {
			log.Println("Error while connecting to minio instance")
			time.Sleep(time.Second * time.Duration(timeout))
		}
	}
}

func UploadData(data io.Reader, name string, size int64) {
	ctx := context.Background();
	_, err := client.PutObject(ctx, BUCKET_NAME, name, data, size, minio.PutObjectOptions{});
	if err != nil {
		log.Printf("storage.UploadData ERROR: Error while calling put object: %s\n", err.Error())
	} else {
		log.Printf("storage.UploadData: uploaded order %s\n", name)
	}
}
