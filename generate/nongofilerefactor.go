// Copyright 2019 Cuttle.ai. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package generate

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
 * This file has the defintions of the non-go file refactor struct which implements the RefactorSource.
 */

//NonGoFileRefactor does refactor non-go files
type NonGoFileRefactor struct{}

//NewNonGoFileRefactor is the constructor for the non go file refactor
func NewNonGoFileRefactor() NonGoFileRefactor {
	return NonGoFileRefactor{}
}

//Initiate will initiate the refactoring of content in non-go source file name provided
//It will send the content through the out channel and expects the refactored input for each import send in the output to come
//If any error happens in the process it will send through the error channel.
//After ranging over the out channel check for the last error from the error channel so that error while
//finishing the refactoring is captured
func (p NonGoFileRefactor) Initiate(file string, out chan string) (chan string, chan error, error) {
	/*
	 * We will create an input channel, error channel
	 * Will initiate a go routine that processes the content and send it to the out channel
	 */
	//creating the input and error output channel
	in, errCh := make(chan string), make(chan error)

	//invoking the go routine to process the content refactoring through the channels
	go processContent(file, in, out, errCh)

	return in, errCh, nil
}

func processContent(file string, in chan string, out chan string, erCh chan error) {
	/*
	 * We will open the file
	 * Scan it line by line
	 * Will send them to the output channel
	 * Will wait for the input channel to get the refactored string
	 * Writes nil to the error channel
	 * Finally overwrite the existing file
	 */
	//opening the file
	f, err := os.Open(file)
	if err != nil {
		//Error while opening the given file
		fmt.Println("Error while opening the file")
		close(out)
		erCh <- err
		return
	}

	//Going to scan the file line by line
	scanner := bufio.NewScanner(f)
	sb := strings.Builder{}
	for scanner.Scan() {
		//Sending them to the output channel
		out <- scanner.Text()

		//waiting for the input channel to give the refactored string
		nS := <-in
		fmt.Fprintln(&sb, nS)

		//outputing nil error to the err channel
		erCh <- nil
	}
	close(out)
	f.Close()

	//overwrting the existing file
	c, _ := os.Create(file)
	defer c.Close()
	//we are ignoring the error, since it is safe to do so in this context
	//we are writing the refactored text into the file itself
	_, err = c.WriteString(sb.String())

	erCh <- err
	close(erCh)
}
