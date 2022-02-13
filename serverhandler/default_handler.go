package serverhandler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//DefaultHandler handles any request made
func DefaultHandler(next func(*RequestState)) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Internal server error. Something went wrong.")

			}
		}()

		s := &RequestState{
			Status: http.StatusOK,
			W:      w,
			R:      r,
			Ps:     ps,
			Resp:   NewResponseType(),
		}

		originHeaderCORS(w, r)
		outfmt, err := checkOutputFormat(r)
		if rerr, ok := err.(*ResponseError); ok {
			for _, v := range rerr.Results {
				s.Resp.AddResult(v)
			}
			w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
			w.WriteHeader(http.StatusNotAcceptable)
			fmt.Fprint(w, `Error: 'Accept' header must be set to 
				'application/json' or 'application/xml'.`)
			return
		}

		next(s)

		if s.Rerr != nil && len(s.Rerr.Results) > 0 {
			s.Resp.Status = "error"
			for _, v := range s.Rerr.Results {
				s.Resp.AddResult(v)
			}

			Response(s.Status, s.Resp, outfmt, s.W)
			return
		}

		Response(s.Status, s.Resp, outfmt, s.W)
	}
}
