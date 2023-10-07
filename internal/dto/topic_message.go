package dto

type TopicMessage struct {
	DeviceID uint    `json:"device_id"`
	Type     string  `json:"type"`
	Value    float64 `json:"value"`
	Time     int64   `json:"time"`
}
