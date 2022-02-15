package metrics

import (
	"fmt"
	"github.com/blang/semver"
	"github.com/prometheus/client_golang/prometheus"
	"sync"
)

const FeatureEnabled = 1
const FeatureDisabled = 0

var (
	constLabels = map[string]string{
		"collector": "anx-provider-collector",
	}
	descProviderBuild = prometheus.NewDesc(getFQDNMetricName("provider_build"),
		"information about the build version of a specific provider", []string{"name", "version"}, constLabels)

	descProviderFeatures = prometheus.NewDesc(getFQDNMetricName("feature"), "provider features and their state",
		[]string{"name"}, constLabels)
)

type ProviderMetrics struct {
	m               *sync.RWMutex
	Name            string
	providerVersion prometheus.Metric
	featureState    map[string]prometheus.Metric
	descriptions    []*prometheus.Desc
}

// NewProviderMetrics returns a prometheus.Collector for Provider Metrics.
func NewProviderMetrics(providerName, providerVersion string) ProviderMetrics {
	description := []*prometheus.Desc{descProviderBuild, descProviderFeatures}

	versionMetric := prometheus.MustNewConstMetric(descProviderBuild, prometheus.GaugeValue,
		1, providerName, providerVersion)

	return ProviderMetrics{
		Name:            providerName,
		providerVersion: versionMetric,
		descriptions:    description,
		m:               &sync.RWMutex{},
		featureState:    map[string]prometheus.Metric{},
	}
}

func (p *ProviderMetrics) Describe(descs chan<- *prometheus.Desc) {
	for _, description := range p.descriptions {
		descs <- description
	}
}

func (p *ProviderMetrics) Collect(metrics chan<- prometheus.Metric) {
	metrics <- p.providerVersion

	p.m.RLock()
	defer p.m.RUnlock()
	for _, counter := range p.featureState {
		metrics <- counter
	}
}

func (p *ProviderMetrics) MarkFeatureEnabled(featureName string) {
	p.m.Lock()
	defer p.m.Unlock()

	featureState := prometheus.MustNewConstMetric(descProviderFeatures, prometheus.CounterValue, FeatureEnabled, featureName)
	p.featureState[featureName] = featureState
}

func (p *ProviderMetrics) MarkFeatureDisabled(featureName string) {
	p.m.Lock()
	defer p.m.Unlock()

	featureState := prometheus.MustNewConstMetric(descProviderFeatures, prometheus.CounterValue, FeatureDisabled, featureName)
	p.featureState[featureName] = featureState
}

func (p *ProviderMetrics) Create(_ *semver.Version) bool {
	return true
}

func (p *ProviderMetrics) ClearState() {}

func (p *ProviderMetrics) FQName() string {
	return "cloud_provider"
}

func getFQDNMetricName(metricName string) string {
	return fmt.Sprintf("cloud_provider_anx_%s", metricName)
}
