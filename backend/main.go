package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/redis/go-redis/v9"
)

var (
	httpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint", "status"},
	)
	httpRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request duration in seconds",
			Buckets: []float64{0.05, 0.1, 0.25, 0.5, 1, 2.5, 5},
		},
		[]string{"method", "endpoint"},
	)
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	http.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		// Simulate business logic
		val, err := rdb.Get(context.Background(), "counter").Result()
		if err == redis.Nil {
			rdb.Set(context.Background(), "counter", 0, 0)
			val = "0"
		}
		
		counter, _ := strconv.Atoi(val)
		counter++
		rdb.Set(context.Background(), "counter", counter, 0)
		
		duration := time.Since(start).Seconds()
		httpRequestDuration.WithLabelValues(r.Method, "/api/data").Observe(duration)
		httpRequestsTotal.WithLabelValues(r.Method, "/api/data", "200").Inc()
		
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"counter":%d, "latency_ms":%.2f}`, counter, duration*1000)
	})

	http.Handle("/metrics", promhttp.Handler())
	
	log.Println("Backend running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}