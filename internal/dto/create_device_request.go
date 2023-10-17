package dto

type CreateDeviceRequest struct {
	Name            string `json:"name" validate:"required"`
	ReadingInterval int    `json:"reading_interval" validate:"required"`
	PushInterval    int    `json:"push_interval" validate:"required"`
}
