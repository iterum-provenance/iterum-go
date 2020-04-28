package util

import (
	"encoding/json"
	"io/ioutil"
)

// WriteJSONFile writes an object as json to the specified file path
func WriteJSONFile(filepath string, data interface{}) error {
	file, err := json.MarshalIndent(data, "", " ")
	if err == nil {
		err = ioutil.WriteFile(filepath, file, 0644)
	}
	return err
}

// ReadJSONFile reads json from a file and stores it in the passed interface which should be a pointer
func ReadJSONFile(filepath string, data interface{}) error {
	file, err := ioutil.ReadFile(filepath)
	if err == nil {
		err = json.Unmarshal([]byte(file), data)
	}
	return err
}
