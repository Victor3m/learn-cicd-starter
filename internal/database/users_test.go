package database

import (
	"context"
	"testing"
)

func TestCreateUser(t *testing.T) {
	db := &mockDB{}
	queries := New(db)
	user := CreateUserParams{
		ID:        "123",
		CreatedAt: "2022-01-01",
		UpdatedAt: "2022-01-01",
		Name:      "John Doe",
		ApiKey:    "abc123",
	}
	err := queries.CreateUser(context.Background(), user)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestGetUser(t *testing.T) {
	db := &mockDB{}
	queries := New(db)
	apiKey := "abc123"
	_, err := queries.GetUser(context.Background(), apiKey)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
