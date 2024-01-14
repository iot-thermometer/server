package dto

type TopicMessage struct {
	DeviceID uint    `json:"device_id"`
	Value    float64 `json:"value"`
	Type     string  `json:"type"`
	Time     int64   `json:"time"`
}
