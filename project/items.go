// Copyright 2019 Cuttle.ai. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package project

import (
	"go/build"
	"os"
	"path/filepath"

	"github.com/cuttle-ai/web-starter/generate"
)

/*
 * This file contains the list of boilerplate code with required refactors
 */

//Separator based on the os
var Separator = string([]rune{filepath.Separator})

//PackagePath is the package name of web starter project
var PackagePath = "github.com" + Separator + "cuttle-ai" + Separator + "web-starter"

//BoilerplatePackagePath is the package path of the boilr plate code
var BoilerplatePackagePath = PackagePath + Separator + "boilerplate"

//GoPath is the gopath in the system
func GoPath() string {
	/*
	 * We will try to get the gopath from the environment variable
	 * If not available we will get it from the default go path
	 */
	p := os.Getenv("GOPATH")
	if len(p) == 0 {
		p = build.Default.GOPATH
	}
	return p + Separator + "src" + Separator
}

//LicensesPath is the location where all license templates are kept
var LicensesPath = GoPath() + PackagePath + Separator + "licenses"

//BoilerplatePath is the absolute path to the boilerplate code
var BoilerplatePath = GoPath() + BoilerplatePackagePath

//VersionPath is the path of the version package in the boilerplate code
var VersionPath = BoilerplatePath + Separator + "version"

//ConfigPath is the path of the config package in the boilerplate code
var ConfigPath = BoilerplatePath + Separator + "config"

//LogPath is the path of the log package in the boilerplate code
var LogPath = BoilerplatePath + Separator + "log"

//RoutesPath is the path of the routes package in the boilerplate code
var RoutesPath = BoilerplatePath + Separator + "routes"

//ResponsePath is the path of the response package in the boilerplate code
var ResponsePath = RoutesPath + Separator + "response"

//LicenseRefactors returns the license refactors to be done in every go source file
func (p *Project) LicenseRefactors() []generate.Refactor {
	return []generate.Refactor{
		{
			Name:    LicenseOrganisation,
			Find:    CuttleAI,
			Replace: p.Author.String(),
			Source:  generate.NewCommentRefactor(),
		},
		{
			Name:    LicenseTemplate,
			Find:    MitStyle,
			Replace: p.Author.String(),
			Source:  generate.NewCommentRefactor(),
		},
	}
}

//BoilerplatePackageRefactors returns the package name refactor to be done in every go source file
func (p *Project) BoilerplatePackageRefactors() generate.Refactor {
	return generate.Refactor{
		Name:    BoilerPlatePackage,
		Find:    BoilerplatePackagePath,
		Replace: p.Package,
		Source:  generate.NewPackageRefactor(),
	}
}

//InitSources will init the sources of the project
func (p *Project) InitSources() {
	p.Sources = []generate.Source{
		{
			Path:     BoilerplatePath,
			FileName: "main.go",
			Refactors: append([]generate.Refactor{
				{
					Name:    MainProjectName,
					Find:    ProjectName,
					Replace: p.Name,
					Source:  generate.NewCommentRefactor(),
				},
				{
					Name:    MainProjectDescription,
					Find:    ProjectDescription,
					Replace: p.Description,
					Source:  generate.NewCommentRefactor(),
				},
				p.BoilerplatePackageRefactors(),
			}, p.LicenseRefactors()...),
		},
		{
			Path:      BoilerplatePath,
			FileName:  ".gitignore",
			Refactors: []generate.Refactor{},
		},
		{
			Path:     BoilerplatePath,
			FileName: "README.md",
			Refactors: []generate.Refactor{
				{
					Name:    ReadmeProjectName,
					Find:    ProjectName,
					Replace: p.Name,
					Source:  generate.NewNonGoFileRefactor(),
				},
				{
					Name:    ReadmeProjectDescription,
					Find:    ProjectDescription,
					Replace: p.Description,
					Source:  generate.NewNonGoFileRefactor(),
				},
				{
					Name:    ReadmePackage,
					Find:    Package,
					Replace: p.Package,
					Source:  generate.NewNonGoFileRefactor(),
				},
				{
					Name:    ReadmeAuthorName,
					Find:    AuthorName,
					Replace: p.Author.Name,
					Source:  generate.NewNonGoFileRefactor(),
				},
				{
					Name:    ReadmeAuthorEmail,
					Find:    AuthorEmail,
					Replace: p.Author.Email,
					Source:  generate.NewNonGoFileRefactor(),
				},
			},
		},

		{
			Path:     LicensesPath + Separator + string(p.License.Type),
			FileName: "LICENSE",
			Refactors: []generate.Refactor{
				{
					Name:    LicenseProjectName,
					Find:    ProjectName,
					Replace: p.Name,
					Source:  generate.NewNonGoFileRefactor(),
				},
				{
					Name:    LicenseOrganisation,
					Find:    Organisation,
					Replace: p.License.Organisation,
					Source:  generate.NewNonGoFileRefactor(),
				},
				{
					Name:    LicenseYear,
					Find:    Year,
					Replace: p.License.Year,
					Source:  generate.NewNonGoFileRefactor(),
				},
			},
		},
		{
			Path:                VersionPath,
			FileName:            "version.go",
			RelativeDestination: "version",
			Refactors:           []generate.Refactor{},
		},
		{
			Path:                ConfigPath,
			FileName:            "config.go",
			RelativeDestination: "config",
			Refactors:           []generate.Refactor{},
		},
		{
			Path:                LogPath,
			FileName:            "log.go",
			RelativeDestination: "log",
			Refactors: []generate.Refactor{
				p.BoilerplatePackageRefactors(),
			},
		},
		{
			Path:                RoutesPath,
			FileName:            "routes.go",
			RelativeDestination: "routes",
			Refactors:           []generate.Refactor{},
		},
		{
			Path:                RoutesPath,
			FileName:            "route.go",
			RelativeDestination: "routes",
			Refactors: []generate.Refactor{
				p.BoilerplatePackageRefactors(),
			},
		},
		{
			Path:                RoutesPath,
			FileName:            "example_test.go",
			RelativeDestination: "routes",
			Refactors: []generate.Refactor{
				p.BoilerplatePackageRefactors(),
			},
		},
		{
			Path:                ResponsePath,
			FileName:            "response.go",
			RelativeDestination: "routes" + Separator + "response",
			Refactors: []generate.Refactor{
				p.BoilerplatePackageRefactors(),
			},
		},
	}
}
