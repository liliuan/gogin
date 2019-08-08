package lib

import (
	"fmt"
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/model"
	"github.com/openzipkin/zipkin-go/reporter"
	reporterHttp "github.com/openzipkin/zipkin-go/reporter/http"
	"gogin/config"
)

var (
	Tracer   *zipkin.Tracer
	Reporter reporter.Reporter
)

func init() {
	serviceName := "api"
	config := common.GetConfig()

	Reporter = reporterHttp.NewReporter(config.Reporter)

	// create endpoint
	endpoint, err := zipkin.NewEndpoint(serviceName, config.HostPort)
	if err != nil {
		fmt.Println("unable to create local endpoint: ", err)
	}

	// init tracer
	Tracer, err = zipkin.NewTracer(Reporter, zipkin.WithLocalEndpoint(endpoint))
	if err != nil {
		fmt.Println("unable to create tracer: ", err)
	}
}

func GetSpan(serviceName, name string, ctx model.SpanContext, kind model.Kind) (zipkin.Span, reporter.Reporter) {
	b := common.GetIsZipkinFlush()
	endporint := &model.Endpoint{}
	endporint.ServiceName = "api:" + serviceName
	childSpan := Tracer.StartSpan(name, zipkin.Parent(ctx), zipkin.FlushOnFinish(b), zipkin.Kind(kind))
	childSpan.SetRemoteEndpoint(endporint)
	return childSpan, Reporter
}
