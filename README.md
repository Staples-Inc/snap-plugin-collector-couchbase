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
/staples/couchbase/{name}/avg_bg_wait_time | float | somthing
/staples/couchbase/{name}/avg_disk_commit_time | float | somthing
/staples/couchbase/{name}/avg_disk_update_time | float | somthing
/staples/couchbase/{name}/bg_wait_count | float | somthing
/staples/couchbase/{name}/bg_wait_total | float | somthing
/staples/couchbase/{name}/bytes_read | float | somthing
/staples/couchbase/{name}/bytes_written | float | somthing
/staples/couchbase/{name}/cas_badval | float | somthing
/staples/couchbase/{name}/cas_hits | float | somthing
/staples/couchbase/{name}/cas_misses | float | somthing
/staples/couchbase/{name}/cmd_get | float | somthing
/staples/couchbase/{name}/cmd_set | float | somthing
/staples/couchbase/{name}/couch_docs_actual_disk_size | float | somthing
/staples/couchbase/{name}/couch_docs_data_size | float | somthing
/staples/couchbase/{name}/couch_docs_disk_size | float | somthing
/staples/couchbase/{name}/couch_docs_fragmentation | float | somthing
/staples/couchbase/{name}/couch_spatial_data_size | float | somthing
/staples/couchbase/{name}/couch_spatial_disk_size | float | somthing
/staples/couchbase/{name}/couch_spatial_ops | float | somthing
/staples/couchbase/{name}/couch_total_disk_size | float | somthing
/staples/couchbase/{name}/couch_views_actual_disk_size | float | somthing
/staples/couchbase/{name}/couch_views_data_size | float | somthing
/staples/couchbase/{name}/couch_views_disk_size | float | somthing
/staples/couchbase/{name}/couch_views_fragmentation | float | somthing
/staples/couchbase/{name}/couch_views_ops | float | somthing
/staples/couchbase/{name}/cpu_idle_ms | float | somthing
/staples/couchbase/{name}/cpu_local_ms | float | somthing
/staples/couchbase/{name}/cpu_utilization_rate | float | somthing
/staples/couchbase/{name}/curr_connections | float | somthing
/staples/couchbase/{name}/curr_items | float | somthing
/staples/couchbase/{name}/curr_items_tot | float | somthing
/staples/couchbase/{name}/decr_hits | float | somthing
/staples/couchbase/{name}/decr_misses | float | somthing
/staples/couchbase/{name}/delete_hits | float | somthing
/staples/couchbase/{name}/delete_misses | float | somthing
/staples/couchbase/{name}/disk_commit_count | float | somthing
/staples/couchbase/{name}/disk_commit_total | float | somthing
/staples/couchbase/{name}/disk_update_count | float | somthing
/staples/couchbase/{name}/disk_update_total | float | somthing
/staples/couchbase/{name}/disk_write_queue | float | somthing
/staples/couchbase/{name}/ep_bg_fetched | float | somthing
/staples/couchbase/{name}/ep_cache_miss_rate | float | somthing
/staples/couchbase/{name}/ep_dcp_2i_backoff | float | somthing
/staples/couchbase/{name}/ep_dcp_2i_count | float | somthing
/staples/couchbase/{name}/ep_dcp_2i_items_remaining | float | somthing
/staples/couchbase/{name}/ep_dcp_2i_items_sent | float | somthing
/staples/couchbase/{name}/ep_dcp_2i_producer_count | float | somthing
/staples/couchbase/{name}/ep_dcp_2i_total_backlog_size | float | somthing
/staples/couchbase/{name}/ep_dcp_2i_total_bytes | float | somthing
/staples/couchbase/{name}/ep_dcp_other_backoff | float | somthing
/staples/couchbase/{name}/ep_dcp_other_count | float | somthing
/staples/couchbase/{name}/ep_dcp_other_items_remaining | float | somthing
/staples/couchbase/{name}/ep_dcp_other_items_sent | float | somthing
/staples/couchbase/{name}/ep_dcp_other_producer_count | float | somthing
/staples/couchbase/{name}/ep_dcp_other_total_backlog_size | float | somthing
/staples/couchbase/{name}/ep_dcp_other_total_bytes | float | somthing
/staples/couchbase/{name}/ep_dcp_replica_backoff | float | somthing
/staples/couchbase/{name}/ep_dcp_replica_count | float | somthing
/staples/couchbase/{name}/ep_dcp_replica_items_remaining | float | somthing
/staples/couchbase/{name}/ep_dcp_replica_items_sent | float | somthing
/staples/couchbase/{name}/ep_dcp_replica_producer_count | float | somthing
/staples/couchbase/{name}/ep_dcp_replica_total_backlog_size | float | somthing
/staples/couchbase/{name}/ep_dcp_replica_total_bytes | float | somthing
/staples/couchbase/{name}/ep_dcp_views+2i_backoff | float | somthing
/staples/couchbase/{name}/ep_dcp_views+2i_count | float | somthing
/staples/couchbase/{name}/ep_dcp_views+2i_items_remaining | float | somthing
/staples/couchbase/{name}/ep_dcp_views+2i_items_sent | float | somthing
/staples/couchbase/{name}/ep_dcp_views+2i_producer_count | float | somthing
/staples/couchbase/{name}/ep_dcp_views+2i_total_backlog_size | float | somthing
/staples/couchbase/{name}/ep_dcp_views+2i_total_bytes | float | somthing
/staples/couchbase/{name}/ep_dcp_views_backoff | float | somthing
/staples/couchbase/{name}/ep_dcp_views_count | float | somthing
/staples/couchbase/{name}/ep_dcp_views_items_remaining | float | somthing
/staples/couchbase/{name}/ep_dcp_views_items_sent | float | somthing
/staples/couchbase/{name}/ep_dcp_views_producer_count | float | somthing
/staples/couchbase/{name}/ep_dcp_views_total_backlog_size | float | somthing
/staples/couchbase/{name}/ep_dcp_views_total_bytes | float | somthing
/staples/couchbase/{name}/ep_dcp_xdcr_backoff | float | somthing
/staples/couchbase/{name}/ep_dcp_xdcr_count | float | somthing
/staples/couchbase/{name}/ep_dcp_xdcr_items_remaining | float | somthing
/staples/couchbase/{name}/ep_dcp_xdcr_items_sent | float | somthing
/staples/couchbase/{name}/ep_dcp_xdcr_producer_count | float | somthing
/staples/couchbase/{name}/ep_dcp_xdcr_total_backlog_size | float | somthing
/staples/couchbase/{name}/ep_dcp_xdcr_total_bytes | float | somthing
/staples/couchbase/{name}/ep_diskqueue_drain | float | somthing
/staples/couchbase/{name}/ep_diskqueue_fill | float | somthing
/staples/couchbase/{name}/ep_diskqueue_items | float | somthing
/staples/couchbase/{name}/ep_flusher_todo | float | somthing
/staples/couchbase/{name}/ep_item_commit_failed | float | somthing
/staples/couchbase/{name}/ep_kv_size | float | somthing
/staples/couchbase/{name}/ep_max_size | float | somthing
/staples/couchbase/{name}/ep_mem_high_wat | float | somthing
/staples/couchbase/{name}/ep_mem_low_wat | float | somthing
/staples/couchbase/{name}/ep_meta_data_memory | float | somthing
/staples/couchbase/{name}/ep_num_non_resident | float | somthing
/staples/couchbase/{name}/ep_num_ops_del_meta | float | somthing
/staples/couchbase/{name}/ep_num_ops_del_ret_meta | float | somthing
/staples/couchbase/{name}/ep_num_ops_get_meta | float | somthing
/staples/couchbase/{name}/ep_num_ops_set_meta | float | somthing
/staples/couchbase/{name}/ep_num_ops_set_ret_meta | float | somthing
/staples/couchbase/{name}/ep_num_value_ejects | float | somthing
/staples/couchbase/{name}/ep_oom_errors | float | somthing
/staples/couchbase/{name}/ep_ops_create | float | somthing
/staples/couchbase/{name}/ep_ops_update | float | somthing
/staples/couchbase/{name}/ep_overhead | float | somthing
/staples/couchbase/{name}/ep_queue_size | float | somthing
/staples/couchbase/{name}/ep_resident_items_rate | float | somthing
/staples/couchbase/{name}/ep_tap_rebalance_count | float | somthing
/staples/couchbase/{name}/ep_tap_rebalance_qlen | float | somthing
/staples/couchbase/{name}/ep_tap_rebalance_queue_backfillremaining | float | somthing
/staples/couchbase/{name}/ep_tap_rebalance_queue_backoff | float | somthing
/staples/couchbase/{name}/ep_tap_rebalance_queue_drain | float | somthing
/staples/couchbase/{name}/ep_tap_rebalance_queue_fill | float | somthing
/staples/couchbase/{name}/ep_tap_rebalance_queue_itemondisk | float | somthing
/staples/couchbase/{name}/ep_tap_rebalance_total_backlog_size | float | somthing
/staples/couchbase/{name}/ep_tap_replica_count | float | somthing
/staples/couchbase/{name}/ep_tap_replica_qlen | float | somthing
/staples/couchbase/{name}/ep_tap_replica_queue_backfillremaining | float | somthing
/staples/couchbase/{name}/ep_tap_replica_queue_backoff | float | somthing
/staples/couchbase/{name}/ep_tap_replica_queue_drain | float | somthing
/staples/couchbase/{name}/ep_tap_replica_queue_fill | float | somthing
/staples/couchbase/{name}/ep_tap_replica_queue_itemondisk | float | somthing
/staples/couchbase/{name}/ep_tap_replica_total_backlog_size | float | somthing
/staples/couchbase/{name}/ep_tap_total_count | float | somthing
/staples/couchbase/{name}/ep_tap_total_qlen | float | somthing
/staples/couchbase/{name}/ep_tap_total_queue_backfillremaining | float | somthing
/staples/couchbase/{name}/ep_tap_total_queue_backoff | float | somthing
/staples/couchbase/{name}/ep_tap_total_queue_drain | float | somthing
/staples/couchbase/{name}/ep_tap_total_queue_fill | float | somthing
/staples/couchbase/{name}/ep_tap_total_queue_itemondisk | float | somthing
/staples/couchbase/{name}/ep_tap_total_total_backlog_size | float | somthing
/staples/couchbase/{name}/ep_tap_user_count | float | somthing
/staples/couchbase/{name}/ep_tap_user_qlen | float | somthing
/staples/couchbase/{name}/ep_tap_user_queue_backfillremaining | float | somthing
/staples/couchbase/{name}/ep_tap_user_queue_backoff | float | somthing
/staples/couchbase/{name}/ep_tap_user_queue_drain | float | somthing
/staples/couchbase/{name}/ep_tap_user_queue_fill | float | somthing
/staples/couchbase/{name}/ep_tap_user_queue_itemondisk | float | somthing
/staples/couchbase/{name}/ep_tap_user_total_backlog_size | float | somthing
/staples/couchbase/{name}/ep_tmp_oom_errors | float | somthing
/staples/couchbase/{name}/ep_vb_total | float | somthing
/staples/couchbase/{name}/evictions | float | somthing
/staples/couchbase/{name}/get_hits | float | somthing
/staples/couchbase/{name}/get_misses | float | somthing
/staples/couchbase/{name}/hibernated_requests | float | somthing
/staples/couchbase/{name}/hibernated_waked | float | somthing
/staples/couchbase/{name}/hit_ratio | float | somthing
/staples/couchbase/{name}/incr_hits | float | somthing
/staples/couchbase/{name}/incr_misses | float | somthing
/staples/couchbase/{name}/mem_actual_free | float | somthing
/staples/couchbase/{name}/mem_actual_used | float | somthing
/staples/couchbase/{name}/mem_free | float | somthing
/staples/couchbase/{name}/mem_total | float | somthing
/staples/couchbase/{name}/mem_used | float | somthing
/staples/couchbase/{name}/mem_used_sys | float | somthing
/staples/couchbase/{name}/misses | float | somthing
/staples/couchbase/{name}/ops | float | somthing
/staples/couchbase/{name}/rest_requests | float | somthing
/staples/couchbase/{name}/spatial/ce5bd2776798930ea5093a18742d4bbf/accesses | float | somthing
/staples/couchbase/{name}/spatial/ce5bd2776798930ea5093a18742d4bbf/data_size | float | somthing
/staples/couchbase/{name}/spatial/ce5bd2776798930ea5093a18742d4bbf/disk_size | float | somthing
/staples/couchbase/{name}/swap_total | float | somthing
/staples/couchbase/{name}/swap_used | float | somthing
/staples/couchbase/{name}/timestamp | float | somthing
/staples/couchbase/{name}/vb_active_eject | float | somthing
/staples/couchbase/{name}/vb_active_itm_memory | float | somthing
/staples/couchbase/{name}/vb_active_meta_data_memory | float | somthing
/staples/couchbase/{name}/vb_active_num | float | somthing
/staples/couchbase/{name}/vb_active_num_non_resident | float | somthing
/staples/couchbase/{name}/vb_active_ops_create | float | somthing
/staples/couchbase/{name}/vb_active_ops_update | float | somthing
/staples/couchbase/{name}/vb_active_queue_age | float | somthing
/staples/couchbase/{name}/vb_active_queue_drain | float | somthing
/staples/couchbase/{name}/vb_active_queue_fill | float | somthing
/staples/couchbase/{name}/vb_active_queue_size | float | somthing
/staples/couchbase/{name}/vb_active_resident_items_ratio | float | somthing
/staples/couchbase/{name}/vb_avg_active_queue_age | float | somthing
/staples/couchbase/{name}/vb_avg_pending_queue_age | float | somthing
/staples/couchbase/{name}/vb_avg_replica_queue_age | float | somthing
/staples/couchbase/{name}/vb_avg_total_queue_age | float | somthing
/staples/couchbase/{name}/vb_pending_curr_items | float | somthing
/staples/couchbase/{name}/vb_pending_eject | float | somthing
/staples/couchbase/{name}/vb_pending_itm_memory | float | somthing
/staples/couchbase/{name}/vb_pending_meta_data_memory | float | somthing
/staples/couchbase/{name}/vb_pending_num | float | somthing
/staples/couchbase/{name}/vb_pending_num_non_resident | float | somthing
/staples/couchbase/{name}/vb_pending_ops_create | float | somthing
/staples/couchbase/{name}/vb_pending_ops_update | float | somthing
/staples/couchbase/{name}/vb_pending_queue_age | float | somthing
/staples/couchbase/{name}/vb_pending_queue_drain | float | somthing
/staples/couchbase/{name}/vb_pending_queue_fill | float | somthing
/staples/couchbase/{name}/vb_pending_queue_size | float | somthing
/staples/couchbase/{name}/vb_pending_resident_items_ratio | float | somthing
/staples/couchbase/{name}/vb_replica_curr_items | float | somthing
/staples/couchbase/{name}/vb_replica_eject | float | somthing
/staples/couchbase/{name}/vb_replica_itm_memory | float | somthing
/staples/couchbase/{name}/vb_replica_meta_data_memory | float | somthing
/staples/couchbase/{name}/vb_replica_num | float | somthing
/staples/couchbase/{name}/vb_replica_num_non_resident | float | somthing
/staples/couchbase/{name}/vb_replica_ops_create | float | somthing
/staples/couchbase/{name}/vb_replica_ops_update | float | somthing
/staples/couchbase/{name}/vb_replica_queue_age | float | somthing
/staples/couchbase/{name}/vb_replica_queue_drain | float | somthing
/staples/couchbase/{name}/vb_replica_queue_fill | float | somthing
/staples/couchbase/{name}/vb_replica_queue_size | float | somthing
/staples/couchbase/{name}/vb_replica_resident_items_ratio | float | somthing
/staples/couchbase/{name}/vb_total_queue_age | float | somthing
/staples/couchbase/{name}/views/a52eab6056871e78f2296dc6d1f4aaaf/accesses | float | somthing
/staples/couchbase/{name}/views/a52eab6056871e78f2296dc6d1f4aaaf/data_size | float | somthing
/staples/couchbase/{name}/views/a52eab6056871e78f2296dc6d1f4aaaf/disk_size | float | somthing
/staples/couchbase/{name}/xdc_ops | float | somthing
