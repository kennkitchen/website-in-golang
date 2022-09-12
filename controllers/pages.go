package controllers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func DisplayPage(c *gin.Context) {
	file, err := os.OpenFile("logs/webserver.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.Info("Page Controller logging activated.")

	switch c.Request.URL.Path {
	case "/":
		log.Info("Home chosen.")
		c.HTML(http.StatusOK, "home.tmpl", gin.H{
			"title": "Kenn Kitchen",
		})
	case "/about":
		log.Info("About chosen.")
		c.HTML(http.StatusOK, "about.tmpl", gin.H{
			"title": "Kenn Kitchen",
		})
	case "/contact":
		log.Info("Contact chosen.")
		c.HTML(http.StatusOK, "contact.tmpl", gin.H{
			"title": "Kenn Kitchen",
		})
	}

}
