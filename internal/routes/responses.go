package routes

import "github.com/Azpect3120/MediaStorageServer/internal/models"

type GetFolderResponse struct {
	Status int             `json:"status"`
	Folder *models.Folder  `json:"folder"`
	Images []*models.Image `josn:"images"`
	Count  int             `json:"count"`
}

type GetImageResponse struct {
	Status int           `json:"status"`
	Image  *models.Image `json:"image"`
}
