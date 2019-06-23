// Copyright 2019 Cuttle.ai. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package generate

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

/*
 * This file has the definitions of the boiler plate code source file
 */

//Source holds the source file and list of refactors to be done on it
type Source struct {
	//Path is absolute path at which the source file exists
	Path string
	//FileName is the name of the source file
	FileName string
	//Refactors is the list of refactors
	Refactors []Refactor
}

//Name returns the name of the source
func (s *Source) Name() string {
	return s.Path + string([]rune{filepath.Separator}) + s.FileName
}

//Generate will generate a source file in the given destination path.
//It will copy the source file and makes the required refactors in the
//newly created destination file.
//The required destination directory has to be provided as argument. The file name will
//be same as the source. If the file name also has to be changed, it has to be specified under the
//refactor list.
func (s *Source) Generate(dst string) error {
	/*
	 * We will copy the source file to the destination
	 * We will open the destination file
	 * Will make all the refactors
	 */
	//copying the source file
	d, err := s.Copy(dst)
	if err != nil {
		//Error while copying the file to the destination
		fmt.Println("Error while generating the source file", s.Name())
		return err
	}

	//will iterate through the refactors and do them
	for _, v := range s.Refactors {
		err = v.Do(d)
		if err != nil {
			fmt.Println("Error while making the refactor", v.String(), "in the destination file for", d)
			return err
		}
	}
	return nil
}

//Copy copies a source file to a given destination. The destination shouldn't have the
//destination file name. It should only contain the absolute path to the destination directory.
//If any error occurs while copying, like unsuccessful copying of the file, or unsuccessful creation of the
//destination directory, it will be reported back. It will also return the absolute path to the destination file.
func (s *Source) Copy(dst string) (string, error) {
	/*
	 * First we need to check whether file exists in the given source
	 * Checking whether the file is a regular file itself
	 * Create the directories in destination if not existing
	 * Identify the destination file name
	 * Copy the file to the destination
	 */
	//checking whether the source file exists and is a regular file
	sourceFileStat, err := os.Stat(s.Name())
	if err != nil {
		//checking whether the source file
		fmt.Println("Error while reading info of the source file", s.Name())
		return "", err
	}

	//checking whether the source file is regular or not
	if !sourceFileStat.Mode().IsRegular() {
		fmt.Println("The given source file is not regular", s.Name())
		return "", fmt.Errorf("%s is not a regular file", s.Name())
	}

	//creating the destination directory if not existing
	err = os.MkdirAll(dst, 0775)
	if err != nil {
		//error while creating the destination directories
		fmt.Println("Error while creating the destination directories", dst)
		return "", err
	}

	//identifying the destination filename
	dstF := dst + string([]rune{filepath.Separator}) + s.FileName

	//copying the source file to the destination
	source, _ := os.Open(s.Name())
	//since the stats of the file is already checked, there is less chance that the
	//source cannot be opened. Hence we are ignoring the error in opening the source file
	defer source.Close()
	dF, _ := os.Create(dstF)
	//There is less chance that the destination file cannot be created.
	//Hence we are ignoring the error. If permission issues creep in
	//the application is meant to fail by default. So it should be fine
	defer dF.Close()
	//copying
	io.Copy(dF, source)
	//since the source and destination are meant to exist, there is less chance that the
	//copying file. Permission issues can creep in. In that case the application is meant to fail.

	return dstF, nil
}
