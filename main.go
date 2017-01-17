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
	plugin.StartCollector(couchbase.CouchBasePlugin{}, Name, Version)
}
