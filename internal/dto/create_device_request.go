package dto

type CreateDeviceRequest struct {
	Name string `json:"name" validate:"required"`
}
