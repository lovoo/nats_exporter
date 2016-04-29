# NATS Exporter

[![GoDoc](https://godoc.org/github.com/lovoo/nats_exporter?status.svg)](https://godoc.org/github.com/lovoo/nats_exporter)

NATS exporter for prometheus.io, written in go.

## Usage

    docker run -d --name nats_exporter -p 9118:9118 lovoo/nats_exporter:latest -nats.addr http://somehost:8222

## Building

    go get -u github.com/lovoo/nats_exporter
    go install github.com/lovoo/nats_exporter

## Contributing

1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request
