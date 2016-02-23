package main

// IMPORTS
import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

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
			bookmark.CreatedAt = time.Now().UTC()
			bookmark.UpdatedAt = time.Now().UTC()
			db.Create(&bookmark)
			c.JSON(201, bookmark)
		} else {
			c.JSON(400, "This bookmark already exist")
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
		bookmark.UpdatedAt = time.Now().UTC()
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
