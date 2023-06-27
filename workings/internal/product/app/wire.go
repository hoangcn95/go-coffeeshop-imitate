//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"google.golang.org/grpc"

	"github.com/hoangcn95/go-coffeeshop-imitate/workings/cmd/product/config"
	"github.com/hoangcn95/go-coffeeshop-imitate/workings/internal/product/app/router"
	"github.com/hoangcn95/go-coffeeshop-imitate/workings/internal/product/infras/inmem"
	"github.com/hoangcn95/go-coffeeshop-imitate/workings/internal/product/service/products"
)

func InitApp(
	cfg *config.Config,
	grpcServer *grpc.Server,
) (*App, error) {
	panic(wire.Build(
		New,
		router.ProductGRPCServerSet,
		inmem.RepositorySet,
		products.ProductServiceSet,
	))
}
