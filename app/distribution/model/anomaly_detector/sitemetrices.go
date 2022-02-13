package anomaly_detector

type SiteMetrices struct {
	SiteID   string  `json:"SiteID"`
	Metrices string  `json:"metrices"`
	Date     string  `json:"date"`
	Value    float64 `json:"value"`
}
