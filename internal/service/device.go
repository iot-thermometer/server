package service

type Device interface {
}

type device struct {
}

func newDeviceService() Device {
	return &device{}
}
