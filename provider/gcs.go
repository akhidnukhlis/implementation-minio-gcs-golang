package provider

import (
	"cloud.google.com/go/storage"
	"context"
	"google.golang.org/api/option"
	entity "implementasi-minio-gcs-golang/entities"
	"implementasi-minio-gcs-golang/helpers"
	"io"
	"os"
	"path/filepath"
)

// GCSAuthService mengimplementasikan file storage service untuk Google Cloud Storage (GCS)
type GCSAuthService struct {
	client *storage.Client
}

// NewGCSAuthService menginisialisasi GCS Service
func NewGCSAuthService(projectID, credentialFileName string) (*GCSAuthService, error) {
	var (
		ctx = context.Background()
	)

	dir := helpers.GetCurrentDirectory()
	credentialFilePath := filepath.Join(dir, "configs/", credentialFileName)

	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialFilePath))
	if err != nil {
		return nil, err
	}

	return &GCSAuthService{
		client: client,
	}, nil
}

// UploadFile mengupload file ke Google Cloud Storage (GCS)
func (g *GCSAuthService) UploadFile(filePath, bucketName, fileName string) (*entity.UploadResult, error) {
	ctx := context.TODO()
	bucket := g.client.Bucket(bucketName)
	object := bucket.Object(fileName)
	writer := object.NewWriter(ctx)
	defer writer.Close()

	file, err := os.Open(filepath.Join(filePath, fileName))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	if _, err := io.Copy(writer, file); err != nil {
		return nil, err
	}

	attrs, err := object.Attrs(ctx)
	if err != nil {
		return nil, err
	}

	return &entity.UploadResult{
		Etag:      attrs.Etag,
		PublicURL: attrs.MediaLink,
	}, nil
}
