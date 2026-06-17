package exportdata

type ExportClusterFormsXLSXResponse struct {
	FileBytes []byte `json:"file_bytes"`
	FileName  string `json:"file_name"`
}
