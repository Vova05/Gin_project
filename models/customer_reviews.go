package models

import (
	"Gin_project/config"
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
	"log"
	"os"
)

type CustomerReview struct {
	ImageCustomer string `json:"imageCustomer"`
	Feedback string `json:"feedback"`
	Name string `json:"name"`
}
type DataCustomerReview struct {
	ImageCustomer string
	Feedback string
	Name string
}

func (customerReview *CustomerReview)GetAllCustomerReview() ([]CustomerReview, error){
	rows, err := config.DB.Query("SELECT photo,feedback,name FROM customer_reviews")
	allCustomerReview := []CustomerReview{}
	if err == nil {
		for rows.Next() {
			var currentCustomerReview CustomerReview
			rows.Scan(
				&currentCustomerReview.ImageCustomer,
				&currentCustomerReview.Feedback,
				&currentCustomerReview.Name,
			)
			currentCustomerReview.ImageCustomer="./Dynamic/customer_reviews_foto/"+currentCustomerReview.ImageCustomer
			log.Println(currentCustomerReview.ImageCustomer)
			fileImg, err2 := os.Open(currentCustomerReview.ImageCustomer)
			if err2 != nil{
				log.Println(err2)
				continue
			}
			imageCode,_,err3 := image.Decode(fileImg)
			if err3 != nil{
				log.Println(err3)
				continue
			}
			buffer := new(bytes.Buffer)
			if err := jpeg.Encode(buffer, imageCode, nil)
				err != nil {
				log.Println("unable to encode image.")
				continue
			}

			currentCustomerReview.ImageCustomer = base64.StdEncoding.EncodeToString(buffer.Bytes())
			allCustomerReview= append(allCustomerReview, currentCustomerReview)
		}
		return allCustomerReview, err
	}
	return allCustomerReview, err
}