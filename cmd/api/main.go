package main

import (
	"os"
	"sync"

	"github.com/johnfercher/medium-api/internal/wireup"

	"github.com/johnfercher/medium-api/internal/adapters/drivens/mysql"
	"github.com/johnfercher/medium-api/internal/services"
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

	productRepository := mysql.NewRepository(db)
	productService := services.New(productRepository)

	endpointmetrics.Start()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		wireup.RunREST(productService)
	}()

	go func() {
		defer wg.Done()
		wireup.RunGRPC(productService)
	}()
	wg.Wait()
}
