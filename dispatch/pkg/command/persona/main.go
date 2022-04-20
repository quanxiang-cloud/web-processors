package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/quanxiang-cloud/cabin/tailormade/client"
)

var (
	endpoint     = flag.String("endpoint", "http://persona", "endpoint")
	api          = flag.String("api", "api/v1/persona/batchSetValue", "api")
	timeout      = flag.Duration("timeout", time.Duration(10)*time.Second, "timeout")
	maxIdleConns = flag.Int("maxIdleConns", 10, "maxIdleConns")
	version      = flag.String("version", "", "version")
	key          = flag.String("key", "", "key")
	value        = flag.String("value", "", "value")
)

func main() {
	flag.Parse()

	a := client.New(client.Config{
		Timeout:      *timeout,
		MaxIdleConns: *maxIdleConns,
	})

	url := fmt.Sprintf("%s/%s", *endpoint, *api)
	err := client.POST(context.Background(), &a, url, struct {
		Params []map[string]interface{} `json:"params"`
	}{
		Params: []map[string]interface{}{
			{
				"version": *version,
				"key":     *key,
				"value":   *value,
			},
		},
	},
		&struct{}{},
	)
	if err != nil {
		panic(err)
	}
}
