package main

import (
	"github.com/Staples-Inc/snap-plugin-collector-couchbase/couchbase"
	"github.com/intelsdi-x/snap/control/plugin"
	"log"
	"os"
)

func main() {
	// Three things are provided:
	// - The definition of the plugin metadata
	// - The implementation satisfying plugin.PublisherPlugin
	// - The publisher config policy satisfying plugin.ConfigRules

	// Define metadata about the plugin
	meta := couchbase.Meta()

	log.Println(meta)
	plugin.Start(meta, couchbase.New(), os.Args[1])

}
