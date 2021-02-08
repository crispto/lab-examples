package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var start time.Time

func main() {

	start = time.Now()

	metricService, err := NewPrometheusService()
	if err != nil {
		panic(err)
	}
	// 为已注册的 汇报项开启 handler
	http.Handle("/metrics", promhttp.Handler())
	// 每次访问 url 记录一次
	http.HandleFunc("/count", AddMetrics(metricService, Count))
	// 这个 metric 是随时间记录
	go func() {
		ticker := time.NewTicker(time.Millisecond * 100)
		defer ticker.Stop()
		for i := 0; i < 100; i++ {
			select {
			case <-ticker.C:
				appMetric := httpReqInstance{
					URL:        "/hello",
					Method:     http.MethodGet,
					StatusCode: "200",
				}
				if i < 10 {
					appMetric.Duration = float64(.0005)
				} else if i < 20 {
					appMetric.Duration = float64(.005)
				} else if i < 30 {
					appMetric.Duration = float64(.5)
				} else if i < 40 {
					appMetric.Duration = float64(1)
				} else if i < 60 {
					appMetric.Duration = float64(2)
				} else if i < 80 {
					appMetric.Duration = float64(5.3)
				} else {
					appMetric.Duration = float64(10.33)
				}

				metricService.report(appMetric)

			}
		}

	}()
	http.ListenAndServe(":8083", nil)
}

type hHandlerFunc func(w *hResponseWriter, r *http.Request)

func AddMetrics(h *Service, next hHandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		appMetric := httpReqInstance{
			URL:    r.URL.Path,
			Method: r.Method,
		}
		cw := hResponseWriter{
			ResponseWriter: w,
		}
		now := time.Now()
		next(&cw, r)
		ela := time.Now().Sub(now).Milliseconds()
		appMetric.StatusCode = strconv.Itoa(cw.code)
		appMetric.StatusCode = "200"
		appMetric.Duration = float64(ela)
		h.report(appMetric)
	})
}

type hResponseWriter struct {
	http.ResponseWriter
	code int
}

func (w *hResponseWriter) WriteHeader(code int) {
	w.ResponseWriter.WriteHeader(code)
	w.code = code
}

func Count(w *hResponseWriter, r *http.Request) {
	elapse := time.Now().Sub(start).Milliseconds()
	get := elapse % 1000
	time.Sleep(time.Duration(get) * time.Millisecond)
	log.Println("count is accessed")
	if get%200 == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		w.code = 500
		io.WriteString(w, "server error")
	} else {
		w.code = 200
		io.WriteString(w, "finish")
	}
}
