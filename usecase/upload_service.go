package usecase

import (
	"errors"
	configs "implementasi-minio-gcs-golang/configs"
	entity "implementasi-minio-gcs-golang/entities"
	"implementasi-minio-gcs-golang/provider"
	"log"
)

type UploadService struct {
	config *configs.Config
}

// NewUploadService membuat instance baru dari upload service
func NewUploadService(config *configs.Config) *UploadService {
	return &UploadService{
		config: config,
	}
}

// UploadFile upload file menggunakan penyedia penyimpanan yang ditentukan
func (u *UploadService) UploadFile(filePath string, bucketName string, fileName string) (*entity.UploadResult, error) {
	// Memeriksa layanan yang aktif berdasarkan kredensial
	var uploadResult *entity.UploadResult

	switch {
	case u.config.MinioEnabled:
		// Inisialisasi layanan autentikasi MinIO
		minIOAuthService, err := provider.NewMinIOAuthService(u.config.MinioEndpoint, u.config.MinioAccessKey, u.config.MinioSecretAccessKey, u.config.MinioSSL)
		if err != nil {
			log.Fatalf("gagal connect ke MinIO: %s", err)
		}

		result, err := minIOAuthService.UploadFile(filePath, bucketName, fileName)
		if err != nil {
			return nil, err
		}

		uploadResult = result
	case u.config.GCSEnabled:
		// Inisialisasi layanan penyimpanan file GCS
		gcsAuthService, err := provider.NewGCSAuthService("", u.config.GCSCredentialsFile)
		if err != nil {
			log.Fatalf("gagal connect ke GCS: %s", err)
		}

		result, err := gcsAuthService.UploadFile(filePath, bucketName, fileName)
		if err != nil {
			return nil, err
		}

		uploadResult = result
	default:
		return nil, errors.New("gagal tidak ada layanan aktif yang ditentukan")
	}

	return uploadResult, nil
}
