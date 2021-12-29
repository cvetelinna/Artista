package domain

type Image struct {
	ID       string
	FolderID string
	UserID   string
	Blob     []byte
}
