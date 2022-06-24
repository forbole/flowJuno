package logging

import (
	"github.com/prometheus/client_golang/prometheus"
)

// StartHeight represents the Telemetry counter used to set the start height of the parsing
var StartHeight = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "juno_initial_height",
		Help: "Initial parsing height.",
	},
)

// WorkerCount represents the Telemetry counter used to track the worker count
var WorkerCount = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "juno_worker_count",
		Help: "Number of active workers.",
	},
)

// WorkerHeight represents the Telemetry counter used to track the last indexed height for each worker
var WorkerHeight = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "juno_last_indexed_height",
		Help: "Height of the last indexed block.",
	},
	[]string{"worker_index"},
)

// ErrorCount represents the Telemetry counter used to track the number of errors emitted
var ErrorCount = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "juno_error_count",
		Help: "Total number of errors emitted.",
	},
)

func init() {
	err := prometheus.Register(StartHeight)
	if err != nil {
		panic(err)
	}

	err = prometheus.Register(WorkerCount)
	if err != nil {
		panic(err)
	}

	err = prometheus.Register(WorkerHeight)
	if err != nil {
		panic(err)
	}

	err = prometheus.Register(ErrorCount)
	if err != nil {
		panic(err)
	}
}
