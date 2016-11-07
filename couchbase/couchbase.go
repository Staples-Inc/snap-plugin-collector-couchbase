package couchbase

import (
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"

	"fmt"
	"strings"
	"time"
)

const (
	Name = "couchbase"
	Version = 1
)

type CouchBasePlugin struct {
	initialized bool
	callDiscovery []string
	app collector
}

// New returns initialized instance of Plugin collector
func New() *CouchBasePlugin {
	self := new(CouchBasePlugin)
	return self
}

// GetConfigPolicy returns plugin config policy
func (p *CouchBasePlugin) GetConfigPolicy() (plugin.ConfigPolicy, error) {
	cfg := plugin.NewConfigPolicy()
	return *cfg, nil
}

// Collect requested metrics. Error is returned when metric collection failed or plugin
// initialization was unsuccessful.
func (p *CouchBasePlugin) CollectMetrics(mts []plugin.Metric) ([]plugin.Metric, error) {
	// check init
	if len(mts) > 0 {
		err := p.init(mts[0].Config)

		if err != nil {
			return nil, err
		}
	} else {
		return mts, nil
	}

	// results to be returned
	results := make([]plugin.Metric, len(mts))

	metrics, err := p.app.Collect()

	if err != nil {
		return nil, err
	}

	t := time.Now()

	for i, mt := range mts {
		value := metrics[parseName(mt.Namespace.Strings())]
		results[i] = plugin.Metric {
			Namespace: mt.Namespace,
			Data:      value,
			Timestamp: t,
		}
	}

	return results, nil
}

// GetMetricTypes returns list of available metrics. If initialization failed
// error is returned.
func (p *CouchBasePlugin) GetMetricTypes(cfg plugin.Config) ([]plugin.Metric, error) {
	//check init
	err := p.init(cfg)
	if err != nil {
		return nil, err
	}

	mts := []plugin.Metric{}

	for _, k := range p.callDiscovery {
		mts = append(mts, plugin.Metric{Namespace: plugin.NewNamespace(makeName(k)...)})
	}
	return mts, nil
}

// init performs one time initialization of plugin. Reads configuration from cfg
// and constructs all service objets that will be used during plugin's lifetime.
// returns error if initialization failed.
func (p *CouchBasePlugin) init(cfg plugin.Config) error {
	if p.initialized {
		return nil
	}

	api, err := cfg.GetString("api_url")
	un, err := cfg.GetString("username")
	pw, err := cfg.GetString("password")
	if err != nil {
		return fmt.Errorf("plugin initalization failed : [%v]", err)
	}

	cfgItems := map[string]string{
		"api_url": api,
		"username": un,
		"password": pw,
	}

	p.app = makeCollector(cfgItems)

	metrics, err := p.app.Discover()
	if err != nil {
		return err
	}

	for _, m := range metrics {
		p.callDiscovery = append(p.callDiscovery, m)
	}

	p.initialized = true
	return nil
}

// for mocking
var makeCollector = func(cfg map[string]string) collector { return NewCollector(cfg) }

// prefix of all namespaces
var namespacePrefix = []string{"staples", "couchbase"}

type collector interface {
	Discover() ([]string, error)
	Collect() (map[string]interface{}, error)
}

// makeName makes namespace from metric path (with segments separated by '/' ) and
// namespace prefix
func makeName(m string) []string {
	res := []string{}
	res = append(res, namespacePrefix...)
	res = append(res, strings.Split(m, "/")...)

	return res
}

// parseName extracts metric path from namespace by trimming prefix and concatenating
// remaining segments with '/'.
func parseName(ns []string) string {
	return strings.Join(ns[len(namespacePrefix):], "/")
}
