// Copyright 2019 Cuttle.ai. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package generate_test

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"testing"

	"github.com/cuttle-ai/web-starter/generate"
)

/*
 * This file contains the tests written for the source code in source.go
 */

var testdataDir = func() string {
	t := os.Getenv("GENERATE_TESTDATA")
	if len(t) == 0 {
		t = "./testdata"
	}
	return t
}()

var sourcetcs = []struct {
	Name     string
	Source   generate.Source
	Dst      string
	Error    error
	Setup    func()
	Teardown func()
	Validate func() (string, bool)
}{
	{
		"Copy file",
		generate.Source{
			Path:      testdataDir,
			FileName:  "main.go",
			Refactors: []generate.Refactor{},
		},
		testdataDir + string([]rune{filepath.Separator}) + "copied",
		nil,
		func() {},
		func() {
			os.RemoveAll(testdataDir + string([]rune{filepath.Separator}) + "copied")
		},
		func() (string, bool) {
			_, err := os.Stat(testdataDir + string([]rune{filepath.Separator}) +
				"copied" + string([]rune{filepath.Separator}) + "main.go")
			if err != nil {
				return err.Error(), false
			}
			return "", true
		},
	},
	{
		"File doesn't exist",
		generate.Source{
			Path:      testdataDir,
			FileName:  "main1.go",
			Refactors: []generate.Refactor{},
		},
		testdataDir + string([]rune{filepath.Separator}) + "copied",
		errors.New("Path doesn't exist"),
		func() {},
		func() {},
		func() (string, bool) {
			_, err := os.Stat(testdataDir + string([]rune{filepath.Separator}) +
				"copied" + string([]rune{filepath.Separator}) + "main.go")
			if err != nil {
				return "", true
			}
			return "Expected no copied file. Got one", false
		},
	},
	{
		"No valid refactor source",
		generate.Source{
			Path:     testdataDir,
			FileName: "main.go",
			Refactors: []generate.Refactor{
				{Name: "Refactor main packages"},
			},
		},
		testdataDir + string([]rune{filepath.Separator}) + "copied",
		errors.New("the refactor source is nil. Not refactoring " + "Refactor main packages"),
		func() {},
		func() {
			os.RemoveAll(testdataDir + string([]rune{filepath.Separator}) + "copied")
		},
		func() (string, bool) {
			_, err := os.Stat(testdataDir + string([]rune{filepath.Separator}) +
				"copied" + string([]rune{filepath.Separator}) + "main.go")
			if err != nil {
				return err.Error(), false
			}
			return "", true
		},
	},
	{
		"File is not regular",
		generate.Source{
			Path:      testdataDir,
			FileName:  "copied",
			Refactors: []generate.Refactor{},
		},
		testdataDir + string([]rune{filepath.Separator}) + "copied",
		errors.New("copied is not a regular file"),
		func() {
			err := os.Mkdir(testdataDir+string([]rune{filepath.Separator})+"copied", 0755)
			if err != nil {
				fmt.Println("Error while making the directory", testdataDir+string([]rune{filepath.Separator})+"copied", err)
			}
		},
		func() {
			os.RemoveAll(testdataDir + string([]rune{filepath.Separator}) + "copied")
		},
		func() (string, bool) {
			_, err := os.Stat(testdataDir + string([]rune{filepath.Separator}) + "copied")
			if err != nil {
				return err.Error(), false
			}
			return "", true
		},
	},
	{
		"Destination directory is a file",
		generate.Source{
			Path:      testdataDir,
			FileName:  "main.go",
			Refactors: []generate.Refactor{},
		},
		testdataDir + string([]rune{filepath.Separator}) + "main.go",
		&os.PathError{Op: "mkdir", Path: testdataDir + string([]rune{filepath.Separator}) + "main.go", Err: syscall.ENOTDIR},
		func() {},
		func() {},
		func() (string, bool) {
			return "", true
		},
	},
}

//TestSource is the test suite for the Generate method of Source data structure
func TestSourceGenerate(t *testing.T) {
	for _, v := range sourcetcs {
		t.Run(v.Name, func(t *testing.T) {
			fmt.Println("Testing", v.Name)
			defer v.Teardown()
			if v.Setup != nil {
				v.Setup()
			}
			err := v.Source.Generate(v.Dst)
			if err == nil && v.Error != nil {
				t.Error("Expected an error.", v.Error.Error(), "Got none.")
				return
			}
			if err != nil && v.Error == nil {
				t.Error("Didn't expect an error. Got one", err.Error())
				return
			}
			if res, ok := v.Validate(); !ok {
				t.Error("Test result validation failed", res)
				return
			}
		})
	}
}
