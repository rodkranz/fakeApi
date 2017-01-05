// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package fakeApi

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"gopkg.in/macaron.v1"

	"github.com/rodkranz/fakeApi/modules/log"
)

type ApiFakeOptions struct {
	DefaultApi string
	BaseFolder string
}

type ApiFake struct {
	*macaron.Context
	Default string
	Folder  string
	Domain  string
	Delay   int
	Headers map[string]string
}

// GetMethodAndStatusCode return method, status code and bool if found or not
func (a *ApiFake) GetMethodAndStatusCode() (string, int, bool) {
	// Get status code and method if it doesn't exist get random
	status, has := a.getHeaderData("X-Fake-Response-Code")
	if !has {
		return a.Context.Req.Method, a.Context.Resp.Status(), false
	}

	i, err := strconv.ParseInt(status, 10, 32)
	if err != nil {
		return a.Context.Req.Method, a.Context.Resp.Status(), false
	}

	return a.Context.Req.Method, int(i), true
}

// GetSeedPath returns the path of seed file.
func (a *ApiFake) GetSeedPath(seed string) (string, error) {
	seed = strings.Replace(seed, "__", "#", -1)
	seed = strings.Replace(seed, "/", "_", -1)
	seed = strings.Replace(seed, "#", "_", -1)

	filePath := fmt.Sprintf("%v/%v.json", a.Folder, seed)
	if isNotExist(filePath) && !isNotExist(fmt.Sprintf("%v/%v.json", a.Default, seed)) {
		filePath = fmt.Sprintf("%v/%v.json", a.Default, seed)
	}

	return filePath, nil
}

// RegisterDomain change domain 'default' to 'X-Fake-Domain'
// if the folder exists with same name of 'X-Fake-Domain'
// it will use this folder as base
func (a *ApiFake) registerDomain() {
	if domain, has := a.getHeaderData("X-Fake-Domain"); has {
		a.Domain = domain
	}

	folder := fmt.Sprintf("%s/%s", a.Folder, a.Domain)
	if isNotExist(folder) {
		return
	}

	a.Folder = folder
}

// RegisterDelay if has header for dela register at fakeApi
func (a *ApiFake) registerDelay() {
	delay, has := a.getHeaderData("X-Fake-Delay")
	if !has {
		return
	}

	// try to convert of string to int64 if has error keep 0 delay
	i, err := strconv.ParseInt(delay, 10, 32)
	if err != nil {
		return
	}

	a.Delay = int(i)
}

// getHeaderData returns first data from header
func (a *ApiFake) getHeaderData(name string) (string, bool) {
	values, has := a.Context.Req.Header[name]
	if has && len(values) > 0 {

		if strings.Contains(values[0], ",") {
			values = strings.Split(values[0], ",")
		}

		return values[0], true
	}

	return "", false
}

// isNotExists return if file exists or not
func isNotExist(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}

// Register Register fake api service
func Register(opt ApiFakeOptions) macaron.Handler {

	// Check if fakes folder exist
	if isNotExist(opt.BaseFolder) {
		if err := os.Mkdir(opt.BaseFolder, 0755); err != nil {
			log.Fatal(4, "Please create a 'fakes' folder: [%v]", opt.BaseFolder)
			log.Fatal(4, "Please create a 'fakes' folder: [%v]", opt.BaseFolder)
		}
	}

	// Check if default folder in faker exist
	p := path.Join(opt.BaseFolder, opt.DefaultApi)
	if isNotExist(p) {
		if err := os.Mkdir(path.Join(p), 0755); err != nil {
			log.Fatal(4, "Please create a 'default' folder: [%v]", p)
		}
	}

	return func(ctx *macaron.Context) {
		api := &ApiFake{
			Delay:   0,
			Domain:  "default",
			Default: opt.DefaultApi,
			Folder:  opt.BaseFolder,
			Context: ctx,
		}

		// AutoConfig load method itself
		api.registerDomain()
		api.registerDelay()

		// Share FakeApi Module for all handlers
		ctx.Map(api)

		// Execute handlers
		ctx.Next()

		// Apply delay
		time.Sleep(time.Duration(api.Delay) * time.Millisecond)
	}
}
