package govalidate

type Default struct {}

func (v Default) Validate(tag string, val interface{}) (bool, error) {
	return true, nil
}