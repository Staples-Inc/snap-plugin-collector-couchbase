package couchbase

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// NewCollector constructs new metricCollector that will request given stats and
// fail if they are unavailable.
func NewCollector(cfg map[string]interface{}) *metricCollector {
	self := new(metricCollector)

	url := cfg["api_url"].(string)
	un := cfg["username"].(string)
	pw := cfg["password"].(string)

	// TODO: Switch for types??
	self.samples = &BucketStats{url: url, un: un, pw: pw}

	// Hard coded values could be removed.
	//self.url = "http://localhost:32777/pools/default/buckets/travel-sample/stats"
	//self.un = "admin"
	//self.pw = "password"

	return self
}

// metricCollector implements logic for discovering available metrics
type metricCollector struct {
	samples metrictap
}

// metric contains name of metric and id of call that collects particular metric.
type metric struct {
	Name string
	Call int
}

// Collect performs given set of calls (indicated by true value in metrics map).
// returns map of metric values (accessible by metric name). If any of requested
// calls fail error is returned.
func (mc *metricCollector) Collect(metrics map[int]bool) (map[string]interface{}, error) {
	samples, err := mc.samples.GetSamples()
	if err != nil {
		return nil, err
	}

	return samples, nil
}

// Discover performs metric discovery. Returns valid metric names and associated
// Call id's. If mandatory request fails error is returned.
func (mc *metricCollector) Discover() ([]metric, error) {
	samples, err := mc.samples.GetSamples()
	if err != nil {
		return nil, err
	}

	res := []metric{}

	for key := range samples {
		res = append(res, metric{Name: key, Call: 0})
	}

	return res, nil
}

// Interface for couchbase
type metrictap interface {
	GetSamples() (samples map[string]interface{}, err error)
}

// BucketStats
type BucketStats struct {
	url string
	un  string
	pw  string
}

func (self BucketStats) GetSamples() (samples map[string]interface{}, err error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", self.url, nil)
	req.SetBasicAuth(self.un, self.pw)
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var couchdata map[string]interface{}
	err = json.Unmarshal(body, &couchdata)
	if err != nil {
		return
	}

	res := couchdata["op"].(map[string]interface{})["samples"].(map[string]interface{})

	samples = make(map[string]interface{})

	for key, value := range res {
		samples[key] = value.([]interface{})[0]
	}

	return
}
