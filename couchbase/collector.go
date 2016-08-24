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

	self.samples = &BucketStats{Url: url, Un: un, Pw: pw}

	return self
}

// metricCollector implements logic for discovering available metrics
type metricCollector struct {
	samples metrictap
}

// Collect returns map of metric values (accessible by metric name).
// If any of requested calls fail error is returned.
func (mc *metricCollector) Collect() (map[string]interface{}, error) {
	samples, err := mc.samples.GetSamples()
	if err != nil {
		return nil, err
	}

	return samples, nil
}

// Discover performs metric discovery. Returns valid metric names and
// if mandatory request fails error is returned.
func (mc *metricCollector) Discover() ([]string, error) {
	samples, err := mc.samples.GetSamples()
	if err != nil {
		return nil, err
	}

	res := []string{}

	for key := range samples {
		res = append(res, key)
	}

	return res, nil
}

// Interface for couchbase api taps
type metrictap interface {
	GetSamples() (samples map[string]interface{}, err error)
}

// BucketStats
type BucketStats struct {
	Url string
	Un  string
	Pw  string
}

func (self BucketStats) GetSamples() (samples map[string]interface{}, err error) {

	// Make init request to /pools/default/buckets/
	var buckets []map[string]interface{}
	self.GetAuthRequest(self.Url, &buckets)

	samples = make(map[string]interface{})

	// For each bucket get name and make /stats request
	for _, b := range buckets {
		name := b["name"].(string)

		var stats map[string]interface{}
		rurl := self.Url + name + "/stats"

		err = self.GetAuthRequest(rurl, &stats)
		if err != nil {
			return
		}

		res := stats["op"].(map[string]interface{})["samples"].(map[string]interface{})

		for key, value := range res {
			sk := name + "/" + key
			samples[sk] = value.([]interface{})[0]
		}
	}

	return
}

func (self BucketStats) GetAuthRequest(rurl string, data interface{}) (err error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", rurl, nil)
	req.SetBasicAuth(self.Un, self.Pw)
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, data)

	return
}
