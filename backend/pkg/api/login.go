package api

import "backend/internal/domain"

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (r LoginRequest) ToDomain() (domain.AdminCredential, error) {
	name, err := domain.NewName(r.Name)
	if err != nil {
		return domain.AdminCredential{}, err
	}

	return domain.NewAdminCredential(name, r.Password)
}

type LoginResponse struct {
	Token string `json:"token"`
}

func NewLoginResponse(token domain.Token) LoginResponse {
	return LoginResponse{Token: token.String()}
}
