package domain

type Image struct {
	ID       string `json:"id"`
	FolderID string `json:"folder_id"`
	UserID   string `json:"user_id"`
	Blob     []byte `json:"blob"`
}
