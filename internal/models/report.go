package models

import "time"

type Report struct {
	FolderName string
	CreatedAt  time.Time
	MediaCount int
	Media      []*MediaData
}

type MediaData struct {
	Name          string
	Format        string
	Size          float64
	UploadedAt    string
}
