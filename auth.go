package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Load private key and store it in global variable as []byte
var key, _ = ioutil.ReadFile(KEY)

// DbUserObject - struct for user object
type DbUserObject struct {
	Name     string `json:"name"`
	Role     string `json:"role"`
	Username string `json:"username"`
}

// Token - struct for token
type Token struct {
	Token   string       `json:"token"`
	Expires int64        `json:"expires"`
	User    DbUserObject `json:"user"`
}

// LoginData - struct for authentication
type LoginData struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func loginHandler(c *gin.Context) {
	var (
		login LoginData
	)
	// bind POST data to struct
	err := c.Bind(&login)
	if err != nil {
		c.JSON(401, "Invalid Credentials Provided")
	} else {
		// return 401 if empty
		if login.Username == "" || login.Password == "" {
			c.JSON(401, "Invalid Credentials Provided")
			return
		}

		// get user from database and fill our struct
		dbUserObject, err := validateUser(login.Username, login.Password)
		if err != nil {
			// return 401 if incorrect user or password
			c.JSON(401, "Invalid Credentials")
			return
		}
		// generate token
		token := genToken(dbUserObject)

		// return token to user
		c.JSON(200, token)
		return
	}

	// return 400 if any other error is encountered
	c.JSON(400, "Error encountered")
}

func validateUser(username string, password string) (DbUserObject, error) {
	var (
		dbUserObj DbUserObject
		user      User
		err       error
	)

	if err = db.Where(&User{Username: username}).First(&user).Error; err != nil {
		fmt.Printf("Received error when retrieving user '%s': %v\n", username, err)
		return DbUserObject{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return DbUserObject{}, err
	}

	dbUserObj.Name = user.FirstName
	dbUserObj.Role = user.Role
	dbUserObj.Username = user.Username

	return dbUserObj, nil
}

func checkUser(username string) (DbUserObject, error) {
	var (
		dbUserObj DbUserObject
		user      User
		err       error
	)

	if err = db.Where(&User{Username: username}).First(&user).Error; err != nil {
		fmt.Printf("Received error when retrieving user '%s': %v\n", username, err)
		return DbUserObject{}, err
	}

	dbUserObj.Name = user.FirstName
	dbUserObj.Role = user.Role
	dbUserObj.Username = user.Username

	return dbUserObj, nil
}

func genToken(user DbUserObject) Token {
	expires := expiresIn(7)
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = expires
	tokenString, err := token.SignedString(key)
	if err != nil {
		return Token{}
	}
	response := Token{tokenString, expires, user}
	return response
}

func expiresIn(numDays int) int64 {
	time := time.Now().Add(time.Hour * time.Duration((numDays * 24))).Unix()
	return time
}
