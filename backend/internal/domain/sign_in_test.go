package domain

import "testing"

func TestNewSignInCredential(t *testing.T) {
	credential, err := NewSignInCredential(" Alice ", "alice@example.com", "  secret  ")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if credential.Name().String() != "Alice" {
		t.Fatalf("expected name Alice, got %s", credential.Name())
	}

	if credential.Email().String() != "alice@example.com" {
		t.Fatalf("unexpected email: %s", credential.Email())
	}

	if credential.Password() != "secret" {
		t.Fatalf("expected trimmed password, got %s", credential.Password())
	}
}

func TestNewSignInCredential_InvalidInputs(t *testing.T) {
	if _, err := NewSignInCredential("", "alice@example.com", "secret"); err == nil {
		t.Fatalf("expected error for empty name")
	}

	if _, err := NewSignInCredential("Alice", "bad-email", "secret"); err == nil {
		t.Fatalf("expected error for invalid email")
	}

	if _, err := NewSignInCredential("Alice", "alice@example.com", "   "); err != ErrInvalidPassword {
		t.Fatalf("expected ErrInvalidPassword, got %v", err)
	}
}
