// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cmd

import (
	"gopkg.in/urfave/cli.v2"
	"fmt"
)

var Docs = &cli.Command{
	Name:        "docs",
	Usage:       "Run Server with docs of seed files",
	Description: `Start server and generate docs for seed files.`,
	Action:      runDocs,
	Flags:       []cli.Flag{},
}

func runDocs(c *cli.Context) error {
	fmt.Fprint(c.App.Writer, "Not implemented yet.")
	return nil
}

