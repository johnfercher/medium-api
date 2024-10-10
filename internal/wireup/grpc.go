package wireup

import (
	"context"
	"net"

	"github.com/johnfercher/medium-api/pkg/observability/log"
	"github.com/johnfercher/medium-api/pkg/observability/log/field"

	"github.com/johnfercher/medium-api/pkg/api"

	"github.com/johnfercher/medium-api/internal/adapters/drivers/grpc"
	"github.com/johnfercher/medium-api/internal/core/ports"
	googleGrpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RunGRPC(ctx context.Context, productService ports.ProductService) {
	log.Info(ctx, "Init GRPC server")

	loggerIntercepter := api.NewLoggerInterceptor()
	metricsInterceptor := api.NewMetricsInterceptor()

	server := googleGrpc.NewServer(googleGrpc.ChainUnaryInterceptor(loggerIntercepter.Intercept, metricsInterceptor.Intercept))

	addr := "0.0.0.0:8082"

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	createProductHandler := grpc.NewCreateProduct(productService)
	grpc.RegisterCreateProductHandlerServer(server, createProductHandler)

	getProductByIDHandler := grpc.NewGetProductByID(productService)
	grpc.RegisterGetProductHandlerServer(server, getProductByIDHandler)

	deleteProductHandler := grpc.NewDeleteProduct(productService)
	grpc.RegisterDeleteProductHandlerServer(server, deleteProductHandler)

	searchProductsHandler := grpc.NewSearchProduct(productService)
	grpc.RegisterSearchProductHandlerServer(server, searchProductsHandler)

	updateProductHandler := grpc.NewUpdateProduct(productService)
	grpc.RegisterUpdateProductHandlerServer(server, updateProductHandler)

	reflection.Register(server)

	log.Info(ctx, "started grpc", field.String("addr", addr))

	if err := server.Serve(listen); err != nil {
		panic(err)
	}
}
