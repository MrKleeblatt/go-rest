package json_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type handleWrapper[P, Q, B, H any, R Response] struct {
	*operation
	handleCallback[P, Q, B, H, R]
}

func (h *handleWrapper[P, Q, B, H, R]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pathvaluesStruct := new(P)
	queryvaluesStruct := new(Q)
	bodyStruct := new(B)
	headersStruct := new(H)
	// TODO: check if argument has bodyBytes, query params, path params
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(bodyReadErr.Code)
		fmt.Fprintln(w, bodyReadErr.String())
		return
	}
	// TODO: check header for content type and parse forms differently than json
	if len(bodyBytes) > 0 {
		err = json.Unmarshal(bodyBytes, &bodyStruct)
		if err != nil {
			w.WriteHeader(jsonExpected.Code)
			fmt.Fprintln(w, jsonExpected.String())
			return
		}
	}
	resultStruct := h.handleCallback(*pathvaluesStruct, *queryvaluesStruct, *bodyStruct, *headersStruct)
	resultBytes, err := json.Marshal(resultStruct)
	if err != nil {
		// TODO: may never happen -> fatal error
		return
	}
	w.WriteHeader(resultStruct.Status().Code)
	w.Write(resultBytes)
}
