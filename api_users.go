package main

// IMPORTS
import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/crypto/bcrypt"
)

func getUsers(c *gin.Context) {
	users := []User{}
	db.Find(&users)
	c.JSON(200, users)
}

func getUser(c *gin.Context) {
	var user User
	id := c.Params.ByName("id")
	db.First(&user, id)
	c.JSON(200, user)
}

func addUser(c *gin.Context) {
	var (
		user User
		err  error
		passHash     []byte
	)
	err = c.BindWith(&user, binding.JSON)
	if err != nil {
		c.JSON(404, "Couldn't bind!")
	} else {
		if db.NewRecord(user) {
			user.CreatedAt = time.Now().UTC()
			user.UpdatedAt = time.Now().UTC()

			// Create password hash
			originalPass := []byte(user.PasswordHash)
			if passHash, err = bcrypt.GenerateFromPassword(originalPass, bcrypt.DefaultCost); err != nil {
				fmt.Printf("%v", err)
			}
			user.PasswordHash = string(passHash)

			db.Create(&user)
			c.JSON(201, user)
		} else {
			c.JSON(400, "This user already exist")
		}
	}
}

func updateUser(c *gin.Context) {
	var (
		user      User
		changeset User
		err error
		passHash     []byte
	)
	id := c.Params.ByName("id")
	db.First(&user, id)
	err = c.BindWith(&changeset, binding.JSON)
	if err != nil {
		c.JSON(404, "Couldn't bind!")
	} else {
		user.UpdatedAt = time.Now().UTC()
		user.Username = changeset.Username
		user.Email = changeset.Email
		if changeset.PasswordHash != "" && changeset.PasswordHash != user.PasswordHash {
			err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(changeset.PasswordHash))
			if err != nil {
				if passHash, err = bcrypt.GenerateFromPassword([]byte(changeset.PasswordHash), bcrypt.DefaultCost); err != nil {
					fmt.Printf("%v", err)
				}
				user.PasswordHash = string(passHash)
			}
		}
		db.Save(&user)
		c.JSON(201, user)
	}
}

func deleteUser(c *gin.Context) {
	var user User
	id := c.Params.ByName("id")
	db.First(&user, id)
	db.Delete(&user)
}
