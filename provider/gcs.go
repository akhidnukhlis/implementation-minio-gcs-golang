package provider

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"google.golang.org/api/option"
	entity "implementation-minio-gcs-golang/entities"
	"implementation-minio-gcs-golang/helpers"
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
func (g *GCSAuthService) UploadFile(fileName, bucketName, filePath string) (*entity.UploadResult, error) {
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

// GCSAuthService mengambil file dari GCS storage
func (g *GCSAuthService) GetFile(fileName, bucketName, filePath string) error {
	ctx := context.TODO()
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("gagal membuat folder penyimpanan: %w", err)
	}
	defer file.Close()

	reader, err := g.client.Bucket(bucketName).Object(fileName).NewReader(ctx)
	if err != nil {
		return fmt.Errorf("gagal mengambil file dari GCS: %w", err)
	}
	defer reader.Close()

	if _, err := io.Copy(file, reader); err != nil {
		return fmt.Errorf("gagal memuat konten file: %w", err)
	}

	return nil
}
