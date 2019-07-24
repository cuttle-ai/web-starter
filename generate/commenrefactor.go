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
 * This file has the defintions of the comment refactor struct which implements the RefactorSource.
 */

//CommentRefactor is the to refactor the comments present in a source code file
type CommentRefactor struct{}

//NewCommentRefactor is the constructor for the comment refactor
//Default value of the separator is |
func NewCommentRefactor() CommentRefactor {
	return CommentRefactor{}
}

//Initiate will initiate the refactoring of comments corresponding to the go source file name provided
//It will send the comments through the out channel and expects the refactored input for each import send in the output to come
//If any error happens in the process it will send through the error channel.
//After ranging over the out channel check for the last error from the error channel so that error while
//finishing the refactoring is captured
func (p CommentRefactor) Initiate(file string, out chan string) (chan string, chan error, error) {
	/*
	 * We will create an input channel, error channel
	 * We will parse the comments of the source file
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

	//invoking the go routine to process the comment refactoring through the channels
	go processComments(file, in, out, errCh, f, fset)

	return in, errCh, nil
}

func processComments(file string, in chan string, out chan string, erCh chan error, f *ast.File, fset *token.FileSet) {
	/*
	 * We will iterate through the comments
	 * Will send them to the output channel
	 * Will wait for the input channel to get the refactored string
	 * Writes nil to the error channel
	 * Finally overwrite the existing file
	 */
	//iterating through the comments and store it in imports array
	for _, v := range f.Comments {
		for _, c := range v.List {
			//Sending them to the output channel
			out <- c.Text

			//waiting for the input channel to give the refactored comments
			c.Text = <-in

			//outputing nil error to the err channel
			erCh <- nil
		}
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
