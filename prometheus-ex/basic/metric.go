package main

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
)

type Service struct {
	httpReq *prometheus.HistogramVec
}

// 创建需要汇报的包 并注册到 prometheus 的默认 handler
func NewPrometheusService() (*Service, error) {
	http := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "test",
		Name:      "my_faver2",
		Help:      "durations of a request in second",
		Buckets:   prometheus.DefBuckets,
	}, []string{"url", "http_method", "status_code"})
	s := &Service{
		httpReq: http,
	}
	err := prometheus.Register(s.httpReq)
	return s, err
}

type httpReqInstance struct {
	URL        string
	Method     string
	StatusCode string
	Duration   float64 `json:"duration"`
}

func (s *Service) report(h httpReqInstance) {
	log.Println("report once")
	s.httpReq.WithLabelValues(h.URL, h.Method, h.StatusCode).Observe(h.Duration)
}
