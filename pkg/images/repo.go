package images

import (
	"Artista/domain"
	"context"
	"database/sql"
)

const insertImageQuery = `INSERT INTO "public"."images" ("folder_id", "user_id", "blob" ) VALUES ($1, $2, $3);`

type pgImagesRepo struct {
	db *sql.DB
}

func newPgImagesRepo(db *sql.DB) *pgImagesRepo {
	return &pgImagesRepo{db: db}
}

func (r *pgImagesRepo) Insert(ctx context.Context, image *domain.Image) error {
	_, err := r.db.ExecContext(ctx, insertImageQuery,
		image.FolderID, image.UserID, image.Blob)
	return err
}
