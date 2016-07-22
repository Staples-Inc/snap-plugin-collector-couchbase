package couchbase

import (
	"github.com/intelsdi-x/snap-plugin-utilities/config"
	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/control/plugin/cpolicy"
	"github.com/intelsdi-x/snap/core"

	"fmt"
	"strings"
	"time"
)

const (
	// Name of plugin
	Name = "couchbase"
	// Version of plugin
	Version = 1
	// Type of plugin
	Type = plugin.CollectorPluginType
)

type CouchBasePlugin struct {
	// is initialized
	initialized bool
	// mapping of metric namespaces and type
	callDiscovery map[string]int
	// an interface for the metrics collector
	couchbase collector
}

// Meta returns plugin's metadata
func Meta() *plugin.PluginMeta {
	return plugin.NewPluginMeta(Name, Version, Type, []string{plugin.SnapGOBContentType}, []string{plugin.SnapGOBContentType})
}

// New returns initialized instance of NetStat Plugin collector
func New() *CouchBasePlugin {
	self := new(CouchBasePlugin)
	self.callDiscovery = map[string]int{}

	//	self.Test()
	return self
}

func (p *CouchBasePlugin) Test() {
	p.init("")
	calls := map[int]bool{}
	datass, _ := p.couchbase.Collect(calls)
	fmt.Println(datass["couch_total_disk_size"].([]interface{}))
}

// GetConfigPolicy returns plugin config policy
func (p *CouchBasePlugin) GetConfigPolicy() (*cpolicy.ConfigPolicy, error) {
	c := cpolicy.New()
	return c, nil
}

// CollectMetrics finds required request ids required to collect given metrics,
// asks collector service for metrics associated with these calls and returns
// requested metrics. Error is returned when metric collection failed or plugin
// initialization was unsuccessful.
func (p *CouchBasePlugin) CollectMetrics(mts []plugin.MetricType) ([]plugin.MetricType, error) {

	// check init
	if len(mts) > 0 {
		err := p.init(mts[0])

		if err != nil {
			return nil, err
		}
	} else {
		return mts, nil
	}

	// results to be returned
	results := make([]plugin.MetricType, len(mts))

	// TODO: Use this to make unique collections. Does nothing right now.
	calls := map[int]bool{}

	metrics, err := p.couchbase.Collect(calls)

	if err != nil {
		return nil, err
	}

	t := time.Now()

	for i, mt := range mts {
		value := metrics[parseName(mt.Namespace().Strings())].([]interface{})[0]
		results[i] = plugin.MetricType{
			Namespace_: mt.Namespace(),
			Data_:      value,
			Timestamp_: t,
		}
	}

	return results, nil
}

// GetMetricTypes returns list of available metrics. If initialization failed
// error is returned.
func (p *CouchBasePlugin) GetMetricTypes(cfg plugin.ConfigType) ([]plugin.MetricType, error) {

	//check init
	err := p.init(cfg)
	if err != nil {
		return nil, err
	}

	mts := []plugin.MetricType{}

	for k := range p.callDiscovery {
		mts = append(mts, plugin.MetricType{Namespace_: core.NewNamespace(makeName(k)...)})
	}
	return mts, nil
}

// init performs one time initialization of plugin. Reads configuration from cfg
// and constructs all service objets that will be used during plugin's lifetime.
// returns error if initialization failed.
func (p *CouchBasePlugin) init(cfg interface{}) error {

	if p.initialized {
		return nil
	}

	cfgItems, err := config.GetConfigItems(cfg, "api_url", "username", "password")
	if err != nil {
		return fmt.Errorf("plugin initalization failed : [%v]", err)
	}

	p.couchbase = makeCollector(cfgItems)

	metrics, err := p.couchbase.Discover()
	if err != nil {
		return err
	}

	for _, m := range metrics {
		// TODO: GROUP TYPE OF METRICS
		p.callDiscovery[m.Name] = 0
	}

	p.initialized = true
	return nil
}

// for mocking
var makeCollector = func(cfg map[string]interface{}) collector { return NewCollector(cfg) }

// prefix of all namespaces
var namespacePrefix = []string{"intel", "couchbase"}

type collector interface {
	Discover() ([]metric, error)
	Collect(metrics map[int]bool) (map[string]interface{}, error)
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
