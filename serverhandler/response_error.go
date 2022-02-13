package serverhandler

import (
	"encoding/json"
	"errors"
	"strings"
)

// ResultError set to fron
type ResultError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Field   string `json:"field"`
}

//ResponseError array of result error
type ResponseError struct {
	Results []ResultError
}

//Error marshaling and identing output
func (rerr *ResponseError) Error() string {
	res, err := json.MarshalIndent(rerr.Results, "", "\t")
	if err != nil {
		//não testável
		panic(err)
	}
	return string(res)
}

//AddError adds error to return struct
func (rerr *ResponseError) AddError(typ, mess, field string) {
	rerr.Results = append(rerr.Results, ResultError{typ, mess, field})
}

// NewErrorFromString to return
func NewErrorFromString(s string) error {
	if strings.ToLower(s) == "null" || s == "" {
		return nil
	}
	return errors.New(s)

}
