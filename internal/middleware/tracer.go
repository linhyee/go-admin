package middleware

import (
	"blog/global"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
)

func Tracing() func(c *gin.Context) {
	return func(c *gin.Context) {
		if !global.ServerSetting.IsTracing || global.Tracer == nil {
			c.Next()
			return
		}
		var newCtx context.Context
		var span opentracing.Span
		spanCtx, err := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(c.Request.Header),
		)
		if err != nil {
			span, newCtx = opentracing.StartSpanFromContextWithTracer(
				c.Request.Context(),
				global.Tracer,
				c.Request.URL.Path,
			)
		} else {
			span, newCtx = opentracing.StartSpanFromContextWithTracer(
				c.Request.Context(),
				global.Tracer,
				c.Request.URL.Path,
				opentracing.ChildOf(spanCtx),
				opentracing.Tag{
					Key:   string(ext.Component),
					Value: "HTTP",
				},
			)
		}
		defer span.Finish()
		var traceId, spanId string
		spanContext := span.Context()
		switch v := spanContext.(type) {
		case jaeger.SpanContext:
			traceId = v.TraceID().String()
			spanId = v.SpanID().String()
		}
		c.Set("X-Trace-ID", traceId)
		c.Set("X-Span-ID", spanId)
		c.Request = c.Request.WithContext(newCtx)
		c.Next()
	}
}
