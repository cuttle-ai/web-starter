// Copyright 2019 Cuttle.ai. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//Package config will have necessary configuration for the application
package config

import (
	"os"
	"strconv"
	"time"
)

var (
	//Port in which the application is being served
	Port = "8080"
	//ResponseTimeout of the api to respond in milliseconds
	ResponseTimeout = time.Duration(100 * time.Millisecond)
	//RequestRTimeout of the api request body read timeout in milliseconds
	RequestRTimeout = time.Duration(20 * time.Millisecond)
	//ResponseWTimeout of the api response write timeout in milliseconds
	ResponseWTimeout = time.Duration(20 * time.Millisecond)
)

func init() {
	/*
	 * We will init the port
	 * We will init the request timeout
	 * We will init the request body read timeout
	 * We will init the request body write timeout
	 */
	//port
	if len(os.Getenv("PORT")) != 0 {
		//Assign the default port as 9090
		Port = os.Getenv("PORT")
	}

	//response timeout
	if len(os.Getenv("RESPONSE_TIMEOUT")) != 0 {
		//if sucessfull convert timeout
		if t, err := strconv.ParseInt(os.Getenv("RESPONSE_TIMEOUT"), 10, 64); err == nil {
			ResponseTimeout = time.Duration(t * int64(time.Millisecond))
		}
	}

	//request body read timeout
	if len(os.Getenv("REQUEST_BODY_READ_TIMEOUT")) != 0 {
		//if sucessfull convert timeout
		if t, err := strconv.ParseInt(os.Getenv("REQUEST_BODY_READ_TIMEOUT"), 10, 64); err == nil {
			RequestRTimeout = time.Duration(t * int64(time.Millisecond))
		}
	}

	//response write
	if len(os.Getenv("RESPOSE_WRITE_TIMEOUT")) != 0 {
		//if sucessfull convert timeout
		if t, err := strconv.ParseInt(os.Getenv("RESPOSE_WRITE_TIMEOUT"), 10, 64); err == nil {
			ResponseWTimeout = time.Duration(t * int64(time.Millisecond))
		}
	}
}

var (
	//PRODUCTION is the switch to turn on and off the Production environment.
	//1: On, 0: Off
	PRODUCTION = 0
)

func init() {
	/*
	 * Will init Production switch
	 */
	//Production
	if len(os.Getenv("PRODUCTION")) != 0 {
		//if sucessfull convert production
		if t, err := strconv.Atoi(os.Getenv("PRODUCTION")); err == nil && (t == 1 || t == 0) {
			PRODUCTION = t
		}
	}
}
