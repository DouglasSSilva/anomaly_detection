package main

import (
	"anomaly_detection/app"
	"anomaly_detection/serverhandler"
	"fmt"
	"log"
	"net/http"
	"os"

	"bitbucket.org/residuall-desenv/bastion/db"
)

func main() {

	// Connecting to db, uses a main source of dbconf and a
	// alternate source as a fallback,
	//ENV is defined on app.yaml if not defined will fallback to development
	env := os.Getenv("ENV")
	goPath := os.Getenv("GOPATH")
	alternateConf := fmt.Sprintf("%s%s", goPath, db.AlternativeDBConfPath)

	db.ConnectDB(env, db.DBConfPath, alternateConf)
	defer db.Conn.Close()

	//Create new router to handle external requests
	router := serverhandler.NewRouter()
	router.HandleOPTIONS = true
	app.URLRoutes(router)
	serverhandler.BuildOptionsRequest(router)

	//based on environment sets host to run API
	host := ":5000"

	if env == "production" {
		host = fmt.Sprintf(":%s", os.Getenv("PORT_SYSTEM"))

	}

	fmt.Println("-anomaly detection server started. Listening on" + host)
	log.Fatal(http.ListenAndServe(host, router))
}
