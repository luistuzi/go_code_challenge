package controller

import (
	"fmt"
	"go_code_challenge/types"
	"go_code_challenge/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	_ "go_code_challenge/docs"
)

type Controller struct {
	service types.DeviceService
}

func NewController(service types.DeviceService) *Controller {
	return &Controller{service: service}
}

func (c *Controller) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/getDevices", c.GetDevicesController).Methods("GET")
	router.HandleFunc("/getDeviceById/{id}", c.GetDeviceByIdController).Methods("GET")
	router.HandleFunc("/getDevicesByBrand/{brand}", c.GetDevicesByBrandController).Methods("GET")
	router.HandleFunc("/getDevicesByState/{state}", c.GetDevicesByStateController).Methods("GET")
	router.HandleFunc("/createDevice", c.CreateDeviceController).Methods("POST")
	router.HandleFunc("/updateDevice", c.UpdateDeviceController).Methods("PATCH")
	router.HandleFunc("/deleteDevice/{id}", c.DeleteDeviceController).Methods("DELETE")

	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler())
}

// GetDevices search for all devices
//
// @Summary Search for all devices
// @Description Return a list of all devices
// @Tags Devices
// @Produce json
// @Success 200 {object} types.Device
// @Failure 400
// @Failure 500
// @Router /getDevices [get]
func (c *Controller) GetDevicesController(w http.ResponseWriter, r *http.Request) {
	log.Println("Starting get devices controller")

	devices, err := c.service.GetDevices()

	log.Println("Devices returned: ", devices)

	if err != nil {
		log.Println("Error occured: ", err)
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf(utils.UnexpectedError, err))
		return
	}

	utils.WriteJson(w, http.StatusOK, devices)

}

// GetDeviceById search for a individual device
//
// @Summary Search for a specific device
// @Description Return a device by its id
// @Tags Devices
// @Produce json
// @Param id path int true "device id"
// @Success 200 {object} types.Device
// @Failure 400
// @Failure 500
// @Router /getDeviceById/{id} [get]
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

	device, err := c.service.GetDeviceById(id)
	if err != nil {
		log.Println("Error occured: ", err)
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf(utils.UnexpectedError, err))
		return
	}

	log.Println("Device returned: ", device)

	utils.WriteJson(w, http.StatusOK, device)

}

// GetDevicesByBrand search for all devices with the specific brand
//
// @Summary Search for all devices with the specific brand
// @Description Return a list of all devices with the specific brand
// @Tags Devices
// @Produce json
// @Param brand path string true "device brand"
// @Success 200 {object} types.Device
// @Failure 400
// @Failure 500
// @Router /getDevicesByBrand/{brand} [get]
func (c *Controller) GetDevicesByBrandController(w http.ResponseWriter, r *http.Request) {

	log.Println("Starting get devices by brand controller")

	vars := mux.Vars(r)
	str, ok := vars["brand"]

	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Missing the Brand of the device"))
		return
	}

	devices, err := c.service.GetDevicesByBrand(str)
	if err != nil {
		log.Println("Error occured: ", err)
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf(utils.UnexpectedError, err))
		return
	}

	log.Println("Devices returned: ", devices)

	utils.WriteJson(w, http.StatusOK, devices)

}

// GetDevicesByState search for all devices with the specific state
//
// @Summary Search for all devices with the specific state
// @Description Return a list of all devices with the specific state
// @Tags Devices
// @Produce json
// @Param state path string true "device state"
// @Success 200 {object} types.Device
// @Failure 400
// @Failure 500
// @Router /getDevicesByState/{state} [get]
func (c *Controller) GetDevicesByStateController(w http.ResponseWriter, r *http.Request) {

	log.Println("Starting get devices by state controller")

	vars := mux.Vars(r)
	str, ok := vars["state"]

	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Missing the State of the device"))
		return
	}

	devices, err := c.service.GetDevicesByState(str)
	if err != nil {
		log.Println("Error occured: ", err)
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf(utils.UnexpectedError, err))
		return
	}

	log.Println("Devices returned: ", devices)

	utils.WriteJson(w, http.StatusOK, devices)

}

// CreateDevice insert a new device.
//
// @Summary Create a new device
// @Description Create a new device
// @Tags Devices
// @Accept json
// @Produce json
// @Param device body types.CreateDevicePayload true "Device Id"
// @Success 201 {object} types.Device
// @Failure 400
// @Failure 409
// @Failure 500
// @Router /createDevice [post]
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

	err := c.service.CreateDevice(device)
	if err != nil {
		log.Println("Error occured: ", err)
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf(utils.UnexpectedError, err))
		return
	}

	log.Println("Device created succesfully: ", device)

	utils.WriteJson(w, http.StatusCreated, device)

}

// UpdateDevice Update parcially the device
//
// @Summary update the device
// @Description Udate only provided fields and with no "in-use" state
// @Tags Devices
// @Accept json
// @Produce json
// @Param device body types.UpdateDevicePayload true "Device Id"
// @Success 204
// @Failure 400
// @Failure 500
// @Router /updateDevice [patch]
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

	deviceValid, err := c.service.GetDeviceById(device.Id)
	if err != nil {
		log.Println("Error occured: ", err)
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf(utils.UnexpectedError, err))
		return
	}

	err = c.service.UpdateDevice(device, utils.CheckState(deviceValid.State))
	if err != nil {
		log.Println("Error occured: ", err)
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf(utils.UnexpectedError, err))
		return
	}
	utils.WriteJson(w, http.StatusOK, device)

	log.Println("Update the device: ", device)

}

// DeleteDevice Delete the device
//
// @Summary delete the device
// @Description Delete the provided device
// @Tags Devices
// @Produce json
// @Param id path int true "device Id"
// @Success 200
// @Failure 400
// @Failure 404
// @Router /deleteDevice/{id} [delete]
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

	device, err := c.service.GetDeviceById(id)
	if err != nil {
		log.Println("Error occured: ", err)
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf(utils.UnexpectedError, err))
		return
	}

	if utils.CheckState(device.State) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Cannot delete in-use devices"))
		return
	}

	err = c.service.DeleteDevice(id)
	if err != nil {
		log.Println("Error occured: ", err)
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf(utils.UnexpectedError, err))
		return
	}

	log.Println("Deleted the device: ", device)

	utils.WriteJson(w, http.StatusOK, "Device deleted")

}
