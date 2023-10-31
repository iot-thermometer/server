package dto

type TopicMessage struct {
	DeviceID     uint    `json:"device_id"`
	SoilMoisture float64 `json:"soil_moisture"`
	Temperature  float64 `json:"temperature"`
	Other        float64 `json:"other"`
	Time         int64   `json:"time"`
}
