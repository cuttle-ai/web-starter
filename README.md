# web-starter [![Go Report Card](https://goreportcard.com/badge/github.com/cuttle-ai/web-starter)](https://goreportcard.com/report/github.com/cuttle-ai/web-starter) [![Build Status](https://ci.cuttle.ai/api/badges/cuttle-ai/web-starter/status.svg)](https://ci.cuttle.ai/cuttle-ai/web-starter) [![Codacy Badge](https://api.codacy.com/project/badge/Grade/feb19038b3dc4b4f86b4f6a16a28e581)](https://www.codacy.com/app/melvinodsa/web-starter?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=cuttle-ai/web-starter&amp;utm_campaign=Badge_Grade) [![Codacy Badge](https://api.codacy.com/project/badge/Coverage/feb19038b3dc4b4f86b4f6a16a28e581)](https://www.codacy.com/app/melvinodsa/web-starter?utm_source=github.com&utm_medium=referral&utm_content=cuttle-ai/web-starter&utm_campaign=Badge_Coverage) [![GoDoc](https://img.shields.io/badge/godoc-Documentation-blue)](https://godoc.org/github.com/cuttle-ai/web-starter)
A boilerplate generator for web application

Its build on top of [cobra](https://github.com/spf13/cobra) and [go-input](https://github.com/tcnksm/go-input) for cmd line interaction

## Pre-requisite
* [Go](https://golang.org/)

## Usage
```sh
$ go get -u github.com/cuttle-ai/web-starter
$ web-starter web-server generate
# Will ask few questions regarding the project name , description etc
Project name
Enter a value (Default is Web Server):

Project description
Enter a value (Default is Backend server):

Author name
Enter a value (Default is cuttle.ai):

Author email
Enter a value (Default is hi@cuttle.ai):

Project destination
Enter a value (Default is /home/melvin/go/src/github.com/hi/web-server):

Package name
Enter a value (Default is github.com/hi/web-server):

Type of license

1. AGPL-3
2. BSD-2
3. BSD-3
4. CLOSED
5. GPL-2
6. MIT
7. UNLICENSED

Enter a number (Default is 6):

Copyright year
Enter a value (Default is 2019):

Organisation
Enter a value (Default is Cuttle.ai):

Installing Web Server

# Now we generated and installed the project with name web-server. We can run it by the following command
$ web-server
2019/08/07 23:15:13 INFO: Starting the server at :8080
```
## Help
```sh
web-starter help
```