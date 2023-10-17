package dto

type FirmwareIndexResponse struct {
	Version int    `json:"version"`
	Source  string `json:"source"`
}
