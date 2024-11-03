package main

import (
	"flag"
	"log"
	"os"
	"strings"

	prometheus "github.com/minhdanh/speedtest-to-prom/pkg/prometheus"
)

func main() {
	metricName := flag.String("metric_name", "", "Name of the metric")
	metricLabels := flag.String("metric_labels", "", "Comma-separated metric labels")
	value := flag.Float64("value", 0.0, "Value of the metric")

	// Parse the command-line flags
	flag.Parse()

	// Check if metric_name is provided
	if *metricName == "" {
		log.Fatal("metric_name is required")
	}

	// Get credentials from environment variables
	promUsername := os.Getenv("PROM_USERNAME")
	promPassword := os.Getenv("PROM_PASSWORD")
	if promUsername == "" || promPassword == "" {
		log.Fatal("PROM_USERNAME and PROM_PASSWORD environment variables must be set")
	}

	metrics := map[string][]prometheus.MetricValue{
		*metricName: {
			{
				Value:  *value,
				Labels: ConvertLabels(*metricLabels),
			},
		},
	}

	err := prometheus.PushPrometheusMetrics(promUsername, promPassword, metrics, nil)
	if err != nil {
		log.Fatalf("Error pushing metrics: %v", err)
	}

	log.Println("Metrics pushed successfully")
}

// ConvertLabels converts a comma-separated string of labels into a map[string]string
func ConvertLabels(labelsStr string) map[string]string {
	labelsMap := make(map[string]string)

	// Split the string by commas
	labels := strings.Split(labelsStr, ",")

	// Iterate over each label
	for _, label := range labels {
		// Split each label by '='
		parts := strings.SplitN(strings.TrimSpace(label), "=", 2)
		if len(parts) == 2 {
			// Add to the map
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			labelsMap[key] = value
		}
	}

	return labelsMap
}
