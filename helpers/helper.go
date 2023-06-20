package helpers

import (
	"log"
	"os"
)

// GetCurrentDirectory untuk melihat direktori saat ini
func GetCurrentDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("gagal meilhat direktori saat ini: %s", err)
	}

	return dir
}
