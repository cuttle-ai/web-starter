// Copyright 2019 Cuttle.ai. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package generate

import (
	"errors"
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

/*
 * This file has the defintions of the package name refactor struct which implements the RefactorSource.
 */

//PackageRefactor is the to refactor the import package names present in a source code file
type PackageRefactor struct {
	//Seperator is the separator used to seperate th import path strings.
	//Refer the constructor for the default of the separator
	Separator string
}

//NewPackageRefactor is the constructor for the package import name refactor
//Default value of the separator is |
func NewPackageRefactor() PackageRefactor {
	return PackageRefactor{
		Separator: "|",
	}
}

//Src will return the import packages corresponding to the go source file name provided
//Will return any error occurred during opening the file or reading the source code
func (p PackageRefactor) Src(file string) (string, error) {
	/*
	 * We will parse the imports of the source file
	 * Will iterate through the imports and join them to produce the file result
	 * Will return the path value of the imports seperated by the seperator
	 */
	//imports to store the list of all the import paths
	imports := []string{}

	//fileset to trcak the source code positions
	fset := token.NewFileSet()

	//parsing the source for the given file
	f, err := parser.ParseFile(fset, file, nil, parser.ImportsOnly)
	if err != nil {
		//error while parsing the source file
		fmt.Println("Error while parsing the go source file", file)
		return "", err
	}

	//iterating through the import paths and store it in imports array
	for _, v := range f.Imports {
		imports = append(imports, v.Path.Value)
	}

	//join the imports and return the final result
	return strings.Join(imports, p.Separator), nil
}

//SetSrc will set the given code as the import packages section for the given file name
func (p PackageRefactor) SetSrc(file, code string) error {
	/*
	 * We will split the given code by the separator
	 * Then will parse the source code file provided
	 * If the number of imports are different it will return an error
	 * Will replace the imports with the values of the given code
	 * Then will write the source code to the given file
	 */
	//splitting the given source code by the separator
	imports := strings.Split(code, p.Separator)

	//parsing the provided source code file
	//fileset to track the source positions
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
	if err != nil {
		//error while parsing the given source code
		fmt.Println("Error while parsing the go source file", file)
		return err
	}

	//checking whether the no of imports is different than the file
	if len(imports) != len(f.Imports) {
		//imports are different
		fmt.Println("The no of refactored imports is different from that of the original source file")
		return errors.New("Couldn't set the refactored imports to the file " + file + " due to the difference in the no. of imports")
	}

	//iterating through the imports and replacing them
	for k, v := range imports {
		f.Imports[k].Path.Value = v
	}

	//writing the refactored source code to the file
	c, _ := os.Create(file)
	defer c.Close()
	//we are ignoring the error, since it is safe to do so in this context
	//we are writing the formatted code into the file itself
	err = format.Node(c, fset, f)
	if err != nil {
		//error while writing the refactored source code to the file
		fmt.Println("Error while writing the formatted code to the given file", file)
		return err
	}
	return nil
}
