package anomaly_detector

import "math"

type AnomalyDetector struct {
	NormalDistribution      NormalDistribution
	SiteMetrices            []SiteMetrices
	OutlierMultiplier       float64
	StrongOutlierMultiplier float64
}

func (anomalyDetector *AnomalyDetector) setMean() {
	somatoriun := float64(0)
	siteMetricesLen := float64(len(anomalyDetector.SiteMetrices))

	for i := range anomalyDetector.SiteMetrices {
		somatoriun += anomalyDetector.SiteMetrices[i].Value
	}

	anomalyDetector.NormalDistribution.Mean = somatoriun / siteMetricesLen
}

func (anomalyDetector *AnomalyDetector) setVariance() {
	somatoriun := float64(0)
	siteMetricesLen := float64(len(anomalyDetector.SiteMetrices))

	for i := range anomalyDetector.SiteMetrices {
		differenceFromTheMean := anomalyDetector.SiteMetrices[i].Value - anomalyDetector.NormalDistribution.Mean
		somatoriun += differenceFromTheMean * differenceFromTheMean
	}

	anomalyDetector.NormalDistribution.Mean = somatoriun / (siteMetricesLen - 1)

}

func (anomalyDetector *AnomalyDetector) setStandardDeviation() {
	anomalyDetector.NormalDistribution.StandardDeviation = math.Sqrt(anomalyDetector.NormalDistribution.Variance)

}

func (anomalyDetector *AnomalyDetector) CreateNormalDistrubution() {
	anomalyDetector.setMean()
	anomalyDetector.setVariance()
}
