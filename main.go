package main

import (
	"flag"
	"log"
	"net/http"
	"net/url"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	listenAddress = flag.String("web.listen", ":9118", "Address on which to expose metrics and web interface.")
	metricsPath   = flag.String("web.path", "/metrics", "Path under which to expose metrics.")
	natsURL       = flag.String("nats.addr", "http://localhost:8222/", "Address of the NATS monitoring.")
	namespace     = flag.String("namespace", "nats", "Namespace for the NATS metrics.")
)

func main() {
	flag.Parse()
	nURL, err := url.Parse(*natsURL)
	if err != nil {
		log.Fatal(err)
	}
	prometheus.MustRegister(NewExporter(nURL))

	log.Printf("Starting Server: %s", *listenAddress)
	handler := prometheus.Handler()
	if *metricsPath == "" || *metricsPath == "/" {
		http.Handle(*metricsPath, handler)
	} else {
		http.Handle(*metricsPath, handler)
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`<html>
			<head><title>NATS Exporter</title></head>
			<body>
			<h1>NATS Exporter</h1>
			<p><a href="` + *metricsPath + `">Metrics</a></p>
			</body>
			</html>`))
		})
	}

	err = http.ListenAndServe(*listenAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
}
