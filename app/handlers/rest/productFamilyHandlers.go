package rest

import (
	"github.com/gorilla/mux"
	"github.com/vamika-digital/wms-resourse/internal/core/ports/driving"
	"github.com/vamika-digital/wms-resourse/pkg/response"
	"net/http"
)

type ProductFamilyHandlers struct {
	service driving.ProductFamilyService
}

func (ch *ProductFamilyHandlers) getAllProductFamilies(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	productFamilies, err := ch.service.GetAllProductFamilies(status)

	if err != nil {
		response.Write(w, r, err.Code, err.AsMessage())
	} else {
		response.Write(w, r, http.StatusOK, productFamilies)
	}
}

func (ch *ProductFamilyHandlers) getProductFamily(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["productFamily_id"]

	productFamily, err := ch.service.GetProductFamily(id)
	if err != nil {
		response.Write(w, r, err.Code, err.AsMessage())
	} else {
		response.Write(w, r, http.StatusOK, productFamily)
	}
}

func (ch *ProductFamilyHandlers) saveProductFamily(w http.ResponseWriter, r *http.Request) {
	//var request2 NewProductFamilyRequest
	//request.Read(, request)

	//productFamily, err := ch.service.GetProductFamily(id)
	//if err != nil {
	//	response.Write(w, r, err.Code, err.AsMessage())
	//} else {
	//	response.Write(w, r, http.StatusOK, productFamily)
	//}
}
