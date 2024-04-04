package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"io"
	"net/http"
	"strings"
	"time"
)

func TraceMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		buf := &bytes.Buffer{}
		lrw := &loggingResponseWriter{
			ResponseWriter: c.Writer,
			buf:            buf,
			start:          time.Now(),
		}
		c.Writer = lrw

		span, _ := opentracing.StartSpanFromContext(c, fmt.Sprintf("begin [ %s ] - %s", c.Request.Method, c.Request.URL.Path))
		defer span.Finish()

		traceID := span.Context().(jaeger.SpanContext).TraceID().String()

		c.Writer.Header().Set("X-Trace-ID", traceID)

		tagRequest(span, c.Request)

		// 将新的 Span 上下文注入到当前请求的上下文中
		ctx := opentracing.ContextWithSpan(c.Request.Context(), span)

		// 使用带有新的 Span 上下文的请求上下文继续处理请求
		c.Request = c.Request.WithContext(ctx)

		c.Next()

		defer func() {
			tagResponse(lrw, c.Request)
		}()
	}
}

func tagRequest(span opentracing.Span, r *http.Request) {

	span.SetTag("http.remote_addr", r.RemoteAddr)
	span.SetTag("http.path", r.URL.Path)
	span.SetTag("http.host", r.Host)
	span.SetTag("http.ip", strings.Split(r.RemoteAddr, ":")[0])
	span.SetTag("http.method", r.Method)
	span.SetTag("http.content_type", r.Header.Get("Content-Type"))

	_ = r.ParseForm()
	body, _ := io.ReadAll(r.Body)

	span.LogKV(
		"Params", r.Form.Encode(),
		"Body", string(body),
	)
}

func tagResponse(lrw *loggingResponseWriter, r *http.Request) {

	span, _ := opentracing.StartSpanFromContext(r.Context(), "Log Response")

	span.SetTag("http.status_code", lrw.statusCode)
	span.LogKV(
		"Response", lrw.buf.String(),
	)
	span.Finish()
}

type loggingResponseWriter struct {
	gin.ResponseWriter
	buf        *bytes.Buffer
	statusCode int
	start      time.Time
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	// 记录状态码
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func (lrw *loggingResponseWriter) Write(b []byte) (int, error) {
	// 将响应写入缓冲区
	lrw.buf.Write(b)
	return lrw.ResponseWriter.Write(b)
}
