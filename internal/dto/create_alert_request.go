package dto

type CreateAlertRequest struct {
	DeviceID        uint    `json:"device_id"`
	Name            string  `json:"name"`
	SoilMoistureMin float64 `json:"soil_moisture_min"`
	SoilMoistureMax float64 `json:"soil_moisture_max"`
	TemperatureMin  float64 `json:"temperature_min"`
	TemperatureMax  float64 `json:"temperature_max"`
}
