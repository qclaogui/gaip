// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package log

import (
	"github.com/go-kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

//lint:ignore faillint It's non-trivial to remove this global variable.
var experimentalFeaturesInUse = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Namespace: "qclaogui",
		Name:      "experimental_features_used_total",
		Help:      "The number of experimental features in use.",
	}, []string{"feature"},
)

// WarnExperimentalUse logs a warning and increments the experimental features metric.
func WarnExperimentalUse(feature string) {
	level.Warn(Logger).Log("msg", "experimental feature in use", "feature", feature)
	experimentalFeaturesInUse.WithLabelValues(feature).Inc()
}
