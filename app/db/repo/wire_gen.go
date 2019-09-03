// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package repo

import (
	"github.com/aristat/golang-example-app/app/config"
	"github.com/aristat/golang-example-app/app/db"
	"github.com/aristat/golang-example-app/app/entrypoint"
	"github.com/aristat/golang-example-app/app/logger"
)

// Injectors from injector.go:

func Build() (*Repo, func(), error) {
	context, cleanup, err := entrypoint.ContextProvider()
	if err != nil {
		return nil, nil, err
	}
	viper, cleanup2, err := config.Provider()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	loggerConfig, cleanup3, err := logger.ProviderCfg(viper)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	zap, cleanup4, err := logger.Provider(context, loggerConfig)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	dbConfig, cleanup5, err := db.Cfg(viper)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	gormDB, cleanup6, err := db.ProviderGORM(context, zap, dbConfig)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	usersRepo, cleanup7, err := NewUsersRepo(gormDB)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	repo, cleanup8, err := Provider(usersRepo)
	if err != nil {
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return repo, func() {
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

func BuildTest() (*Repo, func(), error) {
	gormDB, cleanup, err := db.ProviderGORMTest()
	if err != nil {
		return nil, nil, err
	}
	usersRepo, cleanup2, err := NewUsersRepo(gormDB)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	repo, cleanup3, err := Provider(usersRepo)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return repo, func() {
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}