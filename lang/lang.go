package lang

import (
	"io/ioutil"
	"fmt"
	"os"
	"encoding/json"
	"runtime"
	"path"
)
type MessageError struct {
	Tag string
	Error string
}
func GetMessageErrorFromTag(tag string, value ...interface{} ) (error error) {
	customLangFile := os.Getenv("VALIDATE_LANG_FILE")
	
	_, filename, _, _ := runtime.Caller(0)
	filepath := path.Join(path.Dir(filename), "./en.json")
	
	if(customLangFile != ""){
		filepath = customLangFile
	}

	raw, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	var messagesErrors []MessageError
	json.Unmarshal(raw, &messagesErrors)
	for _, v := range messagesErrors {
		if v.Tag == tag {
			if len(value) == 0 {
				error  = fmt.Errorf(v.Error)
			}else{
				error  = fmt.Errorf(v.Error, value...)
			}
		}
	}
	return
}
