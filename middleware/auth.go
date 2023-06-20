package middleware

import "fmt"

type AuthService struct {
}

// NewAuthService membuat instansi baru dari auth service
func NewAuthService() *AuthService {
	return &AuthService{}
}

// CheckBasicAuth periksa apakah username dan password valid
func (auth *AuthService) CheckBasicAuth(username, password string) error {
	var (
		// Lakukan logika autentikasi di sini
		// Contoh mekanisme autentikasi
		validUsername = "username"
		validPassword = "password"
	)

	if username != validUsername {
		message := "username tidak valid"

		return fmt.Errorf(message)
	}

	if password != validPassword {
		message := "password tidak valid"

		return fmt.Errorf(message)
	}

	return nil
}
