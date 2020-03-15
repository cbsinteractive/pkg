package tracing

import "net/http"

type NoopTracer struct{}

func (NoopTracer) Init() error                        { return nil }
func (NoopTracer) Client(c *http.Client) *http.Client { return c }
func (NoopTracer) Handle(_ interface{ Name(host string) string }, h http.Handler) http.Handler {
	return h
}

