package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/prometheus/client_golang/prometheus"
)

type natsMetrics struct {
	Connections float64 `json:"connections"`
	Routes      float64 `json:"routes"`

	MessagesIn  float64 `json:"in_msgs"`
	MessagesOut float64 `json:"out_msgs"`

	BytesIn  float64 `json:"in_bytes"`
	BytesOut float64 `json:"out_bytes"`

	SlowConsumers float64 `json:"slow_consumers"`
}

var (
	connections = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "gnatsd",
		Name:      "connections",
		Help:      "Active connections to gnatsd",
	})

	routes = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "gnatsd",
		Name:      "routes",
		Help:      "Active routes to gnatsd",
	})

	messageCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "gnatsd",
		Name:      "msg_total",
		Help:      "Count of transferred messages",
	},
		[]string{"direction"},
	)

	bytesCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "gnatsd",
		Name:      "bytes_total",
		Help:      "Amount of transmitted data",
	},
		[]string{"direction"},
	)

	slowConsumers = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "gnatsd",
		Name:      "slow_consumers",
		Help:      "Number of slow consumers",
	})
)

type Exporter struct {
	NatsURL string
}

func NewExporter(natsURL *url.URL) *Exporter {
	natsURL.Path = "/varz"
	return &Exporter{
		NatsURL: natsURL.String(),
	}
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	connections.Describe(ch)
	routes.Describe(ch)
	messageCounter.Describe(ch)
	bytesCounter.Describe(ch)
	slowConsumers.Describe(ch)
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	e.collect()

	connections.Collect(ch)
	routes.Collect(ch)
	messageCounter.Collect(ch)
	bytesCounter.Collect(ch)
	slowConsumers.Collect(ch)
}

func (e *Exporter) collect() {
	var metrics natsMetrics
	resp, err := http.DefaultClient.Get(e.NatsURL)
	if err != nil {
		log.Printf("could not retrieve NATS metrics: %v", err)
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&metrics)
	if err != nil {
		log.Printf("could not decode NATS metrics: %v", err)
	}

	connections.Set(metrics.Connections)
	routes.Set(metrics.Routes)

	messageCounter.WithLabelValues("in").Add(metrics.MessagesIn)
	messageCounter.WithLabelValues("out").Add(metrics.MessagesOut)

	bytesCounter.WithLabelValues("in").Add(metrics.BytesIn)
	bytesCounter.WithLabelValues("out").Add(metrics.BytesOut)

	slowConsumers.Set(metrics.SlowConsumers)
}
