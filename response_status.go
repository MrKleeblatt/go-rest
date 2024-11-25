package json_api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type ResponseStatus struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Details string `json:"-"` // ignore this field in json
}

type Response interface {
	Status() ResponseStatus
}

func (s ResponseStatus) Status() ResponseStatus {
	return s
}

func (r *ResponseStatus) Error() string {
	return errors.New(fmt.Sprintf("%d %s: %s ==> %s",
		r.Code,
		http.StatusText(r.Code),
		r.Message,
		r.Details)).Error()
}

func (r *ResponseStatus) String() string {
	jsonData, err := json.Marshal(r)
	if err != nil {
		// TODO: error logging. This may never happen.
		return ""
	}
	return string(jsonData)
}

var ok = &ResponseStatus{Code: http.StatusOK}
var jsonExpected = &ResponseStatus{Code: http.StatusBadRequest, Message: "json body expected"}
var bodyReadErr = &ResponseStatus{Code: http.StatusInternalServerError, Message: "internal server error", Details: "could not read body"}

func StatusOK() *ResponseStatus { return ok }
