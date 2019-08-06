// Copyright 2019 Cuttle.ai. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package project

/*
 * This file contains the constants used for creating the project out of boilerplate code
 */

const (
	//ProjectName denotes the string to be replaced by the project name in the boilerplate code
	ProjectName = "{{.Name}}"
	//ProjectDescription denotes the string to be replaced by the project description in the boilerplate code
	ProjectDescription = "{{.Description}}"
	//AuthorName denotes the string to be replaced by the author name in the boilerplate code
	AuthorName = "{{.Author.Name}}"
	//AuthorEmail denotes the string to be replaced by the author email in the boilerplate code
	AuthorEmail = "{{.Author.Email}}"
	//Package denotes the string to be replaced by the package in the boilerplate code
	Package = "{{.Package}}"
	//CuttleAI denotes the string to be replaced by the author across the lICENSE information in all files
	CuttleAI = "Cuttle.ai"
	//MitStyle denotes the string to be replaced by the LICENSE Template
	MitStyle = "MIT-style"
	//Year denotes the string to be replaced by the year in LICENSE files
	Year = "{{.Year}}"
	//Organisation denotes the string to be replaced by the orgianisation in LICENSE files
	Organisation = "{{.Organisation}}"
)

const (
	//MainProjectName is the refactor name for the project name in main.go file
	MainProjectName = "Main project name"
	//MainProjectDescription is the refactor name of the project description in main.go file
	MainProjectDescription = "Main project description"
	//BoilerPlatePackage is the boiler plate package name in the import packages
	BoilerPlatePackage = "Boiler plate package name"
	//ReadmeProjectName is the refactor name for the project name in README.md file
	ReadmeProjectName = "README project name"
	//ReadmeProjectDescription is the refactor description for the project description in README.md file
	ReadmeProjectDescription = "README project description"
	//ReadmePackage is the refactor package for the package in README.md file
	ReadmePackage = "README package"
	//ReadmeAuthorName is the refactor for author name for the author name in README.md file
	ReadmeAuthorName = "README author name"
	//ReadmeAuthorEmail is the refactor for author email for the author email in README.md file
	ReadmeAuthorEmail = "README project email"
	//LicenseOrganisation is the refactor for organisation in the license files and license disclaimer in all the files
	LicenseOrganisation = "LICENSE organisation"
	//LicenseTemplate is the refactor for LICENSE type like MIT-style, closed etc.
	LicenseTemplate = "LICENSE template"
	//LicenseYear is the refactor for year in the license file
	LicenseYear = "LICENSE year"
	//LicenseProjectName is the refactor for project name in the license file
	LicenseProjectName = "LICENSE project name"
)
