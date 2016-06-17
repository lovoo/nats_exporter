package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCollect(t *testing.T) {
	var natsResponse = `{
  "server_id": "1234ab567894a9de520315172d46cc80",
  "version": "0.7.2",
  "go": "go1.5.2",
  "host": "0.0.0.0",
  "auth_required": false,
  "ssl_required": false,
  "tls_required": false,
  "tls_verify": false,
  "max_connections": 100,
  "ping_interval": 120000000000,
  "ping_max": 2,
  "http_port": 8222,
  "https_port": 0,
  "max_control_line": 512,
  "max_pending_size": 10000000,
  "cluster_port": 4244,
  "tls_timeout": 0.5,
  "port": 4242,
  "max_payload": 65536,
  "start": "2016-05-04T09:39:52.155400146Z",
  "now": "2016-06-17T09:40:35.231047305Z",
  "uptime": "44d0h0m43s",
  "mem": 84230144,
  "cores": 40,
  "cpu": 39.5,
  "connections": 1,
  "routes": 3,
  "remotes": 3,
  "in_msgs": 1023842918,
  "out_msgs": 998024171,
  "in_bytes": 61818435048,
  "out_bytes": 54370633082,
  "slow_consumers": 0
}`

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, natsResponse)
	}))
	defer srv.Close()

	exporter := Exporter{
		NatsURL: srv.URL,
	}

	if err := exporter.collect(); err != nil {
		t.Fatal(err)
	}
}
