package reports

import (
	"bytes"
	"html/template"
	"math"

	"github.com/Azpect3120/MediaStorageServer/internal/database"
	"github.com/Azpect3120/MediaStorageServer/internal/models"
)

// Generate a report using the folder id
func Generate(db *database.Database, id string) (*models.Report, error) {
	var report models.Report

	folder, err := db.GetFolder(id)

	if err != nil {
		return &report, err
	}

	report.FolderName = folder.Name
	report.CreatedAt = folder.CreatedAt.Format("Mon, Jan 2, 2006 at 3:04pm")

	images, err := db.GetImages(id, math.MaxInt, 1)

	if err != nil {
		return &report, err
	}

	report.MediaCount = len(images)

	var media []*models.MediaData
	for _, image := range images{
		data := &models.MediaData{Name: image.Name, Format: image.Format, Size: float64(image.Size), UploadedAt: image.UploadedAt.Format("Mon, Jan 2, 2006 at 3:04pm")}
		media = append(media, data)
	}

	report.Media = media

	return &report, nil
}

// Convert the report to a string that can be emailed to the user
func String(r *models.Report) (string, error) {
	for i := range r.Media {
		if r.Media[i].Size >= 1000000000 {
			r.Media[i].Size = math.Round(r.Media[i].Size / float64(1000000000) * 10) / 10
			r.Media[i].SizeSuffix = "gb"
		} else if r.Media[i].Size >= 1000000 {
			r.Media[i].Size = math.Round(r.Media[i].Size / float64(1000000) * 10) / 10
			r.Media[i].SizeSuffix = "mb"
		} else if r.Media[i].Size >= 1000 {
			r.Media[i].Size = math.Round(r.Media[i].Size / float64(1000) * 10) / 10
			r.Media[i].SizeSuffix = "kb"
		} else {
			r.Media[i].SizeSuffix = "b"
		}
	}

	const emailTemplate string = `
		Hello,

		Here is your folder report:

		Folder Name: {{.FolderName}}
		Created At: {{.CreatedAt}}
		Media Count: {{.MediaCount}}

		Media Details:
		{{range .Media}}
		  Name: {{.Name}}
		  Format: {{.Format}}
		  Size: {{.Size}}{{.SizeSuffix}}
		  Uploaded At: {{.UploadedAt}}
		{{end}}

		For API details, usage, and information visit https://github.com/Azpect3120/MediaStorageServer

		Best regards,
		Media Storage Server
	`

	tmpl, err := template.New("emailTemplate").Parse(emailTemplate)
	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer

	if err := tmpl.Execute(&buffer, r); err != nil {
		return "", nil
	}

	return buffer.String(), nil
}
