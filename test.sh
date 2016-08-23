#!/bin/bash
./snapctl plugin load ~/bin/snap-plugin-collector-couchbase
./snapctl plugin load ~/src/github.com/intelsdi-x/snap/build/plugin/snap-plugin-publisher-mock-file
./snapctl plugin load ~/src/github.com/intelsdi-x/snap/build/plugin/snap-plugin-processor-passthru

./snapctl task create -t ./examples/tasks/mock-file.yaml 

taskid="$(./snapctl task list | cut -f 1 | grep -v ID)"
echo $taskid
./snapctl task watch $taskid
