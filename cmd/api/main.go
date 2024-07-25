package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/johnfercher/medium-api/internal/adapters/drivens/mysql"
	"github.com/johnfercher/medium-api/internal/adapters/drivers/rest"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/johnfercher/medium-api/internal/services"
	"github.com/johnfercher/medium-api/pkg/api"
	"github.com/johnfercher/medium-api/pkg/chaos"
	"github.com/johnfercher/medium-api/pkg/config"
	"github.com/johnfercher/medium-api/pkg/mysqldriver"
	"github.com/johnfercher/medium-api/pkg/observability/metrics/endpointmetrics"
)

// nolint:gomnd // magic number
func main() {
	cfg, err := config.Load(os.Args)
	if err != nil {
		panic(err)
	}

	db, err := mysqldriver.Start(cfg.Mysql.URL, cfg.Mysql.DB, cfg.Mysql.User, cfg.Mysql.Password)
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	productRepository := mysql.NewRepository(db)
	productService := services.New(productRepository)

	handlers := []api.HTTPHandler{}

	endpointmetrics.Start()

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

	_ = http.ListenAndServe(":8081", r)
}
