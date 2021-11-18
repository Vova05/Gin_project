package controllers

import (
	"Gin_project/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func (con *Controller) SaveProfile(c *gin.Context){

	name, _ := c.GetPostForm("Name")
	surname, _ := c.GetPostForm("Surname")
	mail, _ := c.GetPostForm("Mail")
	telephone, _ := c.GetPostForm("Telephone")
	password, _ := c.GetPostForm("Password")


	var add *models.DataProfileClient
	data := models.DataProfileClient{name,surname,mail,telephone,password, time.Now()}
	add_data, err :=add.CreateProfileClient(data)

	log.Println(add_data)
	log.Println(err)
}

func (con *Controller)  GetRegistration(c *gin.Context){
		data := gin.H{
			"title": "Registration",
		}
		c.HTML(http.StatusOK,"Registration.html", data)
}
