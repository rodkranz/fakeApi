// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package api

import (
	"net/http"

	"github.com/rodkranz/fakeApi/module/context"
	"github.com/rodkranz/fakeApi/module/fakeApi"
	"github.com/rodkranz/fakeApi/router/common"
)

func ApiDocs(ctx *context.APIContext, fakeApi *fakeApi.ApiFake) {
	docs := &common.Docs{
		Domain: fakeApi.Domain,
		Path:   fakeApi.Folder,
	}
	docs.LoadSeeds()

	status := http.StatusOK

	ctx.JSON(status, map[string]interface{}{
		"message":  "List of endpoints available for this domain",
		"status":   status,
		"domain":   fakeApi.Domain,
		"resource": docs,
	})
}
