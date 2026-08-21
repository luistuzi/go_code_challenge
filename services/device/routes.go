package device

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/getDevices", c.GetDevices).Methods("GET")
	router.HandleFunc("/GetDeviceByiD/{id}", c.GetDeviceById).Methods("GET")
	router.HandleFunc("/GetDeviceByBrand/{brand}", c.GetDeviceByBrand).Methods("GET")
	router.HandleFunc("/GetDeviceByState/{state}", c.GetDeviceByState).Methods("GET")
	router.HandleFunc("/devices", c.CreateDevice).Methods("POST")
	router.HandleFunc("/devices/{id}", c.UpdateDevice).Methods("PUT")
	router.HandleFunc("/devices/{id}", c.DeleteDevice).Methods("DELETE")
}

func (c *Controller) GetDevices(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller) GetDeviceById(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller) GetDeviceByBrand(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller) GetDeviceByState(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller) CreateDevice(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller) UpdateDevice(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller) DeleteDevice(w http.ResponseWriter, r *http.Request) {

}
