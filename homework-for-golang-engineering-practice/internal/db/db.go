package db

import (
	"database/sql"

    "github.com/shunshizhen/golang-learning-camp-homework/homework-for-golang-engineering-practice/internal/conf"

	"github.com/google/wire"
	
)

var Provider = wire.NewSet(NewDb)

func NewDb(conf *conf.Config)(db *sql.DB,err error){
    db,err=sql.Open("mysql",conf.Db.Dsn)
    if err!=nil{
        return
    }
    if err=db.Ping();err!=nil{
        return
    }
    return db,nil
}