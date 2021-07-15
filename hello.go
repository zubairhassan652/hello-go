package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, From InvoZone!")
	})
	api := r.Group("/api")
	api.GET("/detail", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"country_name": "pakistan",
			"company_name": "invozone",
			"company_type": "provides technology services around the glob",
			"website": "invozone.com",
		})
	})

	//r.LoadHTMLGlob("templates/*")
	r.LoadHTMLFiles("views/hello.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "hello.html", gin.H{
			"title": "Main website",
		})
	})


	r.Static("/assets", "./views/css")
	r.StaticFS("/static", http.Dir("./views"))
	//r.Use(static.Serve("/static", static.LocalFile("./views", true)))
	r.Run()
}

