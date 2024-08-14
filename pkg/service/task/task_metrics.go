// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package task

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type taskMetrics struct {
	requestDuration prometheus.Histogram
	records         *prometheus.GaugeVec
}

func newTaskMetrics(reg prometheus.Registerer) taskMetrics {
	var m taskMetrics

	m.requestDuration = promauto.With(reg).NewHistogram(prometheus.HistogramOpts{
		Name: "gaip_task_request_duration_seconds",
		Help: "Time spent request a full cycle.",

		NativeHistogramBucketFactor: 1.1,
	})

	m.records = promauto.With(reg).NewGaugeVec(prometheus.GaugeOpts{
		Name: "gaip_task_records",
		Help: "The number of task records.",
	}, []string{"topic"})

	return m
}
