package controllers

type registerRequest struct {
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
}

type registerResponse struct {
	APIKey *string `json:"apikey,omitempty"`
}

type loginRequest struct {
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
}

type loginResponse struct {
	Message *string `json:"message,omitempty"`
}

type apiKeyRequest struct {
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
}

type apiKeyResponse struct {
	APIKey *string `json:"apikey,omitempty"`
}
