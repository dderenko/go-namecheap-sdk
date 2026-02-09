package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	Logs = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "logs",
			Help: "Number of log messages by level",
		},
		[]string{"service", "level", "message"},
	)

	LogsSkipped = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "logs_skipped",
			Help: "Number of skipped log messages.",
		},
		[]string{"service", "level"},
	)
)

// Register registers metrics manually
func Register() error {
	for _, m := range []prometheus.Collector{Logs, LogsSkipped} {
		err := prometheus.Register(m)
		if err != nil {
			return err
		}
	}
	return nil
}
