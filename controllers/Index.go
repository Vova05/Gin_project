package controllers

import (
	"Gin_project/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (_ *Controller)  GetIndex(c *gin.Context){
	var image models.TrustedCompaniesLogo
	var CustomerReview models.CustomerReview
	images, err := image.GetAllTrustedCompaniesLogo()
	CustomersReviews, err2 :=CustomerReview.GetAllCustomerReview()
	if err == nil && err2 == nil {
		data := gin.H{
			"title": "Index",
			"ImageLogo": images,
			"CustomersReviews": CustomersReviews,
		}
		c.HTML(http.StatusOK,"index.html", data)
	}else {
		c.String(http.StatusInternalServerError, err.Error())
		log.Println(err)
	}


}
//var note models.Note
//notes, err := note.GetAll()
//if err == nil {
//c.JSON(http.StatusOK, gin.H{
//"message": "All Notes",
//"notes": notes,
//})
//} else {
//c.String(http.StatusInternalServerError, err.Error())
//}
