// Copyright 2019 Cuttle.ai. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package generate_test

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/cuttle-ai/web-starter/generate"
)

/*
 * This file contains the test for the source code of refactor.go which describes the refactoring of the given source code
 */

type errorRefactor struct{}

func (e errorRefactor) Src(file string) (string, error) {
	return "", nil
}

func (e errorRefactor) SetSrc(file, code string) error {
	return errors.New("Mocking an error")
}

var refactortcs = []struct {
	Name     string
	File     string
	Refactor generate.Refactor
	Error    error
	Setup    func()
	Teardown func()
	Validate func() (string, bool)
}{
	{
		"Normal package refactor test",
		testdataDir + string([]rune{filepath.Separator}) + "main_copy.go",
		generate.Refactor{
			Name:    "import refactor",
			Find:    "cuttle-ai/web-starter",
			Replace: "melvinodsa/test",
			Source:  generate.NewPackageRefactor(),
		},
		nil,
		func() {
			s, err := os.Open(testdataDir + string([]rune{filepath.Separator}) + "main.go")
			if err != nil {
				//error while opeing the main.go test data file
				fmt.Println("Error while opening the testdata/main.go file", err.Error())
				return
			}
			defer s.Close()
			d, err := os.Create(testdataDir + string([]rune{filepath.Separator}) + "main_copy.go")
			if err != nil {
				//error while creating the copy file of main
				fmt.Println("Error while creating the testdata/main_copy.go", err.Error())
				return
			}
			defer d.Close()
			_, err = io.Copy(d, s)
			if err != nil {
				//error while copying the main to main_copy
				fmt.Println("Error while copying from testdata/main.go to testdata/main_copy.go", err.Error())
				return
			}
		},
		func() {
			err := os.Remove(testdataDir + string([]rune{filepath.Separator}) + "main_copy.go")
			if err != nil {
				//error while removing the copy of main.go
				fmt.Println("Error while removing the testdata/main_copy.go", err.Error())
				return
			}
		},
		func() (string, bool) {
			p := generate.NewPackageRefactor()
			s, err := p.Src(testdataDir + string([]rune{filepath.Separator}) + "main_copy.go")
			if err != nil {
				//error while opening the main_copy.go file to validate
				fmt.Println("Error while reading the source of testdata/main_copy.go")
				return err.Error(), false
			}
			if strings.Contains(s, "cuttle-ai/web-starter") {
				//the string shouldn'thave been there if sucessfull test
				return "Expected the string cuttle-ai/web-starter to be not found in testdata/main_copy.go", false
			}
			if !strings.Contains(s, "melvinodsa/test") {
				//the string should have been there if sucessfull test
				return "Expected the string melvinodsa/test to be found in testdata/main_copy.go", false
			}
			return "", true
		},
	},
	{
		"Multiple package refactor test",
		testdataDir + string([]rune{filepath.Separator}) + "main_copy.go",
		generate.Refactor{
			Name:        "import refactor",
			Find:        "cuttle-ai/web-starter",
			Replace:     "melvinodsa/test",
			Source:      generate.NewPackageRefactor(),
			Occurrences: 3,
		},
		nil,
		func() {
			s, err := os.Open(testdataDir + string([]rune{filepath.Separator}) + "main.go")
			if err != nil {
				//error while opeing the main.go test data file
				fmt.Println("Error while opening the testdata/main.go file", err.Error())
				return
			}
			defer s.Close()
			d, err := os.Create(testdataDir + string([]rune{filepath.Separator}) + "main_copy.go")
			if err != nil {
				//error while creating the copy file of main
				fmt.Println("Error while creating the testdata/main_copy.go", err.Error())
				return
			}
			defer d.Close()
			_, err = io.Copy(d, s)
			if err != nil {
				//error while copying the main to main_copy
				fmt.Println("Error while copying from testdata/main.go to testdata/main_copy.go", err.Error())
				return
			}
		},
		func() {
			err := os.Remove(testdataDir + string([]rune{filepath.Separator}) + "main_copy.go")
			if err != nil {
				//error while removing the copy of main.go
				fmt.Println("Error while removing the testdata/main_copy.go", err.Error())
				return
			}
		},
		func() (string, bool) {
			p := generate.NewPackageRefactor()
			s, err := p.Src(testdataDir + string([]rune{filepath.Separator}) + "main_copy.go")
			if err != nil {
				//error while opening the main_copy.go file to validate
				fmt.Println("Error while reading the source of testdata/main_copy.go")
				return err.Error(), false
			}
			if strings.Contains(s, "cuttle-ai/web-starter") {
				//the string shouldn'thave been there if sucessfull test
				return "Expected the string cuttle-ai/web-starter to be not found in testdata/main_copy.go", false
			}
			if !strings.Contains(s, "melvinodsa/test") {
				//the string should have been there if sucessfull test
				return "Expected the string melvinodsa/test to be found in testdata/main_copy.go", false
			}
			return "", true
		},
	},
	{
		"Regex package refactor test",
		testdataDir + string([]rune{filepath.Separator}) + "main_copy.go",
		generate.Refactor{
			Name:    "import refactor",
			Find:    "cuttle-ai/web-starter",
			Replace: "melvinodsa/test",
			Source:  generate.NewPackageRefactor(),
			IsRegex: true,
		},
		nil,
		func() {
			s, err := os.Open(testdataDir + string([]rune{filepath.Separator}) + "main.go")
			if err != nil {
				//error while opeing the main.go test data file
				fmt.Println("Error while opening the testdata/main.go file", err.Error())
				return
			}
			defer s.Close()
			d, err := os.Create(testdataDir + string([]rune{filepath.Separator}) + "main_copy.go")
			if err != nil {
				//error while creating the copy file of main
				fmt.Println("Error while creating the testdata/main_copy.go", err.Error())
				return
			}
			defer d.Close()
			_, err = io.Copy(d, s)
			if err != nil {
				//error while copying the main to main_copy
				fmt.Println("Error while copying from testdata/main.go to testdata/main_copy.go", err.Error())
				return
			}
		},
		func() {
			err := os.Remove(testdataDir + string([]rune{filepath.Separator}) + "main_copy.go")
			if err != nil {
				//error while removing the copy of main.go
				fmt.Println("Error while removing the testdata/main_copy.go", err.Error())
				return
			}
		},
		func() (string, bool) {
			p := generate.NewPackageRefactor()
			s, err := p.Src(testdataDir + string([]rune{filepath.Separator}) + "main_copy.go")
			if err != nil {
				//error while opening the main_copy.go file to validate
				fmt.Println("Error while reading the source of testdata/main_copy.go")
				return err.Error(), false
			}
			if strings.Contains(s, "cuttle-ai/web-starter") {
				//the string shouldn'thave been there if sucessfull test
				return "Expected the string cuttle-ai/web-starter to be not found in testdata/main_copy.go", false
			}
			if !strings.Contains(s, "melvinodsa/test") {
				//the string should have been there if sucessfull test
				return "Expected the string melvinodsa/test to be found in testdata/main_copy.go", false
			}
			return "", true
		},
	},
	{
		"Regex package multiple occurrences refactor test",
		testdataDir + string([]rune{filepath.Separator}) + "main_copy.go",
		generate.Refactor{
			Name:        "import refactor",
			Find:        "cuttle-ai/web-starter",
			Replace:     "melvinodsa/test",
			Source:      generate.NewPackageRefactor(),
			IsRegex:     true,
			Occurrences: 3,
		},
		nil,
		func() {
			s, err := os.Open(testdataDir + string([]rune{filepath.Separator}) + "main.go")
			if err != nil {
				//error while opeing the main.go test data file
				fmt.Println("Error while opening the testdata/main.go file", err.Error())
				return
			}
			defer s.Close()
			d, err := os.Create(testdataDir + string([]rune{filepath.Separator}) + "main_copy.go")
			if err != nil {
				//error while creating the copy file of main
				fmt.Println("Error while creating the testdata/main_copy.go", err.Error())
				return
			}
			defer d.Close()
			_, err = io.Copy(d, s)
			if err != nil {
				//error while copying the main to main_copy
				fmt.Println("Error while copying from testdata/main.go to testdata/main_copy.go", err.Error())
				return
			}
		},
		func() {
			err := os.Remove(testdataDir + string([]rune{filepath.Separator}) + "main_copy.go")
			if err != nil {
				//error while removing the copy of main.go
				fmt.Println("Error while removing the testdata/main_copy.go", err.Error())
				return
			}
		},
		func() (string, bool) {
			p := generate.NewPackageRefactor()
			s, err := p.Src(testdataDir + string([]rune{filepath.Separator}) + "main_copy.go")
			if err != nil {
				//error while opening the main_copy.go file to validate
				fmt.Println("Error while reading the source of testdata/main_copy.go")
				return err.Error(), false
			}
			if strings.Contains(s, "cuttle-ai/web-starter") {
				//the string shouldn'thave been there if sucessfull test
				return "Expected the string cuttle-ai/web-starter to be not found in testdata/main_copy.go", false
			}
			if !strings.Contains(s, "melvinodsa/test") {
				//the string should have been there if sucessfull test
				return "Expected the string melvinodsa/test to be found in testdata/main_copy.go", false
			}
			return "", true
		},
	},
	{
		"Regex package multiple occurrences with not all refactor done test",
		testdataDir + string([]rune{filepath.Separator}) + "main_copy.go",
		generate.Refactor{
			Name:        "import refactor",
			Find:        "cuttle-ai/web-starter",
			Replace:     "melvinodsa/test",
			Source:      generate.NewPackageRefactor(),
			IsRegex:     true,
			Occurrences: 2,
		},
		nil,
		func() {
			s, err := os.Open(testdataDir + string([]rune{filepath.Separator}) + "main.go")
			if err != nil {
				//error while opeing the main.go test data file
				fmt.Println("Error while opening the testdata/main.go file", err.Error())
				return
			}
			defer s.Close()
			d, err := os.Create(testdataDir + string([]rune{filepath.Separator}) + "main_copy.go")
			if err != nil {
				//error while creating the copy file of main
				fmt.Println("Error while creating the testdata/main_copy.go", err.Error())
				return
			}
			defer d.Close()
			_, err = io.Copy(d, s)
			if err != nil {
				//error while copying the main to main_copy
				fmt.Println("Error while copying from testdata/main.go to testdata/main_copy.go", err.Error())
				return
			}
		},
		func() {
			err := os.Remove(testdataDir + string([]rune{filepath.Separator}) + "main_copy.go")
			if err != nil {
				//error while removing the copy of main.go
				fmt.Println("Error while removing the testdata/main_copy.go", err.Error())
				return
			}
		},
		func() (string, bool) {
			p := generate.NewPackageRefactor()
			s, err := p.Src(testdataDir + string([]rune{filepath.Separator}) + "main_copy.go")
			if err != nil {
				//error while opening the main_copy.go file to validate
				fmt.Println("Error while reading the source of testdata/main_copy.go")
				return err.Error(), false
			}
			if !strings.Contains(s, "cuttle-ai/web-starter") {
				//the string shouldn'thave been there if sucessfull test
				return "Expected the string cuttle-ai/web-starter to be found in testdata/main_copy.go", false
			}
			if !strings.Contains(s, "melvinodsa/test") {
				//the string should have been there if sucessfull test
				return "Expected the string melvinodsa/test to be found in testdata/main_copy.go", false
			}
			return "", true
		},
	},
	{
		"File missing",
		testdataDir + string([]rune{filepath.Separator}) + "main_copy.go",
		generate.Refactor{
			Name:    "import refactor",
			Find:    "cuttle-ai/web-starter",
			Replace: "melvinodsa/test",
			Source:  generate.NewPackageRefactor(),
		},
		errors.New("Source file missing"),
		func() {},
		func() {},
		func() (string, bool) {
			return "", true
		},
	},
	{
		"Error Refactor",
		testdataDir + string([]rune{filepath.Separator}) + "main_copy.go",
		generate.Refactor{
			Name:    "import refactor",
			Find:    "cuttle-ai/web-starter",
			Replace: "melvinodsa/test",
			Source:  errorRefactor{},
		},
		errors.New("Error while setting the source"),
		func() {},
		func() {},
		func() (string, bool) {
			return "", true
		},
	},
	{
		"Error Refactor multiple occurences",
		testdataDir + string([]rune{filepath.Separator}) + "main_copy.go",
		generate.Refactor{
			Name:        "import refactor",
			Find:        "cuttle-ai/web-starter",
			Replace:     "melvinodsa/test",
			Source:      errorRefactor{},
			Occurrences: 3,
		},
		errors.New("Error while setting the source"),
		func() {},
		func() {},
		func() (string, bool) {
			return "", true
		},
	},
	{
		"Error Refactor regex",
		testdataDir + string([]rune{filepath.Separator}) + "main_copy.go",
		generate.Refactor{
			Name:    "import refactor",
			Find:    "cuttle-ai/web-starter",
			Replace: "melvinodsa/test",
			Source:  errorRefactor{},
			IsRegex: true,
		},
		errors.New("Error while setting the source"),
		func() {},
		func() {},
		func() (string, bool) {
			return "", true
		},
	},
	{
		"Error Refactor with regex and multiple occurences",
		testdataDir + string([]rune{filepath.Separator}) + "main_copy.go",
		generate.Refactor{
			Name:        "import refactor",
			Find:        "cuttle-ai/web-starter",
			Replace:     "melvinodsa/test",
			Source:      errorRefactor{},
			IsRegex:     true,
			Occurrences: 3,
		},
		errors.New("Error while setting the source"),
		func() {},
		func() {},
		func() (string, bool) {
			return "", true
		},
	},
}

//TestRefactorDo is the test suite for the Do method of refactor struct
func TestRefactorDo(t *testing.T) {
	for _, v := range refactortcs {
		t.Run(v.Name, func(t *testing.T) {
			fmt.Println("Testing", v.Name)
			defer v.Teardown()
			if v.Setup != nil {
				v.Setup()
			}
			err := v.Refactor.Do(v.File)
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
