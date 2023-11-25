package reports

import (
	"bytes"
	"html/template"
	"math"

	"github.com/Azpect3120/MediaStorageServer/internal/database"
	"github.com/Azpect3120/MediaStorageServer/internal/models"
)

// Generate a report using the folder id
func Generate(ch chan models.ReportChannel, db *database.Database, id string) {
	var report models.Report

	chF := make(chan models.FolderChannel)
	go db.GetFolder(chF, id)
	resF := <-chF

	if resF.Error != nil {
		ch <- models.ReportChannel{Report: report, Error: resF.Error}
		return
	}

	report.FolderName = resF.Folder.Name
	report.CreatedAt = resF.Folder.CreatedAt.Format("Mon, Jan 2, 2006 at 3:04pm")

	chI := make(chan models.ImagesChannel)
	go db.GetImages(chI, id)
	resI := <-chI

	if resI.Error != nil {
		ch <- models.ReportChannel{Report: report, Error: resI.Error}
		return
	}

	report.MediaCount = len(resI.Images)

	var media []*models.MediaData
	for _, image := range resI.Images {
		data := &models.MediaData{Name: image.Name, Format: image.Format, Size: float64(image.Size), UploadedAt: image.UploadedAt.Format("Mon, Jan 2, 2006 at 3:04pm")}
		media = append(media, data)
	}

	report.Media = media

	ch <- models.ReportChannel{Report: report, Error: nil}
}

// Convert the report to a string that can be emailed to the user
func String(r *models.Report) (string, error) {
	for i := range r.Media {
		r.Media[i].Size = math.Round(r.Media[i].Size / float64(1000000) * 10) / 10

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
		  Size: {{.Size}}mb
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
