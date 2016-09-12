# Snap Couchbase Collector Plugin
Intelsdi-x [snap](http://github.com/intelsdi-x/snap) plugin for couchbase bucket monitoring. It uses couchbase's provided RESTAPI. Read more about bucket monitoring in this [blog post](http://blog.couchbase.com/monitoring-couchbase-cluster) from couchbase.

### Snap version requires at least
v0.14.0-beta

### Snap version tested up to
v0.15.0-beta

### Supported platforms
darwin-amd64
linux-amd64

### Contributor
Staples, Inc.

### License
Snap is Open Source software released under the Apache 2.0 [License](LICENSE).

## Documentation
### Build
This will build the plugin binary in your $GOPATH/bin
```
$ go install github.com/Staples-Inc/snap-plugin-collector-couchbase 
```

### Run
Make sure your set up a snap config file and a readonly user for your couchbase server.
```json
{
  "control": {
    "plugins": {
      "collector": {
        "couchbase": {
          "all": {
            "api_url": "http://YOURSERVERNAME:8091/pools/default/buckets/",
            "username": "myRO_username",
            "password": "myRO_password"
          }
        }
      }
    }
  }
}
```

Run snapd with the config file. Then use snapctl to list discovered metrics. You can then write your task file. We have an examples in the examples folder.
```
$ ./snapd --plugin-trust 0 --log-level 1 --config /path/to/config/config.json
$ ./snapctl plugin load $GOPATH/bin/snap-plugin-collector-couchbase
$ ./snapctl metric list
```

You can then write your task file. We have an examples here below that also uses the passthru and mock-file plugin.
```yaml
---
  version: 1
  schedule: 
    type: "simple"
    interval: "5s"
  max-failures: 10
  workflow: 
    collect: 
      metrics: 
        /staples/couchbase/*: {}
      config: 
      process: 
        - 
          plugin_name: "passthru"
          process: null
          publish: 
            - 
              plugin_name: "mock-file"
              config: 
                file: "./snap_published_mongo_file.log"
```

### Source structure
```
main.go
couchbase/
  |  collector.go //implements the collection of bucket stats
  |  couchbase.go //implements the snap interfaces.
```

### TODO
- Build Test
- Build With Glide
