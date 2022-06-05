package main

//go:build wireinject
// +build wireinject



import (
	"github.com/google/wire"
	"github.com/shunshizhen/golang-learning-camp-homework/homework-for-golang-engineering-practice/internal/conf"
	"github.com/shunshizhen/golang-learning-camp-homework/homework-for-golang-engineering-practice/internal/db"
)

func InitApp() (*Con, func(), error) {
	panic(wire.Build(conf.Provider, db.Provider, NewApp))
}
