package usecase

import (
	"errors"
	"implementation-minio-gcs-golang/configs"
	entity "implementation-minio-gcs-golang/entities"
	"implementation-minio-gcs-golang/provider"
	"log"
	"os"
	"path/filepath"
)

type GetService struct {
	config *configs.Config
}

// NewGetService membuat instance baru dari get service
func NewGetService(config *configs.Config) *GetService {
	return &GetService{
		config: config,
	}
}

// GetFile get file menggunakan penyedia penyimpanan yang ditentukan
func (g *GetService) GetFile(fileName, bucketName, filePath string) (*entity.FileResult, error) {
	// Memeriksa layanan yang aktif berdasarkan kredensial
	var fileResult *entity.FileResult

	// Buat folder penyimpanan jika tidak ditemukan
	if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
		log.Fatalf("gagal membuat folder penyimpanan: %v", err)
	}

	switch {
	case g.config.MinioEnabled:
		// Inisialisasi layanan autentikasi MinIO
		minIOAuthService, err := provider.NewMinIOAuthService(g.config.MinioEndpoint, g.config.MinioAccessKey, g.config.MinioSecretAccessKey, g.config.MinioSSL)
		if err != nil {
			log.Fatalf("gagal connect ke MinIO: %s", err)
		}

		destinationFile := filepath.Join(filePath, fileName)

		if err := minIOAuthService.GetFile(fileName, bucketName, destinationFile); err != nil {
			return nil, err
		}

		fileResult = &entity.FileResult{
			Source:          "MinIO",
			Bucket:          bucketName,
			Object:          fileName,
			DestinationFile: filePath,
		}
	case g.config.GCSEnabled:
		// Inisialisasi layanan penyimpanan file GCS
		gcsAuthService, err := provider.NewGCSAuthService("", g.config.GCSCredentialsFile)
		if err != nil {
			log.Fatalf("gagal connect ke GCS: %s", err)
		}

		if err := gcsAuthService.GetFile(fileName, bucketName, filePath); err != nil {
			return nil, err
		}

		fileResult = &entity.FileResult{
			Source:          "GCS",
			Bucket:          bucketName,
			Object:          fileName,
			DestinationFile: filePath,
		}
	default:
		return nil, errors.New("gagal tidak ada layanan aktif yang ditentukan")
	}

	return fileResult, nil
}
