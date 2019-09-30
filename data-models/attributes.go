package data_models

import (
	"errors"
)

type data struct {
	attrMap map[string]interface{}
}

type attribues interface {
	addAttr()
	removeAttr()
}

func (d *data) addAttr(name string, value interface{}) error {
	if len(d.attrMap) == 0 {
		d.attrMap = make(map[string]interface{})
	}
	switch value.(type) {
	case int:
		d.attrMap[name] = value.(int)
	case float64:
		d.attrMap[name] = value.(float64)
	case string:
		d.attrMap[name] = value.(string)
	case bool:
		d.attrMap[name] = value.(bool)
	default:
		return errors.New("Attribute is an Unsupported Value Type")
	}
	return nil
}

func (d *data) removeAttr(name string) {
	delete(d.attrMap, name)
}
