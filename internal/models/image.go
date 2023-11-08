package models

import "time"

type Image struct {
	Path       string
	ID         string
	FolderId   string
	Name       string
	Size       int64
	Height     int
	Width      int
	UploadedAt time.Time
}
