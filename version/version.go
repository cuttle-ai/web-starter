// Copyright 2019 Cuttle.ai. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//Package version has the version information about the application
package version

//Version represent a version
type Version struct {
	//CodeName is the code name for the version
	CodeName string
	//Version is the version string of the tool
	Version string
	//Value is the integer value to be used for comparison with other version.
	//This comes handy while developing the migration features
	Value int8
}

var (
	//V1 is the version 1 of the tool
	V1 = Version{
		"TheBuild",
		"v1.0.0",
		0,
	}
)

//Default is the default version of the tool
var Default = V1
