package provider

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	entity "implementation-minio-gcs-golang/entities"
	"os"
	"path/filepath"
)

type MinIOConfig struct {
	Endpoint        string `json:"endpoint"`
	AccessKeyID     string `json:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key"`
	UseSSL          bool   `json:"use_ssl"`
}

// MinIOAuthService mengimplementasikan file storage service untuk MinIO
type MinIOAuthService struct {
	minioClient *minio.Client
}

// NewMinIOAuthService menginisialisasi MiniIO Service
func NewMinIOAuthService(endpoint, accessKey, secretKey string, ssl bool) (*MinIOAuthService, error) {
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:        credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure:       ssl, // Set to true if using HTTPS
		BucketLookup: minio.BucketLookupAuto,
	})
	if err != nil {
		return nil, err
	}

	return &MinIOAuthService{
		minioClient: minioClient,
	}, nil
}

// UploadFile mengupload file ke MinIO storage
func (m *MinIOAuthService) UploadFile(fileName, bucketName, filePath string) (*entity.UploadResult, error) {
	// Membuat cintext context
	ctx := context.TODO()

	// Membuka file
	file, err := os.Open(filepath.Join(filePath, fileName))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Mengambil statistik file
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	// Set optional metadata
	options := minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	}

	etag, err := m.minioClient.PutObject(ctx, bucketName, fileName, file, stat.Size(), options)
	if err != nil {
		return nil, err
	}

	return &entity.UploadResult{
		Etag:      etag.ETag,
		PublicURL: etag.Location,
	}, nil
}

// GetFile mengambil file dari MinIO storage
func (m *MinIOAuthService) GetFile(fileName, bucketName, destinationFile string) error {
	// Membuat cintext context
	ctx := context.TODO()

	err := m.minioClient.FGetObject(ctx, bucketName, fileName, destinationFile, minio.GetObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}
