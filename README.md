# Snap Couchbase Collector Plugin
Intelsdi-x [snap](http://github.com/intelsdi-x/snap) plugin for couchbase bucket monitoring. It uses couchbase's provided RESTAPI. Read more about bucket monitoring in this [blog post](http://blog.couchbase.com/monitoring-couchbase-cluster) from couchbase.

### Snap version requires at least for plugin_v
- v0.14.0-beta - v1.0
- v1.0.0 - v2.0

### Snap version tested up to plugin_v
- v0.15.0-beta - v1.0
- v1.0.0 - v2.0

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
$ go get github.com/Staples-Inc/snap-plugin-collector-couchbase
```
We like to use glide vendor install though.
```
$ cd github.com/Staples-Inc/snap-plugin-collector-couchbase
$ glide up    # Glide will create vendor folder.
$ go build    # Then build.
```

### Run
Run snapteld with a config file if running v1.0. Then use snaptel to load plugin and list metrics names.
```
$ ./snapteld --plugin-trust 0 --log-level 1 --config /path/to/config/config.json
$ ./snaptel plugin load $GOPATH/bin/snap-plugin-collector-couchbase
$ ./snaptel metric list
```

You can then write your task file. We have an example here below that also uses the passthru and mock-file plugin.
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
        /staples/couchbase:
          api_url: "http://MYSERVER:8091/pools/default/buckets/"
          username: "ROUSER"
          password: "ROPW"
      process:
        -
          plugin_name: "test-reverse-processor"
          process: null
          publish:
            -
              plugin_name: "test-file-publisher"
              config:
                file: "/tmp/snap_published_grpc_file.log"
```

### Source structure
```
main.go
couchbase/
  |  collector.go //implements the collection of bucket stats
  |  couchbase.go //implements the snap interfaces.
  |  metrics.go   //contains a string slice of metrics
```

### TODO
- Build Test
