package model

import (
	"time"

	"github.com/iot-thermometer/server/gen"
	"github.com/iot-thermometer/server/internal/util"
	"gorm.io/gorm"
)

type Reading struct {
	gorm.Model
	DeviceID     uint
	SoilMoisture float64 `gorm:"not null"`
	Temperature  float64 `gorm:"not null"`
	Other        float64 `gorm:"not null"`
	MeasuredAt   time.Time
	UploadedAt   time.Time
}

func (r Reading) Protobuf() *gen.Reading {
	return &gen.Reading{
		Id:           util.Int64(int64(r.ID)),
		DeviceId:     util.Int64(int64(r.DeviceID)),
		SoilMoisture: util.Float32(float32(r.SoilMoisture)),
		Temperature:  util.Float32(float32(r.Temperature)),
		Other:        util.Float32(float32(r.Other)),
		MeasuredAt:   util.Int64(r.MeasuredAt.Unix()),
		UploadedAt:   util.Int64(r.UploadedAt.Unix()),
	}
}
