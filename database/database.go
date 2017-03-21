package database

import "github.com/jinzhu/gorm"

type Options struct {
	ServerName   string
	DatabaseName string
}

// Database interface
type DatabaseInterface interface {
	Foo()
	DB() *gorm.DB
}
