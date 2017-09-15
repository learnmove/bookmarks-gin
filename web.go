package main

// IMPORTS
import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spf13/viper"
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
	// Read configuration file
	viper.SetConfigName("app")
	viper.AddConfigPath("./")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("Config file not found...")
		panic(err)
	}

	// connect to database
	db, err = gorm.Open(viper.GetString("database.engine"), getSQLString())
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//initialize it
	db.DB()

	// Configure CLI
	app := cli.NewApp()
	app.Name = "Bookmarks-Gin"
	app.Usage = "Server in Gin for Bookmarks"
	app.Author = "Stefan Jarina"
	app.Email = "stefan@jarina.cz"
	app.Version = "0.0.3"
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

		r.Run(getAddrIP(clc))
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
