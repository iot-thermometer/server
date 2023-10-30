package dto

type TopicMessage struct {
	DeviceID     uint    `json:"device_id"`
	SoilMoisture float64 `json:"soil_moisture"`
	Temperature  float64 `json:"temperature"`
	Time         int64   `json:"time"`
}
