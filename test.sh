#!/bin/bash
snapctl plugin load ~/bin/snap-plugin-collector-couchbase
snapctl plugin load ~/src/snap/plugin/snap-plugin-publisher-file
snapctl plugin load ~/src/snap/plugin/snap-processor-passthru
