package couchbase

import (
	"encoding/json"
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	"io/ioutil"
	"net/http"
)

type BucketMetric struct {
	Bucket string
	Metric string
	Data   interface{}
}

type BucketCollector struct {
	Url    string
	Un     string
	Pw     string
	Client *http.Client
}

func NewCollector(cfg plugin.Config) (bc *BucketCollector, err error) {
	url, err := cfg.GetString("api_url")
	un, err := cfg.GetString("username")
	pw, err := cfg.GetString("password")
	if err != nil {
		return
	}

	bc = &BucketCollector{Url: url, Un: un, Pw: pw, Client: &http.Client{}}
	return
}

func (self BucketCollector) GetSamples() (metrics []BucketMetric, err error) {

	// Make init request to /pools/default/buckets/
	var buckets []map[string]interface{}
	err = self.GetAuthRequest(self.Url, &buckets)
	if err != nil {
		return
	}

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
			metrics = append(metrics, BucketMetric{Bucket: name, Metric: key, Data: value.([]interface{})[0]})
		}
	}

	return
}

func (self BucketCollector) GetAuthRequest(rurl string, data interface{}) (err error) {
	req, err := http.NewRequest("GET", rurl, nil)
	req.SetBasicAuth(self.Un, self.Pw)
	resp, err := self.Client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, data)

	return
}
