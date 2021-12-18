package metrics

import "github.com/prometheus/client_golang/prometheus"

func LastDeployments() (z *prometheus.GaugeVec) {
	result := GetLastXDaysDeployments()
	z = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "versions",
			Subsystem: "last",
			Name:      "deployments_by_status",
			Help:      "Get last 10 days deployments by status",
		},
		[]string{
			"workload",
			"platform",
			"environment",
			"status",
			"date",
		},
	)

	for _, v := range result {
		z.With(prometheus.Labels{
			"workload":    v.Workload,
			"platform":    v.Platform,
			"environment": v.Environment,
			"status":      v.Status,
			"date":        v.Date.String(),
		}).Set(float64(v.Total))
	}
	return
}
