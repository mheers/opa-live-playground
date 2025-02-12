package utils

import (
	"errors"
	"fmt"
	"os"

	"raygun/config"
	"raygun/finder"
	"raygun/log"
	"raygun/parser"
	"raygun/types"
)

func GetPresets() ([]types.TestRecord, error) {
	var entities = make([]string, 0)

	presetsPath := "presets"
	if os.Getenv("PRESETS_PATH") != "" {
		presetsPath = os.Getenv("PRESETS_PATH")
	}

	entities = append(entities, presetsPath)

	/*
	 *  Find the raygun files amidst the files and directories specified on the command line
	 */
	finder := finder.NewFinder(".raygun")

	suite_files, err := finder.FindTargets(entities)

	if err != nil {
		return nil, fmt.Errorf("error finding test suites: %v", err)
	}

	if len(suite_files) == 0 {
		return nil, fmt.Errorf("no .raygun files found in specified location(s)")
	}

	/*
	 *  Parse the .raygun files that we found in the previous step
	 */

	log.Verbose("Parsing Raygun files: %v", suite_files)

	parser := parser.NewRaygunParser(config.SkipOnParseError)

	test_suite_list, err := parser.Parse(suite_files)
	if err != nil {
		return nil, fmt.Errorf("unable to parse test files: %v", err)
	}

	records := make([]types.TestRecord, 0)

	for _, suite := range test_suite_list {
		records = append(records, suite.Tests...)
	}

	return records, nil
}

func GetPreset(name string) (*types.TestRecord, error) {
	presets, err := GetPresets()
	if err != nil {
		return nil, err
	}
	for _, preset := range presets {
		if preset.Name == name {
			return &preset, nil
		}
	}
	return nil, errors.New("preset not found")
}
