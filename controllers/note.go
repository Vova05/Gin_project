package controllers

import (
	"Gin_project/models"
	"bytes"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	"log"
	"net/http"
	"os"
)
type  Controller struct {

}

func (c *Controller) Test(ctx *gin.Context){
	var note models.Note
	notes, err := note.GetAll()
	if err == nil{

	}
	data := gin.H{
		"notes": notes,

	}
	ctx.HTML(http.StatusOK,"index.html", data)
	dataH := gin.H{
		"title": "Test page",
	}
	ctx.HTML(http.StatusOK,"header.html", dataH)
}

func (_ *Controller) GetAllNotes(c *gin.Context) {
	var note models.Note
	notes, err := note.GetAll()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "All Notes",
			"notes": notes,
		})
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
	//c.HTML(// Set the HTTP status to 200 (OK)
	//http.StatusOK,
	//			// Use the index.html template
	//			"form_post.html",
	//			// Pass the data that the page uses (in this case, 'title')
	//			gin.H{
	//				"title": "Page notes",
	//				"body": "Hello world",
	//			},
	//		)

}
func (_ *Controller) GetSingleNote(c *gin.Context) {
	var note models.Note
	id := c.Param("note_id")
	_, err := note.Fetch(id)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Single Note",
			"note": note,
		})
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
}

func (_ *Controller)  GetHome(c *gin.Context){
	var commercialOffer models.CommercialOffers
	commercialOffers, err := commercialOffer.GetAllCommercialOffers()
	if err == nil {
		data := gin.H{
			"title": "Commercial Offers",
			"commercial_offers": commercialOffers,

		}
		c.HTML(http.StatusOK,"home.html", data)
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
}

func (_ *Controller) SaveConsultation(c *gin.Context){
//r.FormValue()
	name, _ := c.GetPostForm("Name")
	phone, _ := c.GetPostForm("Phone")
	message, _ := c.GetPostForm("Message")
	var employeeName2 = "name"
	//var data  models.DataConsultation
	var add *models.DataConsultation
	data := models.DataConsultation{name,phone,message,employeeName2}
	add_data, err :=add.PostConsultation(data)
	//AddData
	//Data := &models.DataConsultation{name,phone,message}
	//log.Println(data)
	log.Println(add_data)
	log.Println(err)
}


func (_ *Controller) CreateNewNote(c *gin.Context) {
	var params models.NoteParams
	var note models.Note
	err := c.BindJSON(&params)
	if err == nil {
		_, creationError := note.Create(params)

		if creationError == nil {
			c.JSON(http.StatusCreated, gin.H{
				"message": "Note created successfully",
				"note": note,
			})
		} else {
			c.String(http.StatusInternalServerError, creationError.Error())
		}
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
}

//Image work
// writeImage encodes an image 'img' in jpeg format and writes it into ResponseWriter.
func writeImage() string{

	tmp_path:="3.jpg"
	path :="C:\\Users\\VovaGlh\\Downloads\\"+tmp_path
	m, err := os.Open(path)
	if err != nil {
		log.Println(err)
	}
	img, _, err := image.Decode(m)
	if err != nil {
		log.Println(err)
	}

	//var img image.Image = imeg
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, img, nil); err != nil {
		log.Fatalln("unable to encode image.")
	}

	str := base64.StdEncoding.EncodeToString(buffer.Bytes())
	return str
}