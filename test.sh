#!/bin/bash
snapctl plugin load ~/bin/snap-plugin-collector-couchbase
snapctl plugin load ~/src/snap/plugin/snap-plugin-publisher-file
snapctl plugin load ~/src/snap/plugin/snap-processor-passthru

snapctl task create -t mock-file.yaml 

taskid="$(snapctl task list | cut -f 1 | grep -v ID)"
echo $taskid
snapctl task watch $taskid
