package govalidate

import (
	"regexp"
	"github.com/rafaelgfirmino/govalidate/lang"
)

type Email struct {
	Name string
}

var mailRegex = regexp.MustCompile(`\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`)

func (v Email) Validate(tag string,val interface{}) (bool, error) {
	if !mailRegex.MatchString(val.(string)) {
		return false, lang.GetMessageErrorFromTag(tag,val.(string))
	}
	return true, nil
}