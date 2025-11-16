package domain

import (
	"strings"

	"github.com/google/uuid"
)

// Token は UUID 形式の管理者トークンを保持する。
type Token struct {
	value uuid.UUID
}

// NewToken は UUID の書式を検証し、不正なら ErrInvalidToken を返す。
func NewToken(value string) (Token, error) {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return Token{}, ErrInvalidToken
	}

	parsed, err := uuid.Parse(trimmed)
	if err != nil {
		return Token{}, ErrInvalidToken
	}

	return Token{value: parsed}, nil
}

// UUID は内部で保持している uuid.UUID 値を返す。
func (t Token) UUID() uuid.UUID {
	return t.value
}

func (t Token) String() string {
	return t.value.String()
}
