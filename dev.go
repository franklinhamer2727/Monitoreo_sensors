package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metric struct {
	name string
	p    *prometheus.GaugeVec
}

var metrics []*Metric

///para poder crear un nuevo vector
func newVec(name string) *prometheus.GaugeVec {
	labels := []string{"id"}
	p := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: name,
		Help: name,
	}, labels)
	prometheus.MustRegister(p)
	return p
}
func isInMetrics(name string) bool {
	for i := 0; i < len(metrics); i++ {
		if metrics[i].name == name {
			return true
		}
	}
	return false
}
func NewMetric(d []Data) {
	for _, data := range d {

		for key, _ := range data.Params {
			if !isInMetrics(key) {
				metrics = append(metrics, &Metric{
					name: key,
					p:    newVec(key),
				})
			}
		}

	}
}
