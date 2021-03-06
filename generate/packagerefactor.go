// Copyright 2019 Cuttle.ai. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package generate

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
)

/*
 * This file has the defintions of the package name refactor struct which implements the RefactorSource.
 */

//PackageRefactor is the to refactor the import package names present in a source code file
type PackageRefactor struct {
}

//NewPackageRefactor is the constructor for the package import name refactor
func NewPackageRefactor() PackageRefactor {
	return PackageRefactor{}
}

//Initiate will initiate the refactoring of import packages corresponding to the go source file name provided
//It will send the imports through the out channel and expects the refactored input for each import send in the output to come
//If any error happens in the process it will send through the error channel.
//After ranging over the out channel check for the last error from the error channel so that error while
//finishing the refactoring is captured
func (p PackageRefactor) Initiate(file string, out chan string) (chan string, chan error, error) {
	/*
	 * We will create an input channel, error channel
	 * We will parse the imports of the source file
	 * Will initiate a go routine that processes the imports and send it to the out channel
	 */
	//fileset to trcak the source code positions
	fset := token.NewFileSet()

	//creating the input and error output channel
	in, errCh := make(chan string), make(chan error)

	//parsing the source for the given file
	f, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
	if err != nil {
		//error while parsing the source file
		fmt.Println("Error while parsing the go source file", file)
		return in, errCh, err
	}

	//invoking the go routine to process the import refactoring through the channels
	go processImports(file, in, out, errCh, f, fset)

	return in, errCh, nil
}

func processImports(file string, in chan string, out chan string, erCh chan error, f *ast.File, fset *token.FileSet) {
	/*
	 * We will iterate through the imports
	 * Will send them to the output channel
	 * Will wait for the input channel to get the refactored string
	 * Writes nil to the error channel
	 * Finally overwrite the existing file
	 */
	//iterating through the import paths and store it in imports array
	for _, v := range f.Imports {
		//Sending them to the output channel
		out <- v.Path.Value

		//waiting for the input channel to give the refactored imports
		v.Path.Value = <-in

		//outputing nil error to the err channel
		erCh <- nil
	}
	close(out)

	//overwrting the existing file
	c, _ := os.Create(file)
	defer c.Close()
	//we are ignoring the error, since it is safe to do so in this context
	//we are writing the formatted code into the file itself
	err := format.Node(c, fset, f)
	if err != nil {
		//error while writing the refactored source code to the file
		fmt.Println("Error while writing the formatted code to the given file", file)
	}

	erCh <- err
	close(erCh)
}
