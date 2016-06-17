# NATS Exporter

[![GoDoc](https://godoc.org/github.com/lovoo/nats_exporter?status.svg)](https://godoc.org/github.com/lovoo/nats_exporter)

NATS exporter for [Prometheus](https://prometheus.io/), written in Go. It extracts several metrics and provides them via the standard HTTP interface for Prometheus to collect them.

Prometheus enables you to monitor those values and add alerting, if necessary.

## Included Metrics

* Active connections to gnatsd
* Active routes to gnatsd
* Count of transferred messages (in/out)
* Amount of transmitted data (in/out)
* Number of slow consumers

## Usage

	Usage of nats_exporter:
	  -namespace string
	    	Namespace for the NATS metrics. (default "nats")
	  -nats.addr string
	    	Address of the NATS monitoring. (default "http://localhost:8222/")
	  -web.listen string
	    	Address on which to expose metrics and web interface. (default ":9118")
	  -web.path string
	    	Path under which to expose metrics. (default "/metrics")

### Building from Source

    go get -u github.com/lovoo/nats_exporter
    go install github.com/lovoo/nats_exporter
    $GOPATH/bin/nats_exporter

### With Docker

    docker run -d --name nats_exporter -p 9118:9118 lovoo/nats_exporter:latest -nats.addr http://somehost:8222

## Contributing

1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request
