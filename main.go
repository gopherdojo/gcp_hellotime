package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	for {
		res, err := http.Get("http://backendhellotime-service.default.svc.cluster.local:8080")
		if err != nil {
			fmt.Println(err)
		}
		defer res.Body.Close()
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(b))

		time.Sleep(5 * time.Second)
	}
}
