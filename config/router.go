package config

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"kennkitchen/controllers"
	"os"
)

func RunServer() {
	file, err := os.OpenFile("logs/webserver.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.Info("Router logging activated.")

	// Build router
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	router.LoadHTMLFiles(
		"templates/home.tmpl",
		"templates/about.tmpl",
		"templates/contact.tmpl",
		"templates/register.tmpl",
		"templates/regresult.tmpl",
		"templates/login.tmpl",
		"templates/logout.tmpl",
		"templates/global/header.tmpl",
		"templates/global/footer.tmpl")
	router.Static("/public", "public")

	router.GET("/", controllers.DisplayPage)
	router.GET("/about", controllers.DisplayPage)
	router.GET("/contact", controllers.DisplayPage)
	router.GET("/register", controllers.DisplayPage)
	router.POST("/register", controllers.AddUser)
	router.GET("/login", controllers.DisplayPage)
	router.GET("/logout", controllers.DisplayPage)
	routerAddr := "127.0.0.1:" + os.Getenv("LISTEN_PORT")

	err = router.Run(routerAddr)
	if err != nil {
		log.Fatal("Unable to serve API: ", err)
	}
}
