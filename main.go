package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"contrib.go.opencensus.io/exporter/stackdriver"
	"contrib.go.opencensus.io/exporter/stackdriver/propagation"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"
)

func main() {
	exporter, err := stackdriver.NewExporter(stackdriver.Options{
		ProjectID: "souzoh-demo-gcp-001",
	})
	if err != nil {
		panic(err)
	}
	trace.RegisterExporter(exporter)

	for {
		ctx := context.Background()
		msg, err := process(ctx)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(msg)

		time.Sleep(5 * time.Second)
	}
}

func process(ctx context.Context) (string, error) {
	ctx, span := trace.StartSpan(ctx, "/process")
	defer span.End()

	client := &http.Client{
		Transport: &ochttp.Transport{
			// Use Google Cloud propagation format.
			Propagation: &propagation.HTTPFormat{},
		},
	}

	req, _ := http.NewRequest("GET", "http://backendhellotime-service.default.svc.cluster.local:8080", nil)

	// The trace ID from the incoming request will be
	// propagated to the outgoing request.
	req = req.WithContext(ctx)

	// The outgoing request will be traced with r's trace ID.
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
