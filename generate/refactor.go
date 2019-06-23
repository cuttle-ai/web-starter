// Copyright 2019 Cuttle.ai. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package generate

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

/*
 * This file has the definitions of the boiler plate code source file refactor
 */

//Refactor is represents a change to be done in a source file
type Refactor struct {
	//Name to identify the refactor
	Name string
	//Find is the string to be found
	Find string
	//Replace is the string to be replaced in the place of the string to be found
	Replace string
	//Occurrences is max the no. of occurrences to be replaced. If the occurrences is < 1,
	//then it is considered as infinite
	Occurrences int
	//IsRegex indicates that the string to be found is a regular expression
	IsRegex bool
	//Source is the source to be refactored
	Source RefactorSource
}

//RefactorSource has to be implmented by any source to act as an source code to be refactored
type RefactorSource interface {
	//Src returns the source code to be refactored.
	//The absolute path to the file from the refactor source code string has to be recovered
	//This method also returns the error if any else nil is returned
	Src(file string) (string, error)
	//SetSrc will set the refactored code to the source.
	//The source file should have the absolute path to the file.
	//Error returned will be nil if no error happens.
	SetSrc(file, code string) error
}

//Do will make the necessary changes in the given file as per the reafctor
//Will return an error something unexpected comes up or refactoring fails
//This method requires the absolute path to the file.
func (r Refactor) Do(file string) error {
	/*
	 * Will check whether the refactor source is nil or not
	 * Will get the source code to be refactored
	 * If the refactor is a regex based one, will handle in accordance
	 * Else simple string find and replace will be done
	 */

	//checking for the nil source
	if r.Source == nil {
		//the source is nil
		return errors.New("the refactor source is nil. Not refactoring " + r.Name)
	}

	//getting the source code to be refactored
	source, err := r.Source.Src(file)
	if err != nil {
		//error while getting the source from the file
		fmt.Println("Error while getting the source code from the file", file)
		return err
	}

	//handling the regex type of refactoring
	if r.IsRegex {
		reg := regexp.MustCompile(r.Find)
		count := r.Occurrences
		//if the count is < 1 then the count is infinite
		//if the count is > 1 find the occurrence till count < 1
		if count < 1 {
			err := r.Source.SetSrc(file, reg.ReplaceAllString(source, r.Replace))
			if err != nil {
				//error while setting refactored string
				fmt.Println("Error while replacing the occurrence refactoring", r.Name, file, r.Find, "->", r.Replace)
				return err
			}
			return nil
		}
		err := r.Source.SetSrc(file, reg.ReplaceAllStringFunc(source, func(s string) string {
			if count < 1 {
				return s
			}
			count--
			return r.Replace
		}))
		if err != nil {
			//error while setting refactored string
			fmt.Println("Error while replacing the occurrence refactoring", r.Name, file, r.Find, "->", r.Replace)
			return err
		}
		return nil
	}

	//handling simple string replacement
	if r.Occurrences < 1 {
		//if the count is < 1 then replace everything
		err := r.Source.SetSrc(file, strings.Replace(source, r.Find, r.Replace, -1))
		if err != nil {
			//error while setting refactored string
			fmt.Println("Error while replacing the occurrence refactoring", r.Name, file, r.Find, "->", r.Replace)
			return err
		}
		return nil
	}
	//in normal cases just replace the occurrences
	err = r.Source.SetSrc(file, strings.Replace(source, r.Find, r.Replace, r.Occurrences))
	if err != nil {
		//error while setting refactored string
		fmt.Println("Error while replacing the occurrence refactoring", r.Name, file, r.Find, "->", r.Replace)
		return err
	}

	return nil
}

//String is the stringer implementation of the Refactor
func (r Refactor) String() string {
	return r.Name
}
