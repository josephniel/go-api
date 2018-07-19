package util

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// ReadFromFile accepts a string representing a relative path to the file from
// the application folder
func ReadFromFile(file string) (b []byte) {
	absPath, err := filepath.Abs(file)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	b, err = ioutil.ReadFile(absPath)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}
	return
}
