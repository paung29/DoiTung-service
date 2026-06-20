package exportdata

type ExportXLSXResponse struct {
	FileBytes []byte `json:"file_bytes"`
	FileName  string `json:"file_name"`
}
