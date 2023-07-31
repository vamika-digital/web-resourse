package request

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

func Read(r *http.Request, requestObj *interface{}) error {
	isXMLRequired := r.Header.Get("Content-Type") == "application/xml"
	if isXMLRequired {
		return xml.NewDecoder(r.Body).Decode(&requestObj)
	} else {
		return json.NewDecoder(r.Body).Decode(&requestObj)
	}
}
