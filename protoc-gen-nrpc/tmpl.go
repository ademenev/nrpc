package main

const tFile = `// This code was autogenerated from {{.GetName}}, do not edit.

{{- $pkgName := GoPackageName .}}
package {{$pkgName}}

import (
	"context"
	"log"
	"time"

	"github.com/golang/protobuf/proto"
	nats "github.com/nats-io/go-nats"
	{{- if Prometheus}}
	"github.com/prometheus/client_golang/prometheus"
	{{- end}}
	"github.com/rapidloop/nrpc"
)

{{range .Service -}}
// {{.GetName}}Server is the interface that providers of the service
// {{.GetName}} should implement.
type {{.GetName}}Server interface {
	{{- range .Method}}
	{{.GetName}}(ctx context.Context, req {{GetPkg $pkgName .GetInputType}}) (resp {{GetPkg $pkgName .GetOutputType}}, err error)
	{{- end}}
}

{{- if Prometheus}}

var (
	// The request completion time, measured at client-side.
	clientRCTFor{{.GetName}} = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "nrpc_client_request_completion_time_seconds",
			Help:       "The request completion time for {{.GetName}} calls, measured client-side.",
			Objectives: map[float64]float64{0.9: 0.01, 0.95: 0.01, 0.99: 0.001},
			ConstLabels: map[string]string{
				"service": "{{.GetName}}",
			},
		},
		[]string{"method"})

	// The handler execution time, measured at server-side.
	serverHETFor{{.GetName}} = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "nrpc_server_handler_execution_time_seconds",
			Help:       "The handler execution time for {{.GetName}} calls, measured server-side.",
			Objectives: map[float64]float64{0.9: 0.01, 0.95: 0.01, 0.99: 0.001},
			ConstLabels: map[string]string{
				"service": "{{.GetName}}",
			},
		},
		[]string{"method"})

	// The counts of calls made by the client, classified by result type.
	clientCallsFor{{.GetName}} = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "nrpc_client_calls_count",
			Help: "The count of calls made by the client.",
			ConstLabels: map[string]string{
				"service": "{{.GetName}}",
			},
		},
		[]string{"method", "result_type"})

	// The counts of requests handled by the server, classified by result type.
	serverRequestsFor{{.GetName}} = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "nrpc_server_requests_count",
			Help: "The count of requests handled by the server.",
			ConstLabels: map[string]string{
				"service": "{{.GetName}}",
			},
		},
		[]string{"method", "result_type"})
)
{{- end}}

// {{.GetName}}Handler provides a NATS subscription handler that can serve a
// subscription using a given {{.GetName}}Server implementation.
type {{.GetName}}Handler struct {
	ctx    context.Context
	nc     *nats.Conn
	server {{.GetName}}Server
}

func New{{.GetName}}Handler(ctx context.Context, nc *nats.Conn, s {{.GetName}}Server) *{{.GetName}}Handler {
	return &{{.GetName}}Handler{
		ctx:    ctx,
		nc:     nc,
		server: s,
	}
}

func (h *{{.GetName}}Handler) Subject() string {
	return "{{.GetName}}.*"
}

func (h *{{.GetName}}Handler) Handler(msg *nats.Msg) {
	// extract method name from subject
	name := nrpc.ExtractFunctionName(msg.Subject)

	// call handler and form response
	var resp proto.Message
	var errstr string
{{- if Prometheus}}
	var elapsed float64
{{- end}}
	switch name {
	{{- $serviceName := .GetName}}{{- range .Method}}
	case "{{.GetName}}":
		var req {{GetPkg $pkgName .GetInputType}}
		if err := proto.Unmarshal(msg.Data, &req); err != nil {
			log.Printf("{{.GetName}}Handler: {{.GetName}} request unmarshal failed: %v", err)
			errstr = "bad request received: " + err.Error()
{{- if Prometheus}}
			serverRequestsFor{{$serviceName}}.WithLabelValues("{{.GetName}}",
				"protobuf_fail").Inc()
{{- end}}
		} else {
{{- if Prometheus}}
			start := time.Now()
{{- end}}
			innerResp, err := h.server.{{.GetName}}(h.ctx, req)
{{- if Prometheus}}
			elapsed = time.Since(start).Seconds()
{{- end}}
			if err != nil {
				log.Printf("{{.GetName}}Handler: {{.GetName}} handler failed: %v", err)
				errstr = err.Error()
{{- if Prometheus}}
				serverRequestsFor{{$serviceName}}.WithLabelValues("{{.GetName}}",
					"handler_fail").Inc()
{{- end}}
			} else {
				resp = &innerResp
			}
		}
{{- end}}
	default:
		log.Printf("{{.GetName}}Handler: unknown name %q", name)
		errstr = "unknown name: " + name
{{- if Prometheus}}
		serverRequestsFor{{.GetName}}.WithLabelValues("{{.GetName}}",
			"name_fail").Inc()
{{- end}}
	}

	// encode and send response
	err := nrpc.Publish(resp, errstr, h.nc, msg.Reply) // error is logged
{{- if Prometheus}}
	if err != nil {
		serverRequestsFor{{$serviceName}}.WithLabelValues(name, "protobuf_fail").Inc()
	} else if len(errstr) == 0 {
		serverRequestsFor{{$serviceName}}.WithLabelValues(name, "success").Inc()
	}

	// report metric to Prometheus
	serverHETFor{{$serviceName}}.WithLabelValues(name).Observe(elapsed)
{{- else}}
	if err != nil {
		log.Println("{{.GetName}}Handler: {{.GetName}} handler failed to publish the response: %s", err)
	}
{{- end}}
}

type {{.GetName}}Client struct {
	nc      *nats.Conn
	Subject string
	Timeout time.Duration
}

func New{{.GetName}}Client(nc *nats.Conn) *{{.GetName}}Client {
	return &{{.GetName}}Client{
		nc:      nc,
		Subject: "{{.GetName}}",
		Timeout: 5 * time.Second,
	}
}
{{$serviceName := .GetName}}
{{- range .Method}}
func (c *{{$serviceName}}Client) {{.GetName}}(req {{GetPkg $pkgName .GetInputType}}) (resp {{GetPkg $pkgName .GetOutputType}}, err error) {
{{- if Prometheus}}
	start := time.Now()
{{- end}}

	// call
	respBytes, err := nrpc.Call(&req, c.nc, c.Subject+".{{.GetName}}", c.Timeout)
	if err != nil {
{{- if Prometheus}}
		clientCallsFor{{$serviceName}}.WithLabelValues("{{.GetName}}",
			"call_fail").Inc()
{{- end}}
		return // already logged
	}

	// decode inner reponse
	if err = proto.Unmarshal(respBytes, &resp); err != nil {
		log.Printf("{{.GetName}}: response unmarshal failed: %v", err)
{{- if Prometheus}}
		clientCallsFor{{$serviceName}}.WithLabelValues("{{.GetName}}",
			"protobuf_fail").Inc()
{{- end}}
		return
	}

{{- if Prometheus}}

	// report total time taken to Prometheus
	elapsed := time.Since(start).Seconds()
	clientRCTFor{{$serviceName}}.WithLabelValues("{{.GetName}}").Observe(elapsed)
	clientCallsFor{{$serviceName}}.WithLabelValues("{{.GetName}}",
		"success").Inc()
{{- end}}

	return
}
{{end -}}
{{- end -}}

{{- if Prometheus}}
func init() {
{{- range .Service}}
	// register metrics for service {{.GetName}}
	prometheus.MustRegister(clientRCTFor{{.GetName}})
	prometheus.MustRegister(serverHETFor{{.GetName}})
	prometheus.MustRegister(clientCallsFor{{.GetName}})
	prometheus.MustRegister(serverRequestsFor{{.GetName}})
{{- end}}
}
{{- end}}`
