package productService

import (
	"encoding/json"
	"net/http"

	"github.com/brudrago/clean-go/core/dto"
)

func (service service) Fetch(response http.ResponseWriter, request *http.Request) {
	paginationRequest, err := dto.FromValuePaginationRequestParams(request)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	products, err := service.useCase.Fetch(paginationRequest)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(products)
}
