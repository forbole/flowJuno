package telemetry

import (
	"fmt"
	"net/http"

	"github.com/forbole/flowJuno/types"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// RunAdditionalOperations runs the module additional operations
func RunAdditionalOperations(cfg types.Config) error {
	err := checkConfig(cfg)
	if err != nil {
		return err
	}
	if !cfg.GetTelemetryConfig().GetEnable() {
		return nil
	}

	go startPrometheus(cfg)

	return nil
}

// checkConfig checks if the given config is valid
func checkConfig(cfg types.Config) error {
	if cfg.GetTelemetryConfig() == nil {
		return fmt.Errorf("no telemetry config found")
	}

	return nil
}

// startPrometheus starts a Prometheus server using the given configuration
func startPrometheus(cfg types.Config) {
	router := mux.NewRouter()
	router.Handle("/metrics", promhttp.Handler())

	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.GetTelemetryConfig().GetPort()), router)
	if err != nil {
		panic(err)
	}
}
