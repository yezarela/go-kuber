package functions

import (
	"context"
	"encoding/json"
	productUsecase "kuba/product/usecase"
	"kuba/utils/tern"
	"net/http"
	"strconv"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

var usecase productUsecase.Usecase

func init() {
	functions.HTTP("ListProduct", ListProduct)

	// DB, config initialization goes here
}

// ListProduct is an HTTP Cloud Function with a request parameter.
func ListProduct(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	limit, _ := strconv.Atoi(query.Get("limit"))
	offset, _ := strconv.Atoi(query.Get("offset"))

	limit = tern.Int(limit, 10)

	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()
	products, err := usecase.ListProduct(ctx, limit, offset)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	json.NewEncoder(w).Encode(products)
}
