# Couchbase Collector Plugin
Intelsdi-x snap plugin for couchbase node/bucket monitoring. It uses couchbase's provided RESTAPI.  

## TODO
- [ ] Implement config options for time intervals for /stats
- [ ] Test with muti node

## Source structure
```
main.go
couchbase/
  |  collector.go
  |  couchbase.go
```

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

Namespace | Data Type | Description
----------|-----------|----------------------
/staples/couchbase/{name}/avg_bg_wait_time | int | exssss 
/staples/couchbase/{name}/avg_disk_commit_time | int | sjsjs
/staples/couchbase/{name}/avg_disk_update_time | int | hshshsh
