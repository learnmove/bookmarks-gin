package main

import (
	"fmt"
	"time"

	"github.com/codegangsta/cli"
	"golang.org/x/crypto/bcrypt"
)

// CreateDBAdmin - Create administrator account
func CreateDBAdmin(clc *cli.Context) {
	var (
		user     User
		err      error
		passHash []byte
	)

	user.CreatedAt = time.Now().UTC()
	user.UpdatedAt = time.Now().UTC()

	originalPass := []byte(clc.String("password"))
	if passHash, err = bcrypt.GenerateFromPassword(originalPass, bcrypt.DefaultCost); err != nil {
		fmt.Printf("%v\n", err)
	}
	user.PasswordHash = string(passHash)
	user.Username = clc.String("username")
	user.FirstName = clc.String("first")
	user.LastName = clc.String("last")
	user.Email = clc.String("email")
	user.Role = "admin"

	if db.NewRecord(user) {
		if err = db.Create(&user).Error; err != nil {
			fmt.Printf("%v\n", err)
		} else {
			fmt.Printf("Admin user with username: '%s' has been created successfuly!\n", clc.String("username"))
		}
	} else {
		fmt.Printf("Admin user with username: '%s' already exist!\n", clc.String("username"))
	}
}
