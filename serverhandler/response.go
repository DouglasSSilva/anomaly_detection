package serverhandler

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type ResponseType struct {
	Status  string        `json:"status"`
	Results []interface{} `json:"results"`
}

//AddResult adds result to response type
func (resp *ResponseType) AddResult(res interface{}) {
	resp.Results = append(resp.Results, res)
}

//NewResponseType creates a response type with ok status
func NewResponseType() ResponseType {
	resp := ResponseType{"ok", []interface{}{}}
	return resp
}

//Response creates response as xml or json to return
func Response(status int, resp ResponseType, outfmt string,
	w http.ResponseWriter) error {

	if outfmt == "json" {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(status)
		return json.NewEncoder(w).Encode(resp)
	} else if outfmt == "xml" {
		w.Header().Set("Content-Type", "application/xml; charset=UTF-8")
		w.WriteHeader(status)
		return xml.NewEncoder(w).Encode(resp)
	}
	return nil
}
