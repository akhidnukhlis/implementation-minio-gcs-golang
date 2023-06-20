package configs

import (
	"gopkg.in/yaml.v3"
	"implementasi-minio-gcs-golang/helpers"
	"log"
	"os"
	"path/filepath"
)

// Config merepresentasikan credential yang konfigurasi
type Config struct {
	GCSEnabled           bool   `yaml:"GCS_ENABLED"`
	GCSCredentialsFile   string `yaml:"GCS_CREDENTIALS_FILE"`
	MinioEnabled         bool   `yaml:"MINIO_ENABLED"`
	MinioEndpoint        string `yaml:"MINIO_ENDPOINT"`
	MinioAccessKey       string `yaml:"MINIO_ACCESS_KEY"`
	MinioSecretAccessKey string `yaml:"MINIO_SECRET_ACCESS_KEY"`
	MinioSSL             bool   `yaml:"MINIO_SSL"`
}

// Membaca file JSON konfigurasi credential
func ReadConfigFromFile() (*Config, error) {
	var credentials Config

	dir := helpers.GetCurrentDirectory()
	filePath := filepath.Join(dir, "configs/config.yml")

	// Membaca file YAML
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("gagal membaca file YAML: %s", err)
	}

	// Parsing data YAML
	err = yaml.Unmarshal(yamlFile, &credentials)
	if err != nil {
		log.Fatalf("gagal membaca data file YAML: %s", err)
	}

	return &credentials, nil
}
