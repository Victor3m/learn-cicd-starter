package database

import (
	"context"
	"testing"
)

func TestCreateNote(t *testing.T) {
	db := &mockDB{}
	queries := New(db)
	note := CreateNoteParams{
		ID:        "123",
		CreatedAt: "2022-01-01",
		UpdatedAt: "2022-01-01",
		Note:      "Hello World",
		UserID:    "123",
	}
	err := queries.CreateNote(context.Background(), note)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestGetNote(t *testing.T) {
	db := &mockDB{}
	queries := New(db)
	note := Note{
		ID: "123",
	}
	_, err := queries.GetNote(context.Background(), note.ID)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestGetNotesForUser(t *testing.T) {
	db := &mockDB{}
	queries := New(db)
	userID := "123"
	notes, err := queries.GetNotesForUser(context.Background(), userID)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(notes) != 0 {
		t.Errorf("expected no notes, got %v", notes)
	}
}
