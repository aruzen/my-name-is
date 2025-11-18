package api

import "backend/internal/domain"

type SignInRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r SignInRequest) ToDomain() (domain.SignInCredential, error) {
	return domain.NewSignInCredential(r.Name, r.Email, r.Password)
}

type SignInResponse struct {
	SessionPayload
}

func NewSignInResponse(session domain.SessionData) SignInResponse {
	return SignInResponse{SessionPayload: NewSessionPayload(session)}
}
