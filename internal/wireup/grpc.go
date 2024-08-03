package wireup

import (
	"fmt"
	"net"

	"github.com/johnfercher/medium-api/internal/adapters/drivers/grpc"
	"github.com/johnfercher/medium-api/internal/core/ports"
	googleGrpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RunGRPC(productService ports.ProductService) {
	fmt.Println("Init GRPC server")
	server := googleGrpc.NewServer()

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

	fmt.Printf("grpc %s...\n", addr)

	if err := server.Serve(listen); err != nil {
		panic(err)
	}
}
