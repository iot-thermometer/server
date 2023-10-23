package procedure

import "github.com/iot-thermometer/server/internal/service"

type Device interface {
}

type deviceProcedure struct {
	deviceService service.Device
}

func newDeviceProcedure(deviceService service.Device) Device {
	return &deviceProcedure{
		deviceService: deviceService,
	}
}
