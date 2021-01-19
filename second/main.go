package main

import (
	"fmt"
	"net/http"
)

// 0.0.0.0:8080 + pattern

/*
	POST - create metricsSlice
	PUT - update metricsSlice
	GET - get metricsSlice
	POST - list metricsSlice
	DELETE - delete metricsSlice
*/

func main() {
	metricsService := &metricsService{}
	metricsService.Init()

	m := http.NewServeMux()
	m.Handle("/metrics/", MakeHandler(*metricsService))
	http.Handle("/", m)

	fmt.Println("listening on port :8080")
	http.ListenAndServe(":8080", m)
}

