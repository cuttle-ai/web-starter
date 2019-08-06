// Copyright 2019 Cuttle.ai. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//Package project has the definitions and implementations required for setting up the project.
//It copies the boilerplate code to the required destination
package project

import (
	"fmt"

	"github.com/cuttle-ai/web-starter/generate"
)

/*
 * This file contians the project defintions and list the refactors required for
 * setting up the boilerplat project
 */

//LicenseType is the type of license
type LicenseType string

const (
	//AGPL3 type license
	AGPL3 LicenseType = "AGPL-3"
	//BSD2 type license
	BSD2 LicenseType = "BSD-2"
	//BSD3 type license
	BSD3 LicenseType = "BSD-3"
	//CLOSED type license
	CLOSED LicenseType = "CLOSED"
	//GPL2 type license
	GPL2 LicenseType = "GPL-2"
	//GPL3 type license
	GPL3 LicenseType = "GPL-3"
	//MIT type license
	MIT LicenseType = "MIT"
	//UNLICENSED type license
	UNLICENSED LicenseType = "UNLICENSED"
)

//License gives info about the license
type License struct {
	//Type is the type of license
	Type LicenseType
	//Year of the license validity
	Year string
	//Organisation that issed the license
	Organisation string
}

//Project struct lists the information about the project
type Project struct {
	//Name of the project
	Name string
	//Description of the project
	Description string
	//Author information for the project
	Author Author
	//Destination target for setting up the boilerplate code
	Destination string
	//Package is the package path to be used by the project
	Package string
	//Sources is the list of sources with refactors in the boilerplate code
	Sources []generate.Source
	//License is the license to be provided for the project
	License License
}

//Author refers to the initial project author
type Author struct {
	//Name of the author
	Name string
	//Email of the author
	Email string
}

//String is the stringer implementation of the author
func (a Author) String() string {
	return a.Name + "<" + a.Email + ">"
}

//Setup will setup a project with boilerplate code.
func (p Project) Setup() error {
	/*
	 * We will init the project sources
	 * We will iterate through the sources
	 * Then will generate the code
	 */
	(&p).InitSources()
	for _, v := range p.Sources {
		err := v.Generate(p.Destination)
		if err != nil {
			//Error while generating the source
			fmt.Println("Error while generating the source while setting up", (&v).Name(), "in the project", p.Name)
			return err
		}
	}
	return nil
}
