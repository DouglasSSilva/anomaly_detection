package anomaly_detector

import "math"

type NormalDistribution struct {
	Mean              float64
	Variance          float64
	StandardDeviation float64
}

func (normalDistribution *NormalDistribution) setMean(siteMetrices []SiteMetrices) {
	somatoriun := float64(0)
	siteMetricesLen := float64(len(siteMetrices))

	for i := range siteMetrices {
		somatoriun += siteMetrices[i].Value
	}

	normalDistribution.Mean = somatoriun / siteMetricesLen
}

func (normalDistribution *NormalDistribution) setVariance(siteMetrices []SiteMetrices) {
	somatoriun := float64(0)
	siteMetricesLen := float64(len(siteMetrices))

	for i := range siteMetrices {
		differenceFromTheMean := siteMetrices[i].Value - normalDistribution.Mean
		somatoriun += differenceFromTheMean * differenceFromTheMean
	}

	normalDistribution.Mean = somatoriun / (siteMetricesLen - 1)

}

func (normalDistribution *NormalDistribution) setStandardDeviation() {
	normalDistribution.StandardDeviation = math.Sqrt(normalDistribution.Variance)

}

func (normalDistribution *NormalDistribution) Create(siteMetrices []SiteMetrices) {
	normalDistribution.setMean(siteMetrices)
	normalDistribution.setVariance(siteMetrices)
	normalDistribution.setStandardDeviation()
}
