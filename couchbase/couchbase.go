package couchbase

import (
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	"time"
)

type CouchBasePlugin struct {
	initialized bool
	collector   *BucketCollector
}

func New() (*CouchBasePlugin, error) {
	return &CouchBasePlugin{
		initialized: false,
		collector:   &BucketCollector{},
	}, nil
}

// GetConfigPolicy returns plugin config policy
func (cbp *CouchBasePlugin) GetConfigPolicy() (plugin.ConfigPolicy, error) {
	policy := plugin.NewConfigPolicy()
	policy.AddNewStringRule([]string{"staples", "couchbase"}, "api_url", true)
	policy.AddNewStringRule([]string{"staples", "couchbase"}, "username", true)
	policy.AddNewStringRule([]string{"staples", "couchbase"}, "password", true)
	return *policy, nil
}

// Collect requested metrics. Error is returned when metric collection failed or plugin
// initialization was unsuccessful.
func (cbp *CouchBasePlugin) CollectMetrics(mts []plugin.Metric) ([]plugin.Metric, error) {

	if cbp.initialized == false {
		collector, err := NewCollector(mts[0].Config)
		if err != nil {
			return nil, err
		}
		cbp.collector = collector
		cbp.initialized = true
	}

	cmts, err := cbp.collector.GetSamples()
	if err != nil {
		return nil, err
	}

	t := time.Now()
	var results []plugin.Metric

	for _, mt := range cmts {
		m := plugin.Metric{
			Namespace: plugin.Namespace{
				plugin.NamespaceElement{Value: "staples"},
				plugin.NamespaceElement{Value: "couchbase"},
				plugin.NamespaceElement{Value: mt.Bucket, Name: "bucket_name", Description: "Bucket Name"},
				plugin.NamespaceElement{Value: mt.Metric},
			},
			Data:      mt.Data,
			Timestamp: t,
		}

		results = append(results, m)
	}

	return results, nil
}

// GetMetricTypes returns list of available metrics. If initialization failed
// error is returned.
func (cbp *CouchBasePlugin) GetMetricTypes(cfg plugin.Config) ([]plugin.Metric, error) {
	mts := []plugin.Metric{}

	for _, m := range Metrics() {
		ns := plugin.NewNamespace("staples", "couchbase").AddDynamicElement("bucket_name", "Bucket Name").AddStaticElement(m)
		mts = append(mts, plugin.Metric{Namespace: ns})
	}
	return mts, nil
}
