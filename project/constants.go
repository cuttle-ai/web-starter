// Copyright 2019 Cuttle.ai. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package project

/*
 * This file contains the constants used for creating the project out of boilerplate code
 */

const (
	//PROJECT_NAME denotes the string to be replaced by the project name in the boilerplate code
	PROJECT_NAME = "{{.Name}}"
	//PROJECT_DESCRIPTION denotes the string to be replaced by the project description in the boilerplate code
	PROJECT_DESCRIPTION = "{{.Description}}"
)

const (
	//MAIN_PROJECT_NAME is the refactor name for the project name in main.go file
	MAIN_PROJECT_NAME = "Project name"
	//MAIN_PROJECT_DESCRIPTION is the refactor name of the project description in main.go file
	MAIN_PROJECT_DESCRIPTION = "Project Description"
	//BOILER_PLATE_PACKAGE is the boiler plate package name in the import packages
	BOILER_PLATE_PACKAGE = "Boiler plate package name"
)
