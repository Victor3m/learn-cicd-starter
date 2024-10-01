package database

import (
	"context"
	"database/sql"
	"testing"
)

func TestNew(t *testing.T) {
	db := &mockDB{}
	queries := New(db)
	if queries.db != db {
		t.Errorf("expected db to be set to %v, got %v", db, queries.db)
	}
}

func TestQueries_WithTx(t *testing.T) {
	db := &mockDB{}
	queries := New(db)
	tx := &sql.Tx{}
	queries = queries.WithTx(tx)
	if queries.db != tx {
		t.Errorf("expected db to be set to %v, got %v", tx, queries.db)
	}
}

type mockDB struct{}

func (m *mockDB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return nil, nil
}

func (m *mockDB) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return nil, nil
}

func (m *mockDB) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return nil, nil
}

func (m *mockDB) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return nil
}
