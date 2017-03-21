package database

type Options struct {
	ServerName   string
	DatabaseName string
}

// Database interface
type DatabaseInterface interface {
	NewSession(database *DatabaseInterface) *DatabaseInterface
}
