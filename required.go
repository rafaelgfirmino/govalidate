package govalidate

import "github.com/rafaelgfirmino/govalidate/lang"

type Required struct {}

func (v Required) Validate(tag string, val interface{}) (bool, error) {

	if val.(string) == "" {
		return false, lang.GetMessageErrorFromTag(tag)
	}
	return true, nil
}