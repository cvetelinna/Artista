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

type UserRepository interface {
	Insert(ctx context.Context, user *domain.User) error
	Fetch(ctx context.Context, user *domain.User) error
	Update(ctx context.Context, user *domain.User) (*domain.User, error)
}

func (r *pgUserRepo) Insert(ctx context.Context, user *domain.User) error {
	_, err := r.db.ExecContext(ctx, insertUserQuery,
		user.Username, user.Password, user.Role)
	return err
}

func (r *pgUserRepo) Fetch(ctx context.Context, user *domain.User) (*domain.User, error) {
	err := r.db.QueryRowContext(ctx, "SELECT username, password FROM users WHERE username=?", user.Username).Scan(&user.Username)
	return user, err
}

func (r *pgUserRepo) Update(ctx context.Context, user *domain.User) (*domain.User, error) {
	err := r.db.QueryRowContext(ctx, "UPDATE username, password FROM users WHERE id=?", user.ID).Scan(&user.Username, &user.Password)
	return user, err
}
