# Couchbase Collector Plugin
Intelsdi-x snap plugin for couchbase node/bucket monitoring. It uses couchbase's provided RESTAPI.  

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
### Source structure
```
main.go
couchbase/
  |  collector.go
  |  couchbase.go
```

### Collected Metrics
This plugin as the ability to collect the following metrics:

#### Buckets

Namespace | Data Type | Description
----------|-----------|----------------------
/staples/couchbase/{name}/avg_bg_wait_time | int | exssss 
/staples/couchbase/{name}/avg_disk_commit_time | int | sjsjs
/staples/couchbase/{name}/avg_disk_update_time | int | hshshsh
