package app

import (
	"github.com/hoangcn95/go-coffeeshop-imitate/workings/cmd/product/config"
	"github.com/hoangcn95/go-coffeeshop-imitate/workings/internal/product/service/products"
	"github.com/hoangcn95/go-coffeeshop-imitate/workings/proto/gen"
)

type App struct {
	Cfg               *config.Config
	ProductService    products.ProductService
	ProductGRPCServer gen.ProductServiceServer
}

func New(
	cfg *config.Config,
	productService products.ProductService,
	productGRPCServer gen.ProductServiceServer,
) *App {
	return &App{
		Cfg:               cfg,
		ProductService:    productService,
		ProductGRPCServer: productGRPCServer,
	}
}
