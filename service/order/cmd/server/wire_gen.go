// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"order/internal/biz"
	"order/internal/conf"
	"order/internal/data"
	"order/internal/server"
	"order/internal/service"
)

// Injectors from wire.go:

// initApp init order application.
func initApp(confServer *conf.Server, registry *conf.Registry, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewDB(confData)
	client := data.NewRedis(confData)
	dataData, cleanup, err := data.NewData(confData, logger, db, client)
	if err != nil {
		return nil, nil, err
	}
	orderRepo := data.NewOrderRepo(dataData, logger)
	orderUsecase := biz.NewOrderUsecase(orderRepo, logger)
	orderService := service.NewOrderService(orderUsecase, logger)
	grpcServer := server.NewGRPCServer(confServer, orderService, logger)
	registrar := server.NewRegistrar(registry)
	app := newApp(logger, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
