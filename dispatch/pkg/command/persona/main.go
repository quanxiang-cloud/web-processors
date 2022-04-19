package main

import (
	"context"
	"flag"
	"os"
	"time"

	"github.com/quanxiang-cloud/cabin/tailormade/client"
)

var (
	url          = flag.String("url", "http://persona/api/v1/persona/batchSetValue", "url")
	timeout      = flag.Duration("timeout", 10*time.Second, "timeout")
	maxIdleConns = flag.Int("maxIdleConns", 10, "maxIdleConns")
	key          = flag.String("key", "", "key")
	value        = flag.String("value", "", "value")
)

func main() {
	flag.Parse()

	a := client.New(client.Config{
		Timeout:      *timeout,
		MaxIdleConns: *maxIdleConns,
	})

	err := client.POST(context.Background(), &a, *url, struct {
		Params []map[string]interface{} `json:"params"`
	}{
		Params: []map[string]interface{}{
			{
				"version": "1.0",
				"key":     *key,
				"value":   *value,
			},
		},
	},
		&struct{}{},
	)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.Sync()

		return
	}
}
