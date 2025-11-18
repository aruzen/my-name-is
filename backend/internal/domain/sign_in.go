package domain

import "strings"

// SignInCredential はサインインAPIの入力値を保持する。
type SignInCredential struct {
	name     Name
	email    Email
	password string
}

// NewSignInCredential は name/email/password を検証して正規化する。
func NewSignInCredential(name, email, password string) (SignInCredential, error) {
	trimmedPassword := strings.TrimSpace(password)
	if trimmedPassword == "" {
		return SignInCredential{}, ErrInvalidPassword
	}

	parsedName, err := NewName(name)
	if err != nil {
		return SignInCredential{}, err
	}

	parsedEmail, err := NewEmail(email)
	if err != nil {
		return SignInCredential{}, err
	}

	return SignInCredential{
		name:     parsedName,
		email:    parsedEmail,
		password: trimmedPassword,
	}, nil
}

func (c SignInCredential) Name() Name {
	return c.name
}

func (c SignInCredential) Email() Email {
	return c.email
}

func (c SignInCredential) Password() string {
	return c.password
}
