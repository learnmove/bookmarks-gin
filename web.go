package main

// IMPORTS
import (
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// db connection object
var db *gorm.DB

func init() {
	// Set KEY PATH based on location of binary
	InitKeyPath()
	// Check if key already exist
	if checkKeyFile() {
		// If file already exists, load and use.
		return
	}
	// If file doesn't exist, try to generate new key
	if err := generateKeyFile(); err != nil {
		// If can't generate file, panic the application as it can't operate without it.
		fmt.Println("Problem when creating KEY file")
		panic(err)
	} else {
		fmt.Println("Key file created!.")
		return
	}
}

// main function
func main() {
	var err error
	// connect to database
	db, err = gorm.Open("mysql", os.Getenv("GO_MYSQL_URI"))
	if err != nil {
		panic(err)
	}
	//initialize it
	db.DB()

	app := cli.NewApp()
	app.Name = "Bookmarks-Gin"
	app.Usage = "Server for bookmarks-gin"
	app.Author = "Stefan Jarina"
	app.Email = "stefan@jarina.cz"
	app.Version = "0.0.2"
	app.Action = func(clc *cli.Context) {
		// initialize gin
		r := gin.Default()

		// get current folder
		pwd, _ := filepath.Abs(filepath.Dir(os.Args[0]))

		// Serve AngularJS
		r.Use(static.Serve("/", static.LocalFile(fmt.Sprintf("%s/public", pwd), true)))

		// Handle login
		r.POST("/login", loginHandler)

		// v1 API routes
		v1 := r.Group("/api/v1", validateRequest)
		{
			v1.GET("/bookmarks", getBookmarks)
			v1.GET("/bookmarks/:id", getBookmark)
			v1.POST("/bookmarks", addBookmark)
			v1.PUT("/bookmarks/:id", updateBookmark)
			v1.DELETE("/bookmarks/:id", deleteBookmark)

			// v1 ADMIN API - Requires Authentication and Authorization - Only 'admin' allowed
			admin := v1.Group("admin", validateAdmin)
			{
				admin.GET("/users", getUsers)
				admin.GET("/users/:id", getUser)
				admin.POST("/users", addUser)
				admin.PUT("/users/:id", updateUser)
				admin.DELETE("/users/:id", deleteUser)
			}
		}

		// start up server
		if clc.Bool("socket") {
			l, err := net.Listen("unix", "/tmp/bookmarks-gin.sock")
			if err != nil {
				panic(err)
			}
			http.Serve(l, r)
		} else {
			r.Run(getAddrIP(clc))
		}
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "ip",
			Usage: "Specify IP to which server should bind",
		},
		cli.StringFlag{
			Name:  "port",
			Usage: "Specify PORT application should bind to",
		},
		cli.BoolFlag{
			Name:  "socket",
			Usage: "Specify if SOCKET should be used",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "initdb",
			Usage: "Initialize database with default admin user",
			Action: func(clc *cli.Context) {
				if clc.IsSet("password") {
					// create any new tables
					db.AutoMigrate(&User{}, &Bookmark{})
					CreateDBAdmin(clc)
				} else {
					fmt.Println("No password specified")
				}
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "username",
					Value: "admin",
					Usage: "Set username for user with admin rights",
				},
				cli.StringFlag{
					Name:  "password",
					Usage: "Set password for user with admin rights",
				},
				cli.StringFlag{
					Name:  "first",
					Value: "Administrator",
					Usage: "Set first name for user with admin rights",
				},
				cli.StringFlag{
					Name:  "last",
					Value: "",
					Usage: "Set last name for user with admin rights",
				},
				cli.StringFlag{
					Name:  "email",
					Value: "foo@bar.com",
					Usage: "Set email for user with admin rights",
				},
			},
		},
	}

	app.Run(os.Args)
}
