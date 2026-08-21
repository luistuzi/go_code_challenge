package service

import (
	"go_code_challenge/internal/model"
	"go_code_challenge/internal/repository"
	"slog"
	"uuid"
)

type Service struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s Service) GetAll() ([]model.Device, error) {
	devices, error := s.repo.Device.GetAll()
	if error != nil {
		return nil, slog.Error("service failed to get user: %w", error)
	}
	return devices, nil
}

func (s Service) Add(newDevice model.Device) (*model.Device, error) {
	repoRequest := model.Device{
		Id:           uuid.New(),
		Name:         newDevice.Name,
		Brand:        newDevice.Brand,
		State:        newDevice.State,
		CreationDate: newDevice.CreationDate,
	}

	s.repo.Device.Add(repoRequest)
	return &repoRequest, nil

}
