package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"
)

// KEY - Global variable for keu used by token validation
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
		address = "0.0.0.0"
	}

	if os.Getenv("OPENSHIFT_NODEJS_PORT") != "" {
		port = os.Getenv("OPENSHIFT_NODEJS_PORT")
	} else if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	} else {
		port = "3000"
	}

	return address + ":" + port
}
