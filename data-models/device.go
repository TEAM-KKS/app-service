package data_models

// device represents a device data structure in the application layer
// ideally, we would get this data one time from EdgeX when the device is originally created
type device struct {
	// name is the name given to the device in EdgeX
	name string
	// commandName is the name of the 'deviceCommands' entity in the EdgeX device profile
	// helps to know what commands this application can send to a device
	commandName string
	// resouceName is the name of the 'deviceResources' entity in the EdgeX device profile
	resouceName string
	// events is the list of supported events for a device
	events []abstractEvent
}
