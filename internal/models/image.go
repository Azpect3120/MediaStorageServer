package models

import "time"

type Image struct {
	ID         string
	FolderId   string
	Name       string
	Size       int64
	Format     string
	UploadedAt time.Time
	Path       string
}
