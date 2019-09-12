package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/tormoder/fit"
)

// LoadFitActivityFromFile loads a single activity from a fit file
func LoadFitActivityFromFile(filePath string) (*fit.ActivityFile, error) {
	fitFile, loadFileError := ioutil.ReadFile(filePath)
	if loadFileError != nil {
		return nil, fmt.Errorf("Unable to load file '%v'. Does it exists?", filePath)
	}
	fit, decodeFileError := fit.Decode(bytes.NewReader(fitFile))
	if decodeFileError != nil {
		return nil, errors.New("Unable to decode file. Is it a valid FIT file?")
	}
	activity, accessActivityError := fit.Activity()
	if accessActivityError != nil {
		return nil, errors.New("Unable to get activity from FIT file. Is it an activity file?")
	}

	return activity, nil
}
