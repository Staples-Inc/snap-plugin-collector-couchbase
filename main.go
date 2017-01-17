package main

import (
	"github.com/Staples-Inc/snap-plugin-collector-couchbase/couchbase"
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
)

const (
	Name    = "couchbase"
	Version = 1
)

func main() {
	cbp, err := couchbase.New()
	if err != nil {
		panic(err)
	}

	plugin.StartCollector(cbp, Name, Version)
}
