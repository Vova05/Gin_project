package main
import (
	"Gin_project/config"
	"Gin_project/controllers"
	"Gin_project/migrations"
	"github.com/gin-gonic/gin"
	"log"
)
func main() {
	_, err := config.InitializeDB()
	if err != nil {
		log.Println("Driver creation failed", err.Error())
	} else {
		// Run all migrations
		migrations.Run()

		router := gin.Default()

		var noteController controllers.NoteController
		router.GET("/notes", noteController.GetAllNotes)
		router.POST("/notes", noteController.CreateNewNote)
		router.GET("/notes/:note_id", noteController.GetSingleNote)
		router.Run(":8000")
	}
}
