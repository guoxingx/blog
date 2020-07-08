// Package mysql provides easy ways to init mysql client of github.com/go-sql-driver/mysql
package mysql

import (
	"database/sql"
	"fmt"

	// call init function of github.com/go-sql-driver/mysql
	"backend/utils"

	// init must be called otherwise an error will occured
	_ "github.com/go-sql-driver/mysql"
)

// Config for mysql
type Config struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

// New returns an instance of sql.DB
func New(config Config) (*sql.DB, error) {
	connBase := "%s:%s@tcp(%s)/%s"
	connStr := fmt.Sprintf(connBase, config.User, config.Password, config.Host, config.Database)
	connStr = connStr + "?charset=utf8&loc=Asia%2FShanghai&parseTime=true"
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}
	return db, err
}

// NewWithConfigPath returns an instance of sql.DB with the given path
func NewWithConfigPath(configPath string) (*sql.DB, error) {
	var config Config
	err := utils.ReadConfig(configPath, &config)
	if err != nil {
		return nil, err
	}
	return New(config)
}
