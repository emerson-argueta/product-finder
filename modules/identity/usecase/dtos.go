package usecase

// RegisterDTO for register usecase
type RegisterDTO struct {
	Email    string
	Password string
}

// LoginDTO for login usecase
type LoginDTO struct {
	Email    string
	Password string
}

// APIKeyDTO for api key usecase
type APIKeyDTO struct {
	Email    string
	Password string
}
