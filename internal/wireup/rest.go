package wireup

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/johnfercher/medium-api/internal/adapters/drivers/rest"
	"github.com/johnfercher/medium-api/internal/core/ports"
	"github.com/johnfercher/medium-api/pkg/api"
	"github.com/johnfercher/medium-api/pkg/chaos"
)

// nolint:gomnd // magic number
func RunREST(productService ports.ProductService) {
	fmt.Println("Init REST server")
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	handlers := []api.HTTPHandler{}

	getProductByIDHandler := rest.NewGetProductByID(productService)
	handlers = append(handlers, getProductByIDHandler)

	searchProductsHandler := rest.NewSearchProducts(productService)
	handlers = append(handlers, searchProductsHandler)

	createProductHandler := rest.NewCreateProduct(productService)
	handlers = append(handlers, createProductHandler)

	updateProductHandler := rest.NewUpdateProduct(productService)
	handlers = append(handlers, updateProductHandler)

	deleteProductHandler := rest.NewDeleteProduct(productService)
	handlers = append(handlers, deleteProductHandler)

	for i, handler := range handlers {
		// Only to inject errors and delay
		chaosHTTPAdapter := chaos.NewChaosHTTPHandler(handler, float64(i*150))
		metricsAdapter := api.NewMetricsHandlerAdapter(chaosHTTPAdapter)

		r.MethodFunc(handler.Verb(), handler.Pattern(), metricsAdapter.AdaptHandler())
	}

	fmt.Printf("rest 0.0.0.0:8081...\n")

	if err := http.ListenAndServe(":8081", r); err == nil {
		panic(err)
	}
}
