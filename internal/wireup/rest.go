package wireup

import (
	"context"
	"net/http"

	"github.com/johnfercher/medium-api/pkg/observability/log"
	"github.com/johnfercher/medium-api/pkg/observability/log/field"

	"github.com/go-chi/chi/v5"
	"github.com/johnfercher/medium-api/internal/adapters/drivers/rest"
	"github.com/johnfercher/medium-api/internal/core/ports"
	"github.com/johnfercher/medium-api/pkg/api"
	"github.com/johnfercher/medium-api/pkg/chaos"
)

// nolint:gomnd // magic number
func RunREST(ctx context.Context, productService ports.ProductService) {
	log.Info(ctx, "Init REST server")
	r := chi.NewRouter()

	r.Use(log.ContextMiddleware)

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

	addr := ":8081"

	log.Info(ctx, "started rest", field.String("addr", addr))

	if err := http.ListenAndServe(addr, r); err == nil {
		panic(err)
	}
}
