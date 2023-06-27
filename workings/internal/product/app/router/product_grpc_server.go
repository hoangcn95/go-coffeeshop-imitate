package router

import (
	"context"

	"github.com/google/wire"
	"github.com/hoangcn95/go-coffeeshop-imitate/workings/internal/product/service/products"
	"github.com/hoangcn95/go-coffeeshop-imitate/workings/proto/gen"
	"github.com/pkg/errors"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type productGRPCServer struct {
	gen.UnimplementedProductServiceServer
	productService products.ProductService
}

var _ gen.ProductServiceServer = (*productGRPCServer)(nil)

// definition custom dependency injection
var ProductGRPCServerSet = wire.NewSet(NewProductGRPCServer)

func NewProductGRPCServer(
	grpcServer *grpc.Server,
	productService products.ProductService,
) gen.ProductServiceServer {
	svc := productGRPCServer{
		productService: productService,
	}

	gen.RegisterProductServiceServer(grpcServer, &svc)

	reflection.Register(grpcServer)

	return &svc
}

func (g *productGRPCServer) GetItemTypes(
	ctx context.Context,
	request *gen.GetItemTypesRequest,
) (*gen.GetItemTypesResponse, error) {
	slog.Info("gRPC client", "http_method", "GET", "http_name", "GetItemTypes")

	res := gen.GetItemTypesResponse{}

	results, err := g.productService.GetItemTypes(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "productGRPCServer-GetItemTypes")
	}

	for _, item := range results {
		res.ItemTypes = append(res.ItemTypes, &gen.ItemTypeDto{
			Name:  item.Name,
			Type:  int32(item.Type),
			Price: item.Price,
			Image: item.Image,
		})
	}

	return &res, nil
}

func (g *productGRPCServer) GetItemsByType(
	ctx context.Context,
	request *gen.GetItemsByTypeRequest,
) (*gen.GetItemsByTypeResponse, error) {
	slog.Info("gRPC client", "http_method", "GET", "http_name", "GetItemsByType", "item_types", request.ItemTypes)

	res := gen.GetItemsByTypeResponse{}

	results, err := g.productService.GetItemsByType(ctx, request.ItemTypes)
	if err != nil {
		return nil, errors.Wrap(err, "productGRPCServer-GetItemsByType")
	}

	for _, item := range results {
		res.Items = append(res.Items, &gen.ItemDto{
			Type:  int32(item.Type),
			Price: item.Price,
		})
	}

	return &res, nil
}
