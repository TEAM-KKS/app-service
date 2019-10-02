package data_models

// eventDescriptor defines the functions that MUST be implemented by events
// this is referenced in the abstract event struct
type eventDescriptor interface {
	// isActive() Confirms whether values in the device that owns this event matches the condition
	// defined by the user when the flow was configured. If it fits the condition, it returns true.
	isActive() bool
	// trigger() sets a value on the device by calling the device's command in EdgeX device service
	trigger(value interface{})
}

// event defines the base construct for an event.
// It mimics an abstract class implementation in object oriented programming paradigms
// All events created should extend this base struct
type event struct {
	name   string
	device device
	eventDescriptor
}
