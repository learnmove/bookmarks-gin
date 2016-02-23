package main

import (
	"reflect"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func validateRequest(c *gin.Context) {
	var err error
	// Get token and user from request header
	t := c.Request.Header.Get("X-Access-Token")
	user := c.Request.Header.Get("X-Key")

	// If empty return error
	if t == "" || user == "" {
		c.AbortWithStatus(401)
		c.JSON(401, "Invalid Token or Key")
		return
	}

	// Decode the token using the original key
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	// Token is not valid or there was error when decoding
	if err != nil || !token.Valid {
		c.AbortWithStatus(401)
		c.JSON(401, "Invalid Token or Key")
		return
	}

	// Token seems valid, but there was unexpected error
	if err != nil && token.Valid {
		c.AbortWithStatus(500)
		c.JSON(500, "Oops something went wrong")
		return
	}

	// Validate user and return error if user not found
	_, err = checkUser(user)

	// if user does not exist, return JSON error
	if err != nil {
		c.AbortWithStatus(401)
		c.JSON(401, "Invalid User")
		return
	}

	// Check if token is expired
	expireIn := int64(reflect.ValueOf(token.Claims["exp"]).Float())

	if expired(expireIn) {
		c.AbortWithStatus(400)
		c.JSON(400, "Token Expired")
		return
	}

	c.Next()
}

func validateAdmin(c *gin.Context) {
	dbUser, _ := checkUser(c.Request.Header.Get("X-Key"))

	if dbUser.Role != "admin" {
		c.AbortWithStatus(403)
		c.JSON(403, "Not Authorized")
		return
	}
}
