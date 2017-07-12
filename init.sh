#!/bin/bash/

go-bindata -o "./modules/bindata/bindata.go" -pkg "bindata" conf/*

bra run