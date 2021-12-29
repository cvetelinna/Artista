package users

import (
	"Artista/domain"
	"context"
	"database/sql"
)

const insertUserQuery = `INSERT INTO "public"."users" ("username", "password", "role") VALUES ($1, $2, $3);`

type pgUserRepo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *pgUserRepo {
	return &pgUserRepo{db: db}
}

func (r *pgUserRepo) Insert(ctx context.Context, user *domain.User) error {
	_, err := r.db.ExecContext(ctx, insertUserQuery,
		user.Username, user.Password, user.Role)
	return err
}
