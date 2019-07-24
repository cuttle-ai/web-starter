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
	//IsRegex indicates that the string to be found is a regular expression
	IsRegex bool
	//Source is the source to be refactored
	Source RefactorSource
}

//RefactorSource has to be implmented by any source to act as an source code to be refactored
type RefactorSource interface {
	//Initiate will read the given file and initiate the refactoring. Through the returned output
	//channel the code to be refactored will be streamed and through the in channel it will replace the code with the refactored one.
	//During processing of each source code stream, if any error occurs it will be pushed to the error channel. Else nil will be pushed to it.
	//The absolute path to the file from the refactor source code string has to be recovered
	//This method also returns the error if any else nil is returned
	Initiate(file string, in chan string) (chan string, chan error, error)
}

//Do will make the necessary changes in the given file as per the reafctor
//Will return an error something unexpected comes up or refactoring fails
//This method requires the absolute path to the file.
func (r Refactor) Do(file string) error {
	/*
	 * Will check whether the refactor source is nil or not
	 * Will initiate the refactor source
	 * Will get the source code to be refactored
	 * If the refactor is a regex based one, will handle in accordance
	 * Else simple string find and replace will be done
	 */

	//checking for the nil source
	if r.Source == nil {
		//the source is nil
		return errors.New("the refactor source is nil. Not refactoring " + r.Name)
	}

	//initiating the source
	in := make(chan string)
	out, errChan, err := r.Source.Initiate(file, in)
	if err != nil {
		//error while initiating the source for the file
		fmt.Println("Error while initiating the source for the file", file)
		return err
	}

	for source := range in {
		//handling the regex type of refactoring
		if r.IsRegex {
			reg := regexp.MustCompile(r.Find)
			//if the count is < 1 then the count is infinite
			//if the count is > 1 find the occurrence till count < 1
			out <- reg.ReplaceAllString(source, r.Replace)
			err := <-errChan
			if err != nil {
				//error while setting refactored string
				fmt.Println("Error while replacing the occurrence refactoring", r.Name, file, r.Find, "->", r.Replace)
				return err
			}
			continue
		}

		//handling simple string replacement
		out <- strings.Replace(source, r.Find, r.Replace, -1)
		err := <-errChan
		if err != nil {
			//error while setting refactored string
			fmt.Println("Error while replacing the occurrence refactoring", r.Name, file, r.Find, "->", r.Replace)
			return err
		}
	}
	err = <-errChan
	return err
}

//String is the stringer implementation of the Refactor
func (r Refactor) String() string {
	return r.Name
}
