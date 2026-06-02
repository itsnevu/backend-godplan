package utils

import (
	"testing"

	"github.com/google/uuid"
)

func TestJWTUtil(t *testing.T) {
	jwtUtil := NewJWTUtil("test-secret-key")

	// Test GenerateToken
	testUserID := uuid.New()
	testTenantID := uuid.New()
	token, err := jwtUtil.GenerateToken(testUserID, "test@example.com", "employee", testTenantID)
	if err != nil {
		t.Fatalf("GenerateToken failed: %v", err)
	}
	if token == "" {
		t.Error("Generated token is empty")
	}

	// Test ValidateToken
	claims, err := jwtUtil.ValidateToken(token)
	if err != nil {
		t.Fatalf("ValidateToken failed: %v", err)
	}
	if claims.UserID != testUserID {
		t.Errorf("Expected UserID %s, got %s", testUserID.String(), claims.UserID.String())
	}
	if claims.Email != "test@example.com" {
		t.Errorf("Expected Email test@example.com, got %s", claims.Email)
	}
	if claims.Role != "employee" {
		t.Errorf("Expected Role employee, got %s", claims.Role)
	}
}

func TestJWTUtil_InvalidToken(t *testing.T) {
	jwtUtil := NewJWTUtil("test-secret-key")

	// Test with invalid token
	_, err := jwtUtil.ValidateToken("invalid-token")
	if err == nil {
		t.Error("Expected error for invalid token, but got none")
	}

	// Test with different secret key
	otherJWTUtil := NewJWTUtil("different-secret-key")
	testUserID := uuid.New()
	testTenantID := uuid.New()
	token, _ := jwtUtil.GenerateToken(testUserID, "test@example.com", "employee", testTenantID)
	_, err = otherJWTUtil.ValidateToken(token)
	if err == nil {
		t.Error("Expected error for token with different secret, but got none")
	}
}
