package serverhandler

import (
	"net/http"
)

func originHeaderCORS(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	}
}

func checkOutputFormat(r *http.Request) (string, error) {
	var mimetype string
	outfmt := r.URL.Query().Get("fmt")
	// guaranties that the default value will be json
	if outfmt == "" || outfmt == "json" {
		outfmt = "json"
		mimetype = "application/json"
	} else if outfmt == "xml" {
		mimetype = "application/xml"
	} else {
		return "text", &ResponseError{
			[]ResultError{
				{"ERR_OUTPUT_FORMAT",
					`Wrong output format given. Options are 'json'
					(default if empty) and 'xml'.`, ""},
			},
		}
	}

	if r.Header.Get("Accept") != mimetype {
		outfmt = "text"
		return outfmt, &ResponseError{
			[]ResultError{
				{"not_acceptable",
					"Accept header must be set to '" + mimetype + "'.",
					""},
			},
		}
	}
	return outfmt, nil
}
