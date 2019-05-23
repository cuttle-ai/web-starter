// Copyright 2019 Cuttle.ai. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routes

import (
	"context"
	"net/http"

	"github.com/cuttle-ai/web-starter/version"

	"github.com/cuttle-ai/web-starter/boilerplate/config"
)

/*
 * This file has the definition of route data structure
 */

//HandlerFunc is the Handler func with the context
type HandlerFunc func(context.Context, http.ResponseWriter, *http.Request)

//Route is a route with explicit versions
type Route struct {
	//Version is the version of the route
	Version string
	//Pattern is the url pattern of the route
	Pattern string
	//HandlerFunc is the handler func of the route
	HandlerFunc HandlerFunc
}

//Register registers the route with the default http handler func
func (r Route) Register(s *http.ServeMux) {
	/*
	 * If the route version is default version then will register it without version string to http handler
	 * Will register the router with the http handler
	 */
	if r.Version == version.Default {
		s.Handle(r.Pattern, http.TimeoutHandler(r, config.ResponseTimeout, "timeout"))
	}
	s.Handle("/"+r.Version+r.Pattern, http.TimeoutHandler(r, config.ResponseTimeout, "timeout"))
}

//ServeHTTP implements HandlerFunc of http package. It makes use of the context of request
func (r Route) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	/*
	 * Will get the context
	 * Will parse the form
	 * Execute request handler func
	 */
	ctx := req.Context()
	req.ParseForm()
	r.Exec(ctx, res, req)
}

//Exec will execute the handler func. By default it will set response content type as as json.
//It will also cancel the context at the end. So no need of explicitly invoking the same in the handler funcs
func (r Route) Exec(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	/*
	 * Will get the cancel for the context
	 * Will set the content type of response as json
	 * Will execute the handlerfunc
	 * Cancelling the context at the end
	 */
	//getting the context cancel
	c, cancel := context.WithCancel(ctx)

	//setting the content type as json
	res.Header().Set("Content-Type", "application/json")

	//executing the handler
	r.HandlerFunc(c, res, req)

	//cancelling the context
	cancel()
}
