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

var separator = string([]rune{filepath.Separator})

//PackagePath is the package name of web starter project
var PackagePath = "github.com" + separator + "cuttle-ai" + separator + "web-starter"

//BoilerplatePackagePath is the package path of the boilr plate code
var BoilerplatePackagePath = PackagePath + separator + "boilerplate"

//BoilerplatePath is the absolute path to the boilerplate code
var BoilerplatePath = func() string {
	/*
	 * We will try to get the gopath from the environment variable
	 * If not available we will get it from the default go path
	 */
	p := os.Getenv("GOPATH")
	if len(p) == 0 {
		p = build.Default.GOPATH
	}
	return p + separator + "src" + separator + BoilerplatePackagePath
}()

//VersionPath is the path of the version package in the boilerplate code
var VersionPath = BoilerplatePath + separator + "version"

//ConfigPath is the path of the config package in the boilerplate code
var ConfigPath = BoilerplatePath + separator + "config"

//LogPath is the path of the log package in the boilerplate code
var LogPath = BoilerplatePath + separator + "log"

//RoutesPath is the path of the routes package in the boilerplate code
var RoutesPath = BoilerplatePath + separator + "routes"

//ResponsePath is the path of the response package in the boilerplate code
var ResponsePath = RoutesPath + separator + "response"

//InitSources will init the sources of the project
func (p *Project) InitSources() {
	p.Sources = []generate.Source{
		{
			Path:     BoilerplatePath,
			FileName: "main.go",
			Refactors: []generate.Refactor{
				{
					Name:    MAIN_PROJECT_NAME,
					Find:    PROJECT_NAME,
					Replace: p.Name,
					Source:  generate.NewCommentRefactor(),
				},
				{
					Name:    MAIN_PROJECT_DESCRIPTION,
					Find:    PROJECT_DESCRIPTION,
					Replace: p.Description,
					Source:  generate.NewCommentRefactor(),
				},
				{
					Name:    BOILER_PLATE_PACKAGE,
					Find:    BoilerplatePackagePath,
					Replace: p.Package,
					Source:  generate.NewPackageRefactor(),
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
				{
					Name:    BOILER_PLATE_PACKAGE,
					Find:    BoilerplatePackagePath,
					Replace: p.Package,
					Source:  generate.NewPackageRefactor(),
				},
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
				{
					Name:    BOILER_PLATE_PACKAGE,
					Find:    BoilerplatePackagePath,
					Replace: p.Package,
					Source:  generate.NewPackageRefactor(),
				},
			},
		},
		{
			Path:                RoutesPath,
			FileName:            "example_test.go",
			RelativeDestination: "routes",
			Refactors: []generate.Refactor{
				{
					Name:    BOILER_PLATE_PACKAGE,
					Find:    BoilerplatePackagePath,
					Replace: p.Package,
					Source:  generate.NewPackageRefactor(),
				},
			},
		},
		{
			Path:                ResponsePath,
			FileName:            "response.go",
			RelativeDestination: "routes" + separator + "response",
			Refactors: []generate.Refactor{
				{
					Name:    BOILER_PLATE_PACKAGE,
					Find:    BoilerplatePackagePath,
					Replace: p.Package,
					Source:  generate.NewPackageRefactor(),
				},
			},
		},
	}
}
