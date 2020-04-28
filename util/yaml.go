package util

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// WriteYAMLFile writes an object as yaml to the specified file path
func WriteYAMLFile(filepath string, data interface{}) error {
	file, err := yaml.Marshal(data)
	if err == nil {
		err = ioutil.WriteFile(filepath, file, 0644)
	}
	return err
}

// ReadYAMLFile reads yaml from a file and stores it in the passed interface which should be a pointer
func ReadYAMLFile(filepath string, data interface{}) error {
	file, err := ioutil.ReadFile(filepath)
	if err == nil {
		err = yaml.Unmarshal([]byte(file), data)
	}
	return err
}
