package main

import (
	"fmt"
	"implementation-minio-gcs-golang/configs"
	"implementation-minio-gcs-golang/middleware"
	"implementation-minio-gcs-golang/usecase"
	"log"
)

func main() {
	var (
		// Contoh username dan password
		username = "username"
		password = "password"

		// Inisialisasi file upload dan download
		bucketName          = "chum-bucket"
		fileName            = "akhidnukhlis-notes.txt"
		filePath            = "external/files/"
		destinationFilePath = "external/destinations/"
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
	resultUpload, err := uploadService.UploadFile(fileName, bucketName, filePath)
	if err != nil {
		log.Fatalf("gagal upload file: %s", err)
	}

	// Buat instance untuk get service
	getService := usecase.NewGetService(config)
	resultRetrieve, err := getService.GetFile(fileName, bucketName, destinationFilePath)
	if err != nil {
		log.Fatalf("gagal mendapatkan file: %s", err)
	}

	fmt.Println("File successfully uploaded.")
	fmt.Println("ETag:", resultUpload.Etag)
	fmt.Println("Public URL:", resultUpload.PublicURL)

	fmt.Println("============================")
	fmt.Println("File successfully retrieve.")
	fmt.Println("Data:", resultRetrieve)
}
