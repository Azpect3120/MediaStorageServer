package models

type Report struct {
	FolderName string
	CreatedAt  string
	MediaCount int
	Media      []*MediaData
}

type MediaData struct {
	Name          string
	Format        string
	Size          float64
	UploadedAt    string
}
