// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package settings

import (
	"encoding/json"
)

type Config struct {
	Title       string
	CrossDomain bool
	Eps         []*Endpoints
}


func (c *Config) RegisterEndpoints(eps *Endpoints) {
	c.Eps = append(c.Eps, eps)
}


type Endpoints struct {
	Url       string
	Endpoints Endpoint
}

///
type Endpoint map[string]interface{}
func (c *Endpoint) Unmarshal(b []byte) error {
	return json.Unmarshal(b, c)
}
