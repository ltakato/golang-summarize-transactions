package services

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"cloud.google.com/go/storage"
)

type StorageClient struct {
	client storage.Client
	ctx    context.Context
}

func NewStorageClient(ctx context.Context) *StorageClient {
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("storage.NewClient: %v", err)
	}
	return &StorageClient{
		client: *client,
		ctx:    ctx,
	}
}

func (s *StorageClient) Upload(file *os.File, bucketName string, objectName string) error {
	var err error
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatalf("failed to close file: %v", err)
		}
	}()

	bucket := s.client.Bucket(bucketName)
	object := bucket.Object(objectName)
	writer := object.NewWriter(s.ctx)

	if _, err = io.Copy(writer, file); err != nil {
		log.Fatalf("io.Copy: %v", err)
	}
	if err = writer.Close(); err != nil {
		log.Fatalf("writer.Close: %v", err)
	}

	fmt.Printf("File %s uploaded to bucket as %s.\n", bucketName, objectName)

	return err
}
