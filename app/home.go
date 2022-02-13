package app

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Home gives back a 200 and let all know we're alive
func Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Residuall Demiurge 1.0 :)")
}

// GCPHealthCheck gives back a 200
func GCPHealthCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Residuall Demiurge 1.0 :)")
}
