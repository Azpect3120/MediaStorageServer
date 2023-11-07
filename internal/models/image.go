package models

import "time"

type Image struct {
	path       string
	id         string
	folderId   string
	name       string
	size       int64
	height     int
	width      int
	uploadedAt time.Time
}
