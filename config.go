package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"
	"github.com/spf13/viper"
)

// KEY - Global variable for key used by token validation
var KEY string

// InitKeyPath - Initialize the location of KEY
func InitKeyPath() {
	pwd, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	KEY = fmt.Sprintf("%s/key.pem", pwd)
}

func getAddrIP(clc *cli.Context) string {
	var address, port string

	if os.Getenv("OPENSHIFT_NODEJS_IP") != "" {
		address = os.Getenv("OPENSHIFT_NODEJS_IP")
	} else if os.Getenv("IP") != "" {
		address = os.Getenv("IP")
	} else if clc.String("ip") != "" {
		address = clc.String("ip")
	} else {
		address = viper.GetString("app.ip")
	}

	if os.Getenv("OPENSHIFT_NODEJS_PORT") != "" {
		port = os.Getenv("OPENSHIFT_NODEJS_PORT")
	} else if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	} else {
		port = viper.GetString("app.port")
	}

	return address + ":" + port
}

func getSQLString() string {
	var engine, host, user, password, sslmode, dbname, filepath, port string
	// read config into variables
	engine = viper.GetString("database.engine")
	host = viper.GetString("database.host")
	port = viper.GetString("database.port")
	user = viper.GetString("database.user")
	password = viper.GetString("database.password")
	dbname = viper.GetString("database.dbname")
	sslmode = viper.GetString("database.sslmode")
	filepath = viper.GetString("database.filepath")

	// build connection string for different databases
	if engine == "mysql" {
		return fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True", user, password, dbname)
	} else if engine == "postgres" {
		return fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s", host, user, dbname, sslmode, password)
	} else if engine == "mssql" {
		return fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", user, password, host, port, dbname)
	}
	// fallback to sqlite3
	return filepath
}
