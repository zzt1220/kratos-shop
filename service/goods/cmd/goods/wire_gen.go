// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"goods/internal/biz"
	"goods/internal/conf"
	"goods/internal/data"
	"goods/internal/server"
	"goods/internal/service"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewDB(confData)
	client := data.NewRedis(confData)
	dataData, cleanup, err := data.NewData(confData, logger, db, client)
	if err != nil {
		return nil, nil, err
	}
	goodsRepo := data.NewGoodsRepo(dataData, logger)
	goodsUsecase := biz.NewGoodsUsecase(goodsRepo, logger)
	categoryRepo := data.NewCategoryRepo(dataData, logger)
	categoryUsecase := biz.NewCategoryUsecase(categoryRepo, logger)
	brandRepo := data.NewBrandRepo(dataData, logger)
	brandUsecase := biz.NewBrandUsecase(brandRepo, logger)
	goodsService := service.NewGoodsService(goodsUsecase, categoryUsecase, brandUsecase, logger)
	grpcServer := server.NewGRPCServer(confServer, goodsService, logger)
	app := newApp(logger, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
