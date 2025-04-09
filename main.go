package main

import (
	"PrometheusF6005/ont"
	internalPrometheus "PrometheusF6005/prometheus"
	"cmp"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	session, err := ont.Login(
		cmp.Or(strings.TrimRight(os.Getenv("ENDPOINT"), "/"), "http://192.168.1.1"),
		cmp.Or(os.Getenv("ONT_USERNAME"), "admin"),
		cmp.Or(os.Getenv("ONT_PASSWORD"), "admin"),
	)
	if err != nil {
		fmt.Println("Login failed:", err)
		return
	}

	collector := internalPrometheus.NewONTCollector(session)
	registry := prometheus.NewRegistry()
	registry.MustRegister(collector)

	http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{
		EnableOpenMetrics: false,
	}))

	log.Fatal(http.ListenAndServe(":80", nil))
}
