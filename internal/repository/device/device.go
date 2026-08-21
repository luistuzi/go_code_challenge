package device

import (
	"go_code_challenge/internal/model"

	"github.com/google/uuid"
)

type Device struct {
	device []model.Device
}

func New() *Device {
	return &Device{device: make([]model.Device, 0)}
}

func (d *Device) GetAll() ([]model.Device, error) {
	return d.device, nil
}

func (d *Device) GetById(id uuid.UUID) (*model.Device, error) {
	return nil, nil
}

func (d *Device) GetByBrand(brand string) ([]Device, error) {
	return nil, nil
}

func (d *Device) GetByState(state string) ([]model.Device, error) {
	return nil, nil
}

func (d *Device) Add(newDevice model.Device) (*model.Device, error) {
	return d.device.append(d.device, newDevice), nil
}

func (d *Device) Update(device Device) (*Device, error) {
	return nil, nil
}

func (d *Device) Delete(id uuid.UUID, device Device) (*model.Device, error) {
	return nil, nil
}
