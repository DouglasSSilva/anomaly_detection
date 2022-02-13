package serverhandler

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"anomaly_detection/db"

	"github.com/julienschmidt/httprouter"
)

//RequestState struct
type RequestState struct {
	Status int
	W      http.ResponseWriter
	R      *http.Request
	Ps     httprouter.Params
	Rerr   *ResponseError
	Resp   ResponseType
}

func StartApp(env, alternateAccess, alternativeDBConfPath string) *Router {

	goPath := os.Getenv("GOPATH")

	var alternateConf string

	if env == "production" || env == "alpha" {
		alternateConf = fmt.Sprintf("%s%s", goPath, alternateAccess)
	} else {
		alternateConf = fmt.Sprintf("%s%s", goPath, alternativeDBConfPath)
	}

	db.ConnectDB(env, db.DBConfPath, alternateConf)
	defer db.Conn.Close()

	//Create new router to handle external requests
	router := NewRouter()
	router.HandleOPTIONS = true
	return router
}

//GETByIntIDHandler receives id and gets header
func GETByIntIDHandler(modelGet interface{}) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		resp := ResponseType{}

		originHeaderCORS(w, r)
		outfmt, err := checkOutputFormat(r)
		if err != nil {
			w.WriteHeader(http.StatusNotAcceptable)
			resp.AddResult(err)
			return
		}

		id, _ := strconv.ParseInt(ps.ByName("id"), 10, 64)

		modelGetFn := (modelGet).(func(int64) interface{})
		result := modelGetFn(id)
		if err != nil {
			resp.AddResult(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		resp.AddResult(result)
		Response(http.StatusOK, resp, outfmt, w)
	}
}

//BuildOptionsRequest creates option request
func BuildOptionsRequest(router *Router) {
	for path := range router.URLAllowMethods {
		router.Options(path, func(s *RequestState) {

			originHeaderCORS(s.W, s.R)

			s.W.Header().Set("Path", path)
			s.W.Header().Set("Access-Control-Allow-Methods",
				"GET, POST, PATCH, PUT, DELETE, OPTIONS")
			s.W.Header().Set("Access-Control-Allow-Headers",
				"Content-Type, Authorization, Accept")
			s.W.Header().Set("Access-Control-Allow-Credentials", "true")
			//w.Header().Set("Access-Control-Allow-Headers",
			//// strings.Join(router.UrlAllowHeaders[path], " "))
			//"Accept, Content-Type, Content-Length, Accept-Encoding, "+
			//"X-CSRF-Token, Authorization")

			s.W.WriteHeader(http.StatusNoContent)
		})
	}
}
