package initialize

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegerConfig "github.com/uber/jaeger-client-go/config"
	"gorm-example/config"
	"io"
)

func InitJaeger(config *config.Runtime) (closer io.Closer, err error) {

	if opentracing.IsGlobalTracerRegistered() {
		return
	}

	// TODO 线上环境，调用阿里云，腾讯云等云服务商等链路追踪

	// 根据配置初始化Tracer 返回Closer
	tracer, closer, err := (&jaegerConfig.Configuration{
		ServiceName: config.ServerName,
		Disabled:    false,
		Sampler: &jaegerConfig.SamplerConfig{
			Type: jaeger.SamplerTypeConst,
			// param的值在0到1之间，设置为1则将所有的Operation输出到Reporter
			Param: 1,
		},
		Reporter: &jaegerConfig.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: config.JaegerHostPort,
		},
	}).NewTracer()
	if err != nil {
		return
	}

	// 设置全局Tracer - 如果不设置将会导致上下文无法生成正确的Span
	opentracing.SetGlobalTracer(tracer)
	return closer, err
}
