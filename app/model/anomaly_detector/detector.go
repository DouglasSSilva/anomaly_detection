package anomaly_detector

type AnomalyDetector struct {
	NormalDistribution      NormalDistribution
	SiteMetrices            []SiteMetrices
	OutlierMultiplier       float64 `json:"OutlierMultiplier"`
	StrongOutlierMultiplier float64 `json:"StrongOutlierMultiplier"`
}
