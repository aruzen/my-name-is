package domain

import "strings"

// Name トリム済み非空の名前
type Name struct {
	value string
}

// NewName 空文字であれば ErrEmptyName を返す。
func NewName(value string) (Name, error) {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return Name{}, ErrEmptyName
	}

	return Name{value: trimmed}, nil
}

func (n Name) String() string {
	return n.value
}
