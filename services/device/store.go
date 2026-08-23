package device

import (
	"database/sql"
	"fmt"
	"go_code_challenge/types"
	"strings"
	"time"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetDevices() ([]*types.Device, error) {
	rows, err := s.db.Query("SELECT * FROM DEVICE")

	if err != nil {
		return nil, err
	}

	devices := make([]*types.Device, 0)
	for rows.Next() {
		d, err := scanRowIntoDevice(rows)
		if err != nil {
			return nil, err
		}
		devices = append(devices, d)
	}

	return devices, nil
}

func (s *Store) GetDeviceById(id int) (*types.Device, error) {
	rows, err := s.db.Query("SELECT * FROM DEVICE WHERE ID = ? ", id)

	if err != nil {
		return nil, err
	}

	d := new(types.Device)
	for rows.Next() {
		d, err = scanRowIntoDevice(rows)
		if err != nil {
			return nil, err
		}
	}

	if d.Id == 0 {
		return nil, fmt.Errorf("Device not Found")
	}

	return d, nil
}

func (s *Store) GetDevicesByBrand(brand string) ([]*types.Device, error) {
	rows, err := s.db.Query("SELECT * FROM DEVICE WHERE BRAND = ? ", brand)

	if err != nil {
		return nil, err
	}

	devices := make([]*types.Device, 0)
	for rows.Next() {
		d, err := scanRowIntoDevice(rows)
		if err != nil {
			return nil, err
		}

		devices = append(devices, d)
	}

	if len(devices) == 0 {
		return nil, fmt.Errorf("Devices not Found")
	}

	return devices, nil
}

func (s *Store) GetDevicesByState(state string) ([]*types.Device, error) {

	rows, err := s.db.Query("SELECT * FROM DEVICE WHERE STATE = ? ", state)

	if err != nil {
		return nil, err
	}

	devices := make([]*types.Device, 0)
	for rows.Next() {
		d, err := scanRowIntoDevice(rows)
		if err != nil {
			return nil, err
		}

		devices = append(devices, d)
	}

	if len(devices) == 0 {
		return nil, fmt.Errorf("Devices not Found")
	}

	return devices, nil
}

func (s *Store) CreateDevice(device types.CreateDevicePayload) error {
	_, err := s.db.Exec("INSERT INTO DEVICE (ID, NAME, BRAND, STATE, CREATIONTIME) VALUES (?, ?, ?, ?, ?)", device.Id, device.Name, device.Brand, device.State, time.Now())
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateDevice(device types.UpdateDevicePayload, checker bool) error {

	query := "UPDATE DEVICE SET "

	args := []interface{}{}
	fields := []string{}

	if device.Name != nil {
		if checker {
			return fmt.Errorf("Cannot update device name while active state")
		}
		fields = append(fields, "NAME = ?")
		args = append(args, device.Name)
	}

	if device.Brand != nil {
		if checker {
			return fmt.Errorf("Cannot update device brand while active state")
		}
		fields = append(fields, "BRAND = ?")
		args = append(args, device.Brand)
	}

	if device.State != nil {
		fields = append(fields, "STATE = ?")
		args = append(args, device.State)
	}

	if len(fields) == 0 {
		return nil
	}

	query += strings.Join(fields, ", ")
	query += " WHERE ID = ? "
	args = append(args, device.Id)

	_, err := s.db.Exec(query, args...)

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeleteDevice(id int) error {

	_, err := s.db.Exec("DELETE FROM DEVICE WHERE ID = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func scanRowIntoDevice(rows *sql.Rows) (*types.Device, error) {
	device := new(types.Device)
	err := rows.Scan(
		&device.Id,
		&device.Name,
		&device.Brand,
		&device.State,
		&device.CreationTime,
	)

	if err != nil {
		return nil, err
	}

	return device, nil
}
