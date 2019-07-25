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
	//AUTHOR_NAME denotes the string to be replaced by the author name in the boilerplate code
	AUTHOR_NAME = "{{.Author.Name}}"
	//AUTHOR_EMAIL denotes the string to be replaced by the author email in the boilerplate code
	AUTHOR_EMAIL = "{{.Author.Email}}"
	//PACKAGE denotes the string to be replaced by the package in the boilerplate code
	PACKAGE = "{{.Package}}"
)

const (
	//MAIN_PROJECT_NAME is the refactor name for the project name in main.go file
	MAIN_PROJECT_NAME = "Main project name"
	//MAIN_PROJECT_DESCRIPTION is the refactor name of the project description in main.go file
	MAIN_PROJECT_DESCRIPTION = "Main project description"
	//BOILER_PLATE_PACKAGE is the boiler plate package name in the import packages
	BOILER_PLATE_PACKAGE = "Boiler plate package name"
	//README_PROJECT_NAME is the refactor name for the project name in README.md file
	README_PROJECT_NAME = "README project name"
	//README_PROJECT_DESCRIPTION is the refactor description for the project description in README.md file
	README_PROJECT_DESCRIPTION = "README project description"
	//README_PACKAGE is the refactor package for the package in README.md file
	README_PACKAGE = "README package"
	//README_AUTHOR_NAME is the refactor author name for the author name in README.md file
	README_AUTHOR_NAME = "README author name"
	//README_AUTHOR_EMAIL is the refactor author email for the author email in README.md file
	README_AUTHOR_EMAIL = "README project email"
)
