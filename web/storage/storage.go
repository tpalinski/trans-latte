package storage

import (
	"log"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var client *minio.Client;

func InitializeMinioClient() {
	accessKey := os.Getenv("MINIO_ACCESS_KEY");
	secretKey := os.Getenv("MINIO_SECRET_KEY");
	address := os.Getenv("MINIO_ADDRESS")

	retries := 5;
	timeout := 2;

	for range retries {
		minioClient, err := minio.New(address, &minio.Options{
			Creds: credentials.NewStaticV4(accessKey, secretKey, ""),
			Secure: true,
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
