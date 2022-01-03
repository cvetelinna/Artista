package folders

import (
	"Artista/domain"
	"context"
	"database/sql"
)

const insertFolderQuery = `INSERT INTO "public".folders ("name", "user_id") VALUES ($1, $2);`

type pgFoldersRepo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *pgFoldersRepo {
	return &pgFoldersRepo{db: db}
}

func (r *pgFoldersRepo) Insert(ctx context.Context, folder *domain.Folder) error {
	_, err := r.db.ExecContext(ctx, insertFolderQuery,
		folder.Name, folder.UserID)
	return err
}
