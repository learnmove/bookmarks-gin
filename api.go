package main

// IMPORTS
import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Bookmark model used by database
type Bookmark struct {
	ID          uint   `json:"id" gorm:"primary_key" sql:"AUTO_INCREMENT"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Url         string `json:"url"`
}

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

	// v1 API
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

func getBookmarks(c *gin.Context) {
	bookmarks := []Bookmark{}
	db.Find(&bookmarks)
	c.JSON(200, bookmarks)
}

func getBookmark(c *gin.Context) {
	var bookmark Bookmark
	id := c.Params.ByName("id")
	db.First(&bookmark, id)
	c.JSON(200, bookmark)
}

func addBookmark(c *gin.Context) {
	var (
		bookmark Bookmark
		err      error
	)
	err = c.BindWith(&bookmark, binding.JSON)
	if err != nil {
		c.JSON(404, "Couldn't bind!")
	} else {
		if db.NewRecord(bookmark) {
			db.Create(&bookmark)
			c.JSON(201, bookmark)
		} else {
			c.JSON(400, "This user already exist")
		}
	}
}

func updateBookmark(c *gin.Context) {
	var (
		bookmark  Bookmark
		changeset Bookmark
	)
	id := c.Params.ByName("id")
	db.First(&bookmark, id)
	err := c.BindWith(&changeset, binding.JSON)
	if err != nil {
		c.JSON(404, "Couldn't bind!")
	} else {
		bookmark.Name = changeset.Name
		bookmark.Description = changeset.Description
		bookmark.Url = changeset.Url
		db.Save(&bookmark)
		c.JSON(201, bookmark)
	}
}

func deleteBookmark(c *gin.Context) {
	var bookmark Bookmark
	id := c.Params.ByName("id")
	db.First(&bookmark, id)
	db.Delete(&bookmark)
}
