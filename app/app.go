package app

import (
	"anomaly_detection/app/controller/anomaly_detector"
	"anomaly_detection/serverhandler"
)

// URLRoutes all routes the system use internally
func URLRoutes(r *serverhandler.Router) {

	r.GET("/", Home)
	r.GET("/health", GCPHealthCheck)
	r.Get("/anomaly_detection", anomaly_detector.FindAnomallies)

}
