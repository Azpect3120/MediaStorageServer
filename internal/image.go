package internal

import "time"

type Image struct {
  path         string
  folderId     string
  name         string
  size         int64
  height       int
  width        int
  uploadedAt   time.Time
}
