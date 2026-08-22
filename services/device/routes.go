package device

import (
	"fmt"
	"go_code_challenge/types"
	"go_code_challenge/utils"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Controller struct {
	store types.DeviceStore
}

func NewController(store types.DeviceStore) *Controller {
	return &Controller{store: store}
}

func (c *Controller) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/getDevices", c.GetDevicesController).Methods("GET")
	router.HandleFunc("/getDeviceByiD/{id}", c.GetDeviceByIdController).Methods("GET")
	router.HandleFunc("/getDeviceByBrand/{brand}", c.GetDevicesByBrandController).Methods("GET")
	router.HandleFunc("/getDeviceByState/{state}", c.GetDevicesByStateController).Methods("GET")
	router.HandleFunc("/createDevice", c.CreateDeviceController).Methods("POST")
	router.HandleFunc("/updateDevice", c.UpdateDeviceController).Methods("PATCH")
	router.HandleFunc("/deleteDevice/{id}", c.DeleteDeviceController).Methods("DELETE")
}

func (c *Controller) GetDevicesController(w http.ResponseWriter, r *http.Request) {

	devices, err := c.store.GetDevices()

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, devices)

}

func (c *Controller) GetDeviceByIdController(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	str, ok := vars["id"]

	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Missing the ID of the device"))
		return
	}

	id, err := strconv.Atoi(str)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	device, err := c.store.GetDeviceById(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, device)

}

func (c *Controller) GetDevicesByBrandController(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	str, ok := vars["brand"]

	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Missing the Brand of the device"))
		return
	}

	devices, err := c.store.GetDevicesByBrand(str)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, devices)

}

func (c *Controller) GetDevicesByStateController(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	str, ok := vars["state"]

	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Missing the State of the device"))
		return
	}

	stateValue, err := utils.ParseDeviceState(str)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	devices, err := c.store.GetDevicesByState(stateValue)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusAccepted, devices)

}

func (c *Controller) CreateDeviceController(w http.ResponseWriter, r *http.Request) {

	var device types.CreateDevicePayload

	if err := utils.ParseJson(r, &device); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(device); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid insert payload request: %v", errors))
	}

	err := c.store.CreateDevice(device)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusCreated, device)

}

func (c *Controller) UpdateDeviceController(w http.ResponseWriter, r *http.Request) {

	var device types.UpdateDevicePayload

	if err := utils.ParseJson(r, &device); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(device); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid update payload request: %v", errors))
	}

	err := c.store.UpdateDevice(device)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJson(w, http.StatusOK, device)

}

func (c *Controller) DeleteDeviceController(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	str, ok := vars["id"]

	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Missing the ID of the device"))
		return
	}

	id, err := strconv.Atoi(str)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = c.store.DeleteDevice(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, "device")

}
