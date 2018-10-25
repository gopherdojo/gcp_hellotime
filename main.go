package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"contrib.go.opencensus.io/exporter/stackdriver"
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

	res, err := http.Get("http://backendhellotime-service.default.svc.cluster.local:8080")
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
