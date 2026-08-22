package device

import (
	"go_code_challenge/types"
)

type mockDeviceStore struct{}

func (m *mockDeviceStore) CreateDevice(d types.CreateDevicePayload) error {
	return nil
}
