package types

import (
	"time"
)

type Device struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	Brand        string    `json:"brand"`
	State        string    `json:"state"`
	CreationTime time.Time `json:"creation_time"`
}

type CreateDevicePayload struct {
	Id    int    `json:"id" validate:"required"`
	Name  string `json:"name" validate:"required"`
	Brand string `json:"brand" validate:"required"`
	State string `json:"state" validate:"required"`
}

type UpdateDevicePayload struct {
	Id    int     `json:"id" validate:"required"`
	Name  *string `json:"name"`
	Brand *string `json:"brand"`
	State *string `json:"state"`
}

type DeviceStore interface {
	CreateDevice(device CreateDevicePayload) error
	GetDevices() ([]*Device, error)
	GetDeviceById(id int) (*Device, error)
	GetDevicesByBrand(brand string) ([]*Device, error)
	GetDevicesByState(state string) ([]*Device, error)
	UpdateDevice(device UpdateDevicePayload, checker bool) error
	DeleteDevice(id int) error
}
