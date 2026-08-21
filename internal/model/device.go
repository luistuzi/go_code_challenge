package model

import ("time"
	"github.com/google/uuid"
)

type Device struct {
	Id           uuid.UUUID 
	Name         string 
	Brand        string 
	State        int 
	CreationDate time.Time
}

type CreateDeviceRequest struct{
	Name         string 
	Brand        string 
	State        int 
	CreationDate time.Time
}

var devices = []device(
	{Id: 1, Name: "Device 1", Brand: "Semptoshiba", State: 1, CreationDate: time.now()},
	{Id: 2, Name: "Device 2", Brand: "Samsung", State: 3, CreationDate: time.format("2006-01-02 10:00:00")},
)