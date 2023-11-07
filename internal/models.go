package internal

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

type Folder struct {
	id        string
	name      string
	createdAt time.Time
}
