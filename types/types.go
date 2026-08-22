package types

import (
	"go_code_challenge/utils"
	"time"
)

type Device struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	Brand        string    `json:"brand"`
	State        int       `json:"state"`
	CreationTime time.Time `json:"creation_time"`
}

type CreateDevicePayload struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Brand string `json:"brand"`
	State int    `json:"state"`
}

type UpdateDevicePayload struct {
	Id    int     `json:"id"`
	Name  *string `json:"name"`
	Brand *string `json:"brand"`
	State *int    `json:"state"`
}

type DeleteDevicePayload struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Brand string `json:"brand"`
}

type DeviceStore interface {
	CreateDevice(device CreateDevicePayload) error
	GetDevices() ([]*Device, error)
	GetDeviceById(id int) (*Device, error)
	GetDevicesByBrand(brand string) ([]*Device, error)
	GetDevicesByState(state utils.DeviceState) ([]*Device, error)
	UpdateDevice(device UpdateDevicePayload) error
	DeleteDevice(id int) error
}
