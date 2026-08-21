package repository

import (
	"go_code_challenge/internal/model"
	"go_code_challenge/internal/repository/device"

	"github.com/google/uuid"
)

type Repository struct {
	Device interface {
		GetAll() ([]model.Device, error)
		GetById(id uuid.UUID) (*model.Device, error)
		GetByBrand(brand string) ([]model.Device, error)
		GetByState(state int) ([]model.Device, error)
		Add(device model.Device) (*model.Device, error)
		Update(device model.Device) (*model.Device, error)
		Delete(id uuid.UUID) error
	}
}

func New() *Repository {
	return &Repository{
		Device: device.New(),
	}
}
