package device

// Device represents a logged in device in the system.
type Device interface {
	GetDeviceID() string //GetDeviceID returns the device ID of the device 
	GetDeviceModel() *string //GetDeviceModel returns the model of the device
}

// New Device represents information of a device extracted from a new log in session
type NewDevice interface {
    GetDeviceID() string //GetDeviceID returns the device ID of the device 
	GetDeviceModel() string //GetDeviceModel returns the model of the device
}

