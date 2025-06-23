package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"kennkitchen/models"
	"kennkitchen/utils"
	"net/http"
	"os"
)

func AddUser(c *gin.Context) {
	file, err := os.OpenFile("logs/webserver.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.Info("User Controller logging activated.")

	username := c.PostForm("username")
	email := username + "@gmail.com"
	password := c.PostForm("password")

	log.Info("Username: ", username)
	log.Info("Password: ", password)

	if len(username) < 8 || len(password) < 8 {
		err := http.StatusNotAcceptable
		log.Warning(http.StatusText(err), err)
		return
	}

	// TODO this should be checking the database to see if username already exists
	// TODO (maybe much later) use HTMX on the front-end to check username before submission
	//if _, ok := users[username]; ok {
	//	err := http.StatusConflict
	//	log.Warning(http.StatusText(err), err)
	//	return
	//}

	// TODO add a password confirmation field
	hashedPassword, _ := utils.HashPassword(password)
	//users[username] = Login{
	//	HashedPassword: hashedPassword,
	//}
	log.Info("Hashed Password: ", hashedPassword)

	// TODO insert new user into database
	_, err = models.InsertUserRow(email, password)
	if err != nil {
		log.Error(err)
	}

	// TODO (maybe much later) do a confirmation email

	c.HTML(http.StatusOK, "regresult.tmpl", gin.H{
		"title": "Kenn Kitchen",
	})

}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateToken(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		log.Fatal(err)
	}
	return base64.URLEncoding.EncodeToString(bytes)
}
