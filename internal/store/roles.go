package store

import (
	"context"
	"database/sql"
)

type Role struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Level int    `json:"level"`
}

type RoleStore struct {
	db *sql.DB
}

func (s *RoleStore) GetByName(ctx context.Context, roleName string) (*Role, error) {
	query := `
		SELECT id, name, level FROM roles WHERE name = $1
	`

	role := &Role{}

	err := s.db.QueryRowContext(ctx, query, roleName).Scan(
		&role.ID,
		&role.Name,
		&role.Level,
	)

	if err != nil {
		return nil, err
	}

	return role, nil
}
