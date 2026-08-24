package device

import (
	"database/sql"
	"fmt"
	"go_code_challenge/types"
	"go_code_challenge/utils"
	"log"
	"strings"
	"time"
)

type Service struct {
	db *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{db: db}
}

func (s *Service) GetDevices() ([]*types.Device, error) {

	rows, err := s.db.Query("SELECT * FROM DEVICE")

	if err != nil {
		return nil, err
	}

	log.Println(utils.SelectExecuted)

	devices := make([]*types.Device, 0)
	for rows.Next() {
		d, err := scanRowIntoDevice(rows)
		if err != nil {
			return nil, err
		}
		log.Println("Current row: ", d)
		devices = append(devices, d)
	}

	return devices, nil

}

func (s *Service) GetDeviceById(id int) (*types.Device, error) {

	rows, err := s.db.Query("SELECT * FROM DEVICE WHERE ID = ? ", id)

	if err != nil {
		return nil, err
	}

	log.Println(utils.SelectExecuted)

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

func (s *Service) GetDevicesByBrand(brand string) ([]*types.Device, error) {

	rows, err := s.db.Query("SELECT * FROM DEVICE WHERE BRAND = ? ", brand)

	if err != nil {
		return nil, err
	}

	log.Println(utils.SelectExecuted)

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

func (s *Service) GetDevicesByState(state string) ([]*types.Device, error) {

	rows, err := s.db.Query("SELECT * FROM DEVICE WHERE STATE = ? ", state)

	if err != nil {
		return nil, err
	}

	log.Println(utils.SelectExecuted)

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

func (s *Service) CreateDevice(device types.CreateDevicePayload) error {

	_, err := s.db.Exec("INSERT INTO DEVICE (ID, NAME, BRAND, STATE, CREATIONTIME) VALUES (?, ?, ?, ?, ?)", device.Id, device.Name, device.Brand, device.State, time.Now())
	if err != nil {
		return err
	}

	return nil

}

func (s *Service) UpdateDevice(device types.UpdateDevicePayload, checker bool) error {

	query := "UPDATE DEVICE SET "

	args := []interface{}{}
	fields := []string{}

	if device.Name != nil {
		if checker {
			return fmt.Errorf("Cannot update device name while in-use state")
		}
		fields = append(fields, "NAME = ?")
		args = append(args, device.Name)
	}

	if device.Brand != nil {
		if checker {
			return fmt.Errorf("Cannot update device brand while in-use state")
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

func (s *Service) DeleteDevice(id int) error {

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
