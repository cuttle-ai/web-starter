// Copyright 2019 Cuttle.ai. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package generate

import "os"

/*
 * This file has the definitions of the boiler plate code source file refactor
 */

//Refactor is represents a change to be done in a source file
type Refactor struct {
	//Name to identify the refactor
	Name string
}

//Do will make the necessary changes in the given file as per the reactor
//Will return an error something unexpected comes up or refactoring fails
func (r Refactor) Do(f *os.File) error {

	return nil
}

//String is the stringer implementation of the Refactor
func (r Refactor) String() string {
	return r.Name
}
