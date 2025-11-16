package domain

import "strings"

// UserData はキーを正規化した任意メタデータを保持する。
type UserData map[string]interface{}

// NewUserData はキーをトリムし、空キーがあれば ErrInvalidUserDataKey を返す。
func NewUserData(raw map[string]interface{}) (UserData, error) {
	if raw == nil {
		return UserData{}, nil
	}

	copied := make(UserData, len(raw))
	for key, value := range raw {
		trimmed := strings.TrimSpace(key)
		if trimmed == "" {
			return nil, ErrInvalidUserDataKey
		}

		copied[trimmed] = value
	}

	return copied, nil
}

// Clone は外部から改変されないよう浅いコピーを返す。
func (u UserData) Clone() UserData {
	if u == nil {
		return nil
	}

	copied := make(UserData, len(u))
	for key, value := range u {
		copied[key] = value
	}

	return copied
}
