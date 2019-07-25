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

func (e errorRefactor) Initiate(file string, out chan string) (chan string, chan error, error) {
	in := make(chan string)
	erCh := make(chan error)
	go func(in, out chan string, erCh chan error) {
		out <- ""
		<-in
		erCh <- errors.New("Mocking an error")
		erCh <- nil
	}(in, out, erCh)
	return in, erCh, nil
}

func commonValidTest(file string, checkOld, checkNew bool, oldStr, newStr string, source generate.RefactorSource) (string, bool) {
	out := make(chan string)
	in, erCh, err := source.Initiate(testdataDir+string([]rune{filepath.Separator})+file, out)
	if err != nil {
		//error while opening the main_copy.go file to validate
		fmt.Println("Error while reading the source of testdata/main_copy.go")
		return err.Error(), false
	}
	foundNew := false
	foundOld := false
	for s := range out {
		if strings.Contains(s, oldStr) {
			foundOld = true
		}
		if strings.Contains(s, newStr) {
			foundNew = true
		}
		in <- s
		<-erCh
	}
	<-erCh
	if (checkOld && foundOld) || (!checkOld && !foundOld) {
		//the string shouldn'thave been there if sucessfull test
		return "Expected the string cuttle-ai/web-starter to be not found in testdata/main_copy.go", false
	}

	if (checkNew && foundNew) || (!checkNew && !foundNew) {
		//the string should have been there if sucessfull test
		return "Expected the string melvinodsa/test to be found in testdata/main_copy.go", false
	}
	return "", true
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
			return commonValidTest("main_copy.go", true, false, "cuttle-ai/web-starter", "melvinodsa/test", generate.NewPackageRefactor())
		},
	},
	{
		"Normal comment refactor test",
		testdataDir + string([]rune{filepath.Separator}) + "main_copy.go",
		generate.Refactor{
			Name:    "comment project name refactor",
			Find:    "{{.Name}}",
			Replace: "New Name",
			Source:  generate.NewCommentRefactor(),
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
			return commonValidTest("main_copy.go", true, false, "{{.Name}}", "New Name", generate.NewCommentRefactor())
		},
	},
	{
		"Normal non-go file refactor test",
		testdataDir + string([]rune{filepath.Separator}) + "README_copy.md",
		generate.Refactor{
			Name:    "project name refactor",
			Find:    "{{.Name}}",
			Replace: "New Name",
			Source:  generate.NewNonGoFileRefactor(),
		},
		nil,
		func() {
			s, err := os.Open(testdataDir + string([]rune{filepath.Separator}) + "README.md")
			if err != nil {
				//error while opeing the README.md test data file
				fmt.Println("Error while opening the testdata/README.md file", err.Error())
				return
			}
			defer s.Close()
			d, err := os.Create(testdataDir + string([]rune{filepath.Separator}) + "README_copy.md")
			if err != nil {
				//error while creating the copy file of readme
				fmt.Println("Error while creating the testdata/README_copy.md", err.Error())
				return
			}
			defer d.Close()
			_, err = io.Copy(d, s)
			if err != nil {
				//error while copying the main to README_copy
				fmt.Println("Error while copying from testdata/README.md to testdata/README_copy.md", err.Error())
				return
			}
		},
		func() {
			err := os.Remove(testdataDir + string([]rune{filepath.Separator}) + "README_copy.md")
			if err != nil {
				//error while removing the copy of README.md
				fmt.Println("Error while removing the testdata/README_copy.md", err.Error())
				return
			}
		},
		func() (string, bool) {
			return commonValidTest("README_copy.md", true, false, "{{.Name}}", "New Name", generate.NewNonGoFileRefactor())
		},
	},
	{
		"Multiple package refactor test",
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
			return commonValidTest("main_copy.go", true, false, "cuttle-ai/web-starter", "melvinodsa/test", generate.NewPackageRefactor())
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
			return commonValidTest("main_copy.go", true, false, "cuttle-ai/web-starter", "melvinodsa/test", generate.NewPackageRefactor())
		},
	},
	{
		"Regex package multiple occurrences refactor test",
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
			return commonValidTest("main_copy.go", true, false, "cuttle-ai/web-starter", "melvinodsa/test", generate.NewPackageRefactor())
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
