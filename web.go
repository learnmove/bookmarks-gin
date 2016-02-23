package main

// IMPORTS
import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// db connection object
var db gorm.DB

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
	// create any new tables
	db.AutoMigrate(&Bookmark{})

	// initialize gin
	r := gin.Default()

	// get current folder
	pwd, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	// Serve AngularJS
	r.Use(static.Serve("/", static.LocalFile(fmt.Sprintf("%s/public", pwd), true)))

	// v1 API routes
	v1 := r.Group("/api/v1")
	{
		v1.GET("/bookmarks", getBookmarks)
		v1.GET("/bookmarks/:id", getBookmark)
		v1.POST("/bookmarks", addBookmark)
		v1.PUT("/bookmarks/:id", updateBookmark)
		v1.DELETE("/bookmarks/:id", deleteBookmark)
	}

	// start up server
	r.Run(os.Getenv("IP") + ":" + os.Getenv("PORT"))
}