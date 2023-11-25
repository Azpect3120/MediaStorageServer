package models

type FolderChannel struct {
	Folder *Folder
	Error  error
}

type ImageChannel struct {
	Image *Image
	Error error
}

type ImagesChannel struct {
	Images []*Image
	Error  error
}

type IDChannel struct {
	ID    string
	Error error
}

type ReportChannel struct {
	Report Report
	Error  error
}
