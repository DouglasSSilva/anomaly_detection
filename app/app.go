package app

import "anomaly_detection/serverhandler"

// URLRoutes all routes the system use internally
func URLRoutes(r *serverhandler.Router) {

	r.GET("/", Home)
	r.GET("/health", GCPHealthCheck)

}
