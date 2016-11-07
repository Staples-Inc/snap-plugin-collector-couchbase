package main

import (
	"github.com/Staples-Inc/snap-plugin-collector-couchbase/couchbase"
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
)

func main() {
	plugin.StartCollector(couchbase.New(), couchbase.Name, couchbase.Version)
}
