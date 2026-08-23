package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"go_code_challenge/types"

	"github.com/gorilla/mux"
)

func TestDeviceControllers(t *testing.T) {

	deviceService := &mockDeviceService{}
	controller := NewController(deviceService)

	t.Run("should get all devices", func(t *testing.T) {

		req, err := http.NewRequest(http.MethodGet, "/getDevices", nil)

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/getDevices", controller.GetDevicesController).Methods(http.MethodGet)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Expected status code %d but got %d", http.StatusOK, rr.Code)
		}
	})

	t.Run("should get a device by id", func(t *testing.T) {

		req, err := http.NewRequest(http.MethodGet, "/getDeviceById/123", nil)

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/getDeviceById/{id}", controller.GetDeviceByIdController).Methods(http.MethodGet)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Expected status code %d but got %d", http.StatusOK, rr.Code)
		}
	})

	t.Run("should fail with non number id on get device", func(t *testing.T) {

		req, err := http.NewRequest(http.MethodGet, "/getDeviceById/abc", nil)

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/getDeviceById/{id}", controller.GetDeviceByIdController).Methods(http.MethodGet)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should get a device by brand", func(t *testing.T) {

		req, err := http.NewRequest(http.MethodGet, "/getDevicesByBrand/testbrand", nil)

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/getDevicesByBrand/{brand}", controller.GetDevicesByBrandController).Methods(http.MethodGet)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Expected status code %d but got %d", http.StatusOK, rr.Code)
		}
	})

	t.Run("should get a device by state", func(t *testing.T) {

		req, err := http.NewRequest(http.MethodGet, "/getDevicesByState/active", nil)

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/getDevicesByState/{state}", controller.GetDevicesByStateController).Methods(http.MethodGet)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Expected status code %d but got %d", http.StatusOK, rr.Code)
		}
	})

	t.Run("should create a device", func(t *testing.T) {
		payload := types.CreateDevicePayload{
			Id:    1,
			Name:  "device1",
			Brand: "brand 1",
			State: "active",
		}

		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/createDevice", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/createDevice", controller.CreateDeviceController).Methods(http.MethodPost)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, rr.Code)
		}
	})

	t.Run("should update a device", func(t *testing.T) {
		name := "device1"
		brand := "brand1"
		state := "active"

		payload := types.UpdateDevicePayload{
			Id:    1,
			Name:  &name,
			Brand: &brand,
			State: &state,
		}

		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/updateDevice", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/updateDevice", controller.UpdateDeviceController).Methods(http.MethodPost)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
		}
	})

	t.Run("should update only the name device", func(t *testing.T) {
		name := "device1"

		payload := types.UpdateDevicePayload{
			Id:    1,
			Name:  &name,
			Brand: nil,
			State: nil,
		}

		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/updateDevice", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/updateDevice", controller.UpdateDeviceController).Methods(http.MethodPost)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
		}
	})

	t.Run("should update only the brand device", func(t *testing.T) {
		brand := "brand1"

		payload := types.UpdateDevicePayload{
			Id:    1,
			Name:  nil,
			Brand: &brand,
			State: nil,
		}

		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/updateDevice", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/updateDevice", controller.UpdateDeviceController).Methods(http.MethodPost)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
		}
	})

	t.Run("should update only the state device", func(t *testing.T) {
		state := "active"

		payload := types.UpdateDevicePayload{
			Id:    1,
			Name:  nil,
			Brand: nil,
			State: &state,
		}

		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/updateDevice", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/updateDevice", controller.UpdateDeviceController).Methods(http.MethodPost)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
		}
	})

	t.Run("should delete a device", func(t *testing.T) {

		req, err := http.NewRequest(http.MethodGet, "/deleteDevice/123", nil)

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/deleteDevice/{id}", controller.DeleteDeviceController).Methods(http.MethodGet)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Expected status code %d but got %d", http.StatusOK, rr.Code)
		}
	})

	t.Run("should raise an error while deleting a device", func(t *testing.T) {

		req, err := http.NewRequest(http.MethodGet, "/deleteDevice/avsd", nil)

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/deleteDevice/{id}", controller.DeleteDeviceController).Methods(http.MethodGet)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, rr.Code)
		}
	})

}

type mockDeviceService struct{}

func (m *mockDeviceService) GetDevices() ([]*types.Device, error) {
	return []*types.Device{}, nil
}

func (m *mockDeviceService) GetDeviceById(id int) (*types.Device, error) {
	return &types.Device{}, nil
}

func (m *mockDeviceService) GetDevicesByBrand(brand string) ([]*types.Device, error) {
	return []*types.Device{}, nil
}

func (m *mockDeviceService) GetDevicesByState(state string) ([]*types.Device, error) {
	return []*types.Device{}, nil
}

func (m *mockDeviceService) CreateDevice(d types.CreateDevicePayload) error {
	return nil
}

func (m *mockDeviceService) UpdateDevice(d types.UpdateDevicePayload, checker bool) error {
	return nil
}

func (m *mockDeviceService) DeleteDevice(id int) error {
	return nil
}
