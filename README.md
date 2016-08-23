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
/staples/couchbase/{name}/avg_bg_wait_time 
/staples/couchbase/{name}/avg_disk_commit_time 
/staples/couchbase/{name}/avg_disk_update_time 
/staples/couchbase/{name}/bg_wait_count 
/staples/couchbase/{name}/bg_wait_total 
/staples/couchbase/{name}/bytes_read 
/staples/couchbase/{name}/bytes_written 
/staples/couchbase/{name}/cas_badval 
/staples/couchbase/{name}/cas_hits 
/staples/couchbase/{name}/cas_misses 
/staples/couchbase/{name}/cmd_get 
/staples/couchbase/{name}/cmd_set 
/staples/couchbase/{name}/couch_docs_actual_disk_size 
/staples/couchbase/{name}/couch_docs_data_size 
/staples/couchbase/{name}/couch_docs_disk_size 
/staples/couchbase/{name}/couch_docs_fragmentation 
/staples/couchbase/{name}/couch_spatial_data_size 
/staples/couchbase/{name}/couch_spatial_disk_size 
/staples/couchbase/{name}/couch_spatial_ops 
/staples/couchbase/{name}/couch_total_disk_size 
/staples/couchbase/{name}/couch_views_actual_disk_size 
/staples/couchbase/{name}/couch_views_data_size 
/staples/couchbase/{name}/couch_views_disk_size 
/staples/couchbase/{name}/couch_views_fragmentation 
/staples/couchbase/{name}/couch_views_ops 
/staples/couchbase/{name}/cpu_idle_ms 
/staples/couchbase/{name}/cpu_local_ms 
/staples/couchbase/{name}/cpu_utilization_rate 
/staples/couchbase/{name}/curr_connections 
/staples/couchbase/{name}/curr_items 
/staples/couchbase/{name}/curr_items_tot 
/staples/couchbase/{name}/decr_hits 
/staples/couchbase/{name}/decr_misses 
/staples/couchbase/{name}/delete_hits 
/staples/couchbase/{name}/delete_misses 
/staples/couchbase/{name}/disk_commit_count 
/staples/couchbase/{name}/disk_commit_total 
/staples/couchbase/{name}/disk_update_count 
/staples/couchbase/{name}/disk_update_total 
/staples/couchbase/{name}/disk_write_queue 
/staples/couchbase/{name}/ep_bg_fetched 
/staples/couchbase/{name}/ep_cache_miss_rate 
/staples/couchbase/{name}/ep_dcp_2i_backoff 
/staples/couchbase/{name}/ep_dcp_2i_count 
/staples/couchbase/{name}/ep_dcp_2i_items_remaining 
/staples/couchbase/{name}/ep_dcp_2i_items_sent 
/staples/couchbase/{name}/ep_dcp_2i_producer_count 
/staples/couchbase/{name}/ep_dcp_2i_total_backlog_size 
/staples/couchbase/{name}/ep_dcp_2i_total_bytes 
/staples/couchbase/{name}/ep_dcp_other_backoff 
/staples/couchbase/{name}/ep_dcp_other_count 
/staples/couchbase/{name}/ep_dcp_other_items_remaining 
/staples/couchbase/{name}/ep_dcp_other_items_sent 
/staples/couchbase/{name}/ep_dcp_other_producer_count 
/staples/couchbase/{name}/ep_dcp_other_total_backlog_size 
/staples/couchbase/{name}/ep_dcp_other_total_bytes 
/staples/couchbase/{name}/ep_dcp_replica_backoff 
/staples/couchbase/{name}/ep_dcp_replica_count 
/staples/couchbase/{name}/ep_dcp_replica_items_remaining 
/staples/couchbase/{name}/ep_dcp_replica_items_sent 
/staples/couchbase/{name}/ep_dcp_replica_producer_count 
/staples/couchbase/{name}/ep_dcp_replica_total_backlog_size 
/staples/couchbase/{name}/ep_dcp_replica_total_bytes 
/staples/couchbase/{name}/ep_dcp_views+2i_backoff 
/staples/couchbase/{name}/ep_dcp_views+2i_count 
/staples/couchbase/{name}/ep_dcp_views+2i_items_remaining 
/staples/couchbase/{name}/ep_dcp_views+2i_items_sent 
/staples/couchbase/{name}/ep_dcp_views+2i_producer_count 
/staples/couchbase/{name}/ep_dcp_views+2i_total_backlog_size 
/staples/couchbase/{name}/ep_dcp_views+2i_total_bytes 
/staples/couchbase/{name}/ep_dcp_views_backoff 
/staples/couchbase/{name}/ep_dcp_views_count 
/staples/couchbase/{name}/ep_dcp_views_items_remaining 
/staples/couchbase/{name}/ep_dcp_views_items_sent 
/staples/couchbase/{name}/ep_dcp_views_producer_count 
/staples/couchbase/{name}/ep_dcp_views_total_backlog_size 
/staples/couchbase/{name}/ep_dcp_views_total_bytes 
/staples/couchbase/{name}/ep_dcp_xdcr_backoff 
/staples/couchbase/{name}/ep_dcp_xdcr_count 
/staples/couchbase/{name}/ep_dcp_xdcr_items_remaining 
/staples/couchbase/{name}/ep_dcp_xdcr_items_sent 
/staples/couchbase/{name}/ep_dcp_xdcr_producer_count 
/staples/couchbase/{name}/ep_dcp_xdcr_total_backlog_size 
/staples/couchbase/{name}/ep_dcp_xdcr_total_bytes 
/staples/couchbase/{name}/ep_diskqueue_drain 
/staples/couchbase/{name}/ep_diskqueue_fill 
/staples/couchbase/{name}/ep_diskqueue_items 
/staples/couchbase/{name}/ep_flusher_todo 
/staples/couchbase/{name}/ep_item_commit_failed 
/staples/couchbase/{name}/ep_kv_size 
/staples/couchbase/{name}/ep_max_size 
/staples/couchbase/{name}/ep_mem_high_wat 
/staples/couchbase/{name}/ep_mem_low_wat 
/staples/couchbase/{name}/ep_meta_data_memory 
/staples/couchbase/{name}/ep_num_non_resident 
/staples/couchbase/{name}/ep_num_ops_del_meta 
/staples/couchbase/{name}/ep_num_ops_del_ret_meta 
/staples/couchbase/{name}/ep_num_ops_get_meta 
/staples/couchbase/{name}/ep_num_ops_set_meta 
/staples/couchbase/{name}/ep_num_ops_set_ret_meta 
/staples/couchbase/{name}/ep_num_value_ejects 
/staples/couchbase/{name}/ep_oom_errors 
/staples/couchbase/{name}/ep_ops_create 
/staples/couchbase/{name}/ep_ops_update 
/staples/couchbase/{name}/ep_overhead 
/staples/couchbase/{name}/ep_queue_size 
/staples/couchbase/{name}/ep_resident_items_rate 
/staples/couchbase/{name}/ep_tap_rebalance_count 
/staples/couchbase/{name}/ep_tap_rebalance_qlen 
/staples/couchbase/{name}/ep_tap_rebalance_queue_backfillremaining 
/staples/couchbase/{name}/ep_tap_rebalance_queue_backoff 
/staples/couchbase/{name}/ep_tap_rebalance_queue_drain 
/staples/couchbase/{name}/ep_tap_rebalance_queue_fill 
/staples/couchbase/{name}/ep_tap_rebalance_queue_itemondisk 
/staples/couchbase/{name}/ep_tap_rebalance_total_backlog_size 
/staples/couchbase/{name}/ep_tap_replica_count 
/staples/couchbase/{name}/ep_tap_replica_qlen 
/staples/couchbase/{name}/ep_tap_replica_queue_backfillremaining 
/staples/couchbase/{name}/ep_tap_replica_queue_backoff 
/staples/couchbase/{name}/ep_tap_replica_queue_drain 
/staples/couchbase/{name}/ep_tap_replica_queue_fill 
/staples/couchbase/{name}/ep_tap_replica_queue_itemondisk 
/staples/couchbase/{name}/ep_tap_replica_total_backlog_size 
/staples/couchbase/{name}/ep_tap_total_count 
/staples/couchbase/{name}/ep_tap_total_qlen 
/staples/couchbase/{name}/ep_tap_total_queue_backfillremaining 
/staples/couchbase/{name}/ep_tap_total_queue_backoff 
/staples/couchbase/{name}/ep_tap_total_queue_drain 
/staples/couchbase/{name}/ep_tap_total_queue_fill 
/staples/couchbase/{name}/ep_tap_total_queue_itemondisk 
/staples/couchbase/{name}/ep_tap_total_total_backlog_size 
/staples/couchbase/{name}/ep_tap_user_count 
/staples/couchbase/{name}/ep_tap_user_qlen 
/staples/couchbase/{name}/ep_tap_user_queue_backfillremaining 
/staples/couchbase/{name}/ep_tap_user_queue_backoff 
/staples/couchbase/{name}/ep_tap_user_queue_drain 
/staples/couchbase/{name}/ep_tap_user_queue_fill 
/staples/couchbase/{name}/ep_tap_user_queue_itemondisk 
/staples/couchbase/{name}/ep_tap_user_total_backlog_size 
/staples/couchbase/{name}/ep_tmp_oom_errors 
/staples/couchbase/{name}/ep_vb_total 
/staples/couchbase/{name}/evictions 
/staples/couchbase/{name}/get_hits 
/staples/couchbase/{name}/get_misses 
/staples/couchbase/{name}/hibernated_requests 
/staples/couchbase/{name}/hibernated_waked 
/staples/couchbase/{name}/hit_ratio 
/staples/couchbase/{name}/incr_hits 
/staples/couchbase/{name}/incr_misses 
/staples/couchbase/{name}/mem_actual_free 
/staples/couchbase/{name}/mem_actual_used 
/staples/couchbase/{name}/mem_free 
/staples/couchbase/{name}/mem_total 
/staples/couchbase/{name}/mem_used 
/staples/couchbase/{name}/mem_used_sys 
/staples/couchbase/{name}/misses 
/staples/couchbase/{name}/ops 
/staples/couchbase/{name}/rest_requests 
/staples/couchbase/{name}/spatial/ce5bd2776798930ea5093a18742d4bbf/accesses 
/staples/couchbase/{name}/spatial/ce5bd2776798930ea5093a18742d4bbf/data_size 
/staples/couchbase/{name}/spatial/ce5bd2776798930ea5093a18742d4bbf/disk_size 
/staples/couchbase/{name}/swap_total 
/staples/couchbase/{name}/swap_used 
/staples/couchbase/{name}/timestamp 
/staples/couchbase/{name}/vb_active_eject 
/staples/couchbase/{name}/vb_active_itm_memory 
/staples/couchbase/{name}/vb_active_meta_data_memory 
/staples/couchbase/{name}/vb_active_num 
/staples/couchbase/{name}/vb_active_num_non_resident 
/staples/couchbase/{name}/vb_active_ops_create 
/staples/couchbase/{name}/vb_active_ops_update 
/staples/couchbase/{name}/vb_active_queue_age 
/staples/couchbase/{name}/vb_active_queue_drain 
/staples/couchbase/{name}/vb_active_queue_fill 
/staples/couchbase/{name}/vb_active_queue_size 
/staples/couchbase/{name}/vb_active_resident_items_ratio 
/staples/couchbase/{name}/vb_avg_active_queue_age 
/staples/couchbase/{name}/vb_avg_pending_queue_age 
/staples/couchbase/{name}/vb_avg_replica_queue_age 
/staples/couchbase/{name}/vb_avg_total_queue_age 
/staples/couchbase/{name}/vb_pending_curr_items 
/staples/couchbase/{name}/vb_pending_eject 
/staples/couchbase/{name}/vb_pending_itm_memory 
/staples/couchbase/{name}/vb_pending_meta_data_memory 
/staples/couchbase/{name}/vb_pending_num 
/staples/couchbase/{name}/vb_pending_num_non_resident 
/staples/couchbase/{name}/vb_pending_ops_create 
/staples/couchbase/{name}/vb_pending_ops_update 
/staples/couchbase/{name}/vb_pending_queue_age 
/staples/couchbase/{name}/vb_pending_queue_drain 
/staples/couchbase/{name}/vb_pending_queue_fill 
/staples/couchbase/{name}/vb_pending_queue_size 
/staples/couchbase/{name}/vb_pending_resident_items_ratio 
/staples/couchbase/{name}/vb_replica_curr_items 
/staples/couchbase/{name}/vb_replica_eject 
/staples/couchbase/{name}/vb_replica_itm_memory 
/staples/couchbase/{name}/vb_replica_meta_data_memory 
/staples/couchbase/{name}/vb_replica_num 
/staples/couchbase/{name}/vb_replica_num_non_resident 
/staples/couchbase/{name}/vb_replica_ops_create 
/staples/couchbase/{name}/vb_replica_ops_update 
/staples/couchbase/{name}/vb_replica_queue_age 
/staples/couchbase/{name}/vb_replica_queue_drain 
/staples/couchbase/{name}/vb_replica_queue_fill 
/staples/couchbase/{name}/vb_replica_queue_size 
/staples/couchbase/{name}/vb_replica_resident_items_ratio 
/staples/couchbase/{name}/vb_total_queue_age 
/staples/couchbase/{name}/views/a52eab6056871e78f2296dc6d1f4aaaf/accesses 
/staples/couchbase/{name}/views/a52eab6056871e78f2296dc6d1f4aaaf/data_size 
/staples/couchbase/{name}/views/a52eab6056871e78f2296dc6d1f4aaaf/disk_size 
/staples/couchbase/{name}/xdc_ops 
