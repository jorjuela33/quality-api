package mssql

import (
	"github.com/jinzhu/gorm"
	"github.com/jorjuela33/quality-api/database"
)

// MSSQLDB
type MSSQLDB struct {
	GormDB  *gorm.DB
	options *database.Options
}

func (database *MSSQLDB) NewSession() *MSSQLDB {
	options := database.options
	db, err := gorm.Open("mssql", "server="+options.ServerName+";database="+options.DatabaseName+";user id=sa;password=Qu4l1ty;port=1433")
	if err != nil {
		panic(err)
	}
	database.GormDB = db
	return database
}

func New(options *database.Options) *MSSQLDB {
	database := &MSSQLDB{}
	database.options = options
	return database
}
