package device

import (
	"fmt"
	"go_code_challenge/types"
	"go_code_challenge/utils"
	"log"
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
	router.HandleFunc("/getDeviceById/{id}", c.GetDeviceByIdController).Methods("GET")
	router.HandleFunc("/getDevicesByBrand/{brand}", c.GetDevicesByBrandController).Methods("GET")
	router.HandleFunc("/getDevicesByState/{state}", c.GetDevicesByStateController).Methods("GET")
	router.HandleFunc("/createDevice", c.CreateDeviceController).Methods("POST")
	router.HandleFunc("/updateDevice", c.UpdateDeviceController).Methods("PATCH")
	router.HandleFunc("/deleteDevice/{id}", c.DeleteDeviceController).Methods("DELETE")
}

func (c *Controller) GetDevicesController(w http.ResponseWriter, r *http.Request) {
	log.Println("Starting get devices controller")

	devices, err := c.store.GetDevices()

	log.Println("Devices returned: ", devices)

	if err != nil {
		log.Println("Error occured: ", err)
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("Not expected error occured: %v", err))
		return
	}

	utils.WriteJson(w, http.StatusOK, devices)

}

func (c *Controller) GetDeviceByIdController(w http.ResponseWriter, r *http.Request) {

	log.Println("Starting get device by id controller")

	vars := mux.Vars(r)
	str, ok := vars["id"]

	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Missing the ID of the device"))
		return
	}

	id, err := strconv.Atoi(str)
	if err != nil {
		log.Println("Error occured: ", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid get payload request: %v", err))
		return
	}

	device, err := c.store.GetDeviceById(id)
	if err != nil {
		log.Println("Error occured: ", err)
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("Not expected error occured: %v", err))
		return
	}

	log.Println("Device returned: ", device)

	utils.WriteJson(w, http.StatusOK, device)

}

func (c *Controller) GetDevicesByBrandController(w http.ResponseWriter, r *http.Request) {

	log.Println("Starting get devices by brand controller")

	vars := mux.Vars(r)
	str, ok := vars["brand"]

	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Missing the Brand of the device"))
		return
	}

	devices, err := c.store.GetDevicesByBrand(str)
	if err != nil {
		log.Println("Error occured: ", err)
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("Not expected error occured: %v", err))
		return
	}

	log.Println("Devices returned: ", devices)

	utils.WriteJson(w, http.StatusOK, devices)

}

func (c *Controller) GetDevicesByStateController(w http.ResponseWriter, r *http.Request) {

	log.Println("Starting get devices by state controller")

	vars := mux.Vars(r)
	str, ok := vars["state"]

	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Missing the State of the device"))
		return
	}

	devices, err := c.store.GetDevicesByState(str)
	if err != nil {
		log.Println("Error occured: ", err)
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("Not expected error occured: %v", err))
		return
	}

	log.Println("Devices returned: ", devices)

	utils.WriteJson(w, http.StatusOK, devices)

}

func (c *Controller) CreateDeviceController(w http.ResponseWriter, r *http.Request) {

	log.Println("Starting create device controller")

	var device types.CreateDevicePayload

	if err := utils.ParseJson(r, &device); err != nil {
		log.Println("Error occured: ", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid insert payload request: %v", err))
		return
	}

	log.Println("Json Parsed")

	if err := utils.Validate.Struct(device); err != nil {
		errors := err.(validator.ValidationErrors)
		log.Println("Error occured: ", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid insert payload request: %v", errors))
	}

	log.Println("Json Validated")

	err := c.store.CreateDevice(device)
	if err != nil {
		log.Println("Error occured: ", err)
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("Not expected error occured: %v", err))
		return
	}

	log.Println("Device created succesfully: ", device)

	utils.WriteJson(w, http.StatusCreated, device)

}

func (c *Controller) UpdateDeviceController(w http.ResponseWriter, r *http.Request) {

	log.Println("Starting update device controller")

	var device types.UpdateDevicePayload

	if err := utils.ParseJson(r, &device); err != nil {
		log.Println("Error occured: ", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid update payload request: %v", err))
		return
	}

	if err := utils.Validate.Struct(device); err != nil {
		errors := err.(validator.ValidationErrors)
		log.Println("Error occured: ", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid update payload request: %v", errors))
	}

	deviceValid, err := c.store.GetDeviceById(device.Id)
	if err != nil {
		log.Println("Error occured: ", err)
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("Not expected error occured: %v", err))
		return
	}

	err = c.store.UpdateDevice(device, utils.CheckState(deviceValid.State))
	if err != nil {
		log.Println("Error occured: ", err)
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("Not expected error occured: %v", err))
		return
	}
	utils.WriteJson(w, http.StatusOK, device)

	log.Println("Update the device: ", device)

}

func (c *Controller) DeleteDeviceController(w http.ResponseWriter, r *http.Request) {

	log.Println("Starting delete device controller")

	vars := mux.Vars(r)
	str, ok := vars["id"]

	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Missing the ID of the device"))
		return
	}

	id, err := strconv.Atoi(str)
	if err != nil {
		log.Println("Error occured: ", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid delete payload request: %v", err))
		return
	}

	device, err := c.store.GetDeviceById(id)
	if err != nil {
		log.Println("Error occured: ", err)
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("Not expected error occured: %v", err))
		return
	}

	if utils.CheckState(device.State) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Cannot delete active devices"))
		return
	}

	err = c.store.DeleteDevice(id)
	if err != nil {
		log.Println("Error occured: ", err)
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("Not expected error occured: %v", err))
		return
	}

	log.Println("Deleted the device: ", device)

	utils.WriteJson(w, http.StatusOK, "Device deleted")

}
