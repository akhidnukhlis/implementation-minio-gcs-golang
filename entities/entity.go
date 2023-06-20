package entity

// UploadResult struct untuk mengelola hasil upload file
type UploadResult struct {
	Etag      string
	PublicURL string
}

// FileResult struct untuk menampung file
type FileResult struct {
	Source          string
	Bucket          string
	Object          string
	DestinationFile string
	Success         bool
	ErrorMessage    string
}
