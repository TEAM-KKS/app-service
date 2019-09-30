package data_models

type device struct {
	name         string
	ip           string
	port         string
	manufacturer string
	attributes   data   // stores attributes and values for each device
	profile      string // device profile name on EdgeX
}

func (d *device) addAttr(name string, value interface{}) error {
	return d.attributes.addAttr(name, value)
}

func (d *device) removeAttr(name string) {
	d.attributes.removeAttr(name)
}
