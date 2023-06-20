package main

import (
	"fmt"
	"implementasi-minio-gcs-golang/configs"
	"implementasi-minio-gcs-golang/middleware"
	"implementasi-minio-gcs-golang/usecase"
	"log"
)

func main() {
	var (
		// Contoh username dan password
		username = "username"
		password = "password"

		// Inisialisasi file upload
		bucketName = "chum-bucket"
		fileName   = "notes.txt"
		filePath   = "external/files/"
	)

	// Membaca file JSON yang berisi konfigurasi
	config, err := configs.ReadConfigFromFile()
	if err != nil {
		log.Fatalf("failed to read config: %s", err)
	}

	// Menjalankan layanan autentikasi
	service := middleware.NewAuthService()
	if err := service.CheckBasicAuth(username, password); err != nil {
		log.Fatalf("gagal user tidak terautentikasi: %s", err)
	}

	// Buat instance untuk upload service
	uploadService := usecase.NewUploadService(config)
	result, err := uploadService.UploadFile(filePath, bucketName, fileName)
	if err != nil {
		log.Fatalf("failed to upload file: %s", err)
	}

	fmt.Println("File successfully uploaded.")
	fmt.Println("ETag:", result.Etag)
	fmt.Println("Public URL:", result.PublicURL)
}
