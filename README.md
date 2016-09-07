# Couchbase Collector Plugin
Intelsdi-x snap plugin for couchbase bucket monitoring. It uses couchbase's provided RESTAPI. Read more about bucket monitoring in this [blog post](http://blog.couchbase.com/monitoring-couchbase-cluster) from couchbase.

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
Snap is Open Source software released under the Apache 2.0 License.

## Documentation
### Build
```
$ go install // this will build the binary in your $GOPATH/bin 
```

### Source structure
```
main.go
couchbase/
  |  collector.go
  |  couchbase.go
```

### TODO
- Build Test
- Build With Glide
- 
