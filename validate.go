package govalidate

import (
	"reflect"
	"strings"
)

const tagName= "validate"

type Validate interface {
	// validate method performs validation and returns result and optional error.
	Validate(string, interface{}) (bool, error)
}

func Struct(s interface{}) map[string][]interface{} {
	v := reflect.ValueOf(s)
	erros := make(map[string][]interface{})
	for i := 0; i < v.NumField(); i++ {
		fieldName := strings.ToLower(v.Type().Field(i).Name)
		if v.Field(i).Kind() == reflect.Struct{
			if v.Field(i).NumField() > 0 {
				r := reflect.Indirect(v.Field(i))
				errosRecursive  := Struct(r.Interface())
				for k, v := range errosRecursive {
					erros[fieldName] = append(erros[k],map[string]interface{}{k: v} )
				}
			}
		}
		tag := v.Type().Field(i).Tag.Get(tagName)

		if tag == "" || tag == "-" {
			continue
		}

		tags := strings.Split(tag,",")
		for _, tagValue := range tags {
			validator := getValidator(tagValue)
			valid, err := validator.Validate(tagValue,v.Field(i).Interface())
			if !valid && err != nil {
				erros[fieldName] = append(erros[fieldName],err.Error() )
			}
		}
	}
	return erros
}

func getValidator(tag string) Validate{
	if tag, ok := TagMap[tag]; ok {
		return tag
	}
	return Default{}
}

var TagMap = map[string]Validate {
	"email": Email{},
	"required": Required{},
}