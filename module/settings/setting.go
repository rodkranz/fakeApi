// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package settings

import "github.com/rodkranz/fakeApi/module/entity"

var (
	CrossDomain bool
	Title       string
	Folder      string
	Urls        map[string]entity.Endpoint
)

func init() {
	CrossDomain = true
	Title = "GoLang"
	Folder = "json"
	Urls = make(map[string]entity.Endpoint, 0)
}


//func LoadInfo(p string) error {
//	data, err := ioutil.ReadFile(p)
//	if err != nil {
//		return err
//	}
//
//	ep := &Endpoint{}
//	if err = ep.Unmarshal(data); err != nil {
//		log.Fatalf("Something is worng with settings: %v", err.Error())
//		os.Exit(1)
//	}
//
//	name := path.Base(p)
//	name = strings.Replace(name, path.Ext(p), "", -1)
//	name = strings.Replace(name, "_", "/", -1)
//
//	eps := &Endpoints{
//		Url:       fmt.Sprintf("/%v", name),
//		Endpoints: *ep,
//	}
//
//	Api.RegisterEndpoints(eps)
//	return nil
