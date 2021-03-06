// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/shunshizhen/golang-learning-camp-homework/homework-for-golang-engineering-practice/internal/conf"
	"github.com/shunshizhen/golang-learning-camp-homework/homework-for-golang-engineering-practice/internal/db"
)

// Injectors from wire.go:

func InitApp() (*Con, func(), error) {
	config, err := conf.New()
	if err != nil {
		return nil, nil, err
	}
	sqlDB, err := db.NewDb(config)
	if err != nil {
		return nil, nil, err
	}
	con := NewApp(sqlDB)
	return con, func() {
	}, nil
}
