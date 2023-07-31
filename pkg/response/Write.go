package response

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

func Write(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	var err error
	isXMLRequired := r.Header.Get("Content-Type") == "application/xml"
	if isXMLRequired {
		w.Header().Add("Content-Type", "application/xml")
	} else {
		w.Header().Add("Content-Type", "application/json")
	}
	w.WriteHeader(code)
	if isXMLRequired {
		err = xml.NewEncoder(w).Encode(data)
	} else {
		err = json.NewEncoder(w).Encode(data)
	}
	if err != nil {
		panic(err)
	}
}
