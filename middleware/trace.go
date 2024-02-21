package middleware

import (
	"bytes"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"net/http"
	"strings"
	"time"
)

func TraceMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		buf := &bytes.Buffer{}
		lrw := &loggingResponseWriter{ResponseWriter: w, buf: buf}

		span, _ := opentracing.StartSpanFromContext(r.Context(), fmt.Sprintf("begin [ %s ] - %s", r.Method, r.URL.Path))

		traceID := span.Context().(jaeger.SpanContext).TraceID().String()

		w.Header().Set("X-Trace-ID", traceID)

		tagRequest(span, r)

		ctx := opentracing.ContextWithSpan(r.Context(), span)
		r = r.WithContext(ctx)

		span.Finish()

		defer func() {
			tagResponse(lrw, r)
		}()

		next(lrw, r)
	}
}

func tagRequest(span opentracing.Span, r *http.Request) {

	span.SetTag("http.remote_addr", r.RemoteAddr)
	span.SetTag("http.path", r.URL.Path)
	span.SetTag("http.host", r.Host)
	span.SetTag("http.ip", strings.Split(r.RemoteAddr, ":")[0])
	span.SetTag("http.method", r.Method)

	_ = r.ParseForm()

	span.LogKV(
		"Params", r.URL.Query(),
		"Body", r.Form.Encode(),
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
	http.ResponseWriter
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
