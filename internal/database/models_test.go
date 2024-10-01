package database

import (
	"testing"
)

func TestNote(t *testing.T) {
	note := Note{
		ID:        "123",
		CreatedAt: "2022-01-01",
		UpdatedAt: "2022-01-01",
		Note:      "Hello World",
		UserID:    "123",
	}
	if note.ID != "123" {
		t.Errorf("expected ID to be %v, got %v", "123", note.ID)
	}
	if note.CreatedAt != "2022-01-01" {
		t.Errorf("expected CreatedAt to be %v, got %v", "2022-01-01", note.CreatedAt)
	}
	if note.UpdatedAt != "2022-01-01" {
		t.Errorf("expected UpdatedAt to be %v, got %v", "2022-01-01", note.UpdatedAt)
	}
	if note.Note != "Hello World" {
		t.Errorf("expected Note to be %v, got %v", "Hello World", note.Note)
	}
	if note.UserID != "123" {
		t.Errorf("expected UserID to be %v, got %v", "123", note.UserID)
	}
}

func TestUser(t *testing.T) {
	user := User{
		ID:        "123",
		CreatedAt: "2022-01-01",
		UpdatedAt: "2022-01-01",
		Name:      "John Doe",
		ApiKey:    "abc123",
	}
	if user.ID != "123" {
		t.Errorf("expected ID to be %v, got %v", "123", user.ID)
	}
	if user.CreatedAt != "2022-01-01" {
		t.Errorf("expected CreatedAt to be %v, got %v", "2022-01-01", user.CreatedAt)
	}
	if user.UpdatedAt != "2022-01-01" {
		t.Errorf("expected UpdatedAt to be %v, got %v", "2022-01-01", user.UpdatedAt)
	}
	if user.Name != "John Doe" {
		t.Errorf("expected Name to be %v, got %v", "John Doe", user.Name)
	}
	if user.ApiKey != "abc123" {
		t.Errorf("expected ApiKey to be %v, got %v", "abc123", user.ApiKey)
	}
}
