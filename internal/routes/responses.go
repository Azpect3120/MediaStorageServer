package routes

import "github.com/Azpect3120/MediaStorageServer/internal/models"

type GetFolderResponse struct {
	Status int            `json:"status"`
	Folder *models.Folder `json:"folder"`
}

type GetFolderImagesResponse struct {
	Status int             `json:"status"`
	Images []*models.Image `json:"images"`
}

type GetImageResponse struct {
	Status int           `json:"status"`
	Image  *models.Image `json:"image"`
}
