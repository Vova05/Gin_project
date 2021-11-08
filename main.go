package main

import (
	"Gin_project/config"
	"Gin_project/controllers"
	"Gin_project/middlewares"
	_ "github.com/gin-gonic/contrib/secure"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
	"log"
)
func main() {
	_, err := config.InitializeDB()
	if err != nil {
		log.Println("Driver creation failed", err.Error())
	} else {
		// Run all migrations
		//migrations.Run()

		//server := gin.Default()
		server := gin.New()
		//server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())

		var noteController controllers.NoteController
		//server.Use(static.Serve("/",static.LocalFile("./views",true)))

		server.Static("/css","./templates/css")
		server.LoadHTMLGlob("templates/*.html")

		view_page := server.Group("/view_bank")
		{
			view_page.GET("/home", noteController.GetHome)
//			view_page.GET("/about/bank")
	//		view_page.GET("/about/company")
		}

		main_groupe := server.Group("/note",gin.Recovery(),middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())
		{
			//main_groupe.Use(gin.Recovery(), )
			main_groupe.GET("/notes", noteController.GetAllNotes)
			main_groupe.POST("/notes", noteController.CreateNewNote)
			main_groupe.GET("/notes/:note_id", noteController.GetSingleNote)
		}

		//server.GET("/notes", )
		//server.GET("/form_post", func(c *gin.Context) {
		//	c.HTML(// Set the HTTP status to 200 (OK)
		//		http.StatusOK,
		//		// Use the index.html template
		//		"form_post.html",
		//		// Pass the data that the page uses (in this case, 'title')
		//		gin.H{
		//			"title": "Home Page",
		//			"body": "Hello world",
		//		},
		//	)
		//})

		view := server.Group("/view", gindump.Dump())
		{
			view.GET("/test",gin.Recovery(), noteController.Test)
		}
		server.Run(":8000")
	}
}
