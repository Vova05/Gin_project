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

type TrustedCompaniesLogo struct {
	ImageLogo string `json:"imageLogo"`
}
type DataTrustedCompaniesLogo struct {
	ImageLogo image.Image
}

func (trustedCompaniesLogo *TrustedCompaniesLogo)GetAllTrustedCompaniesLogo() ([]TrustedCompaniesLogo, error){
	rows, err := config.DB.Query("SELECT image FROM partner_companies")
	allTrustedCompaniesLogo := []TrustedCompaniesLogo{}
	if err == nil {
		for rows.Next() {
			var currentTrustedCompaniesLogo TrustedCompaniesLogo
			var tmp string
			rows.Scan(
				&tmp,
				)
			tmp="./Dynamic/partner_companies/"+tmp
			log.Println(tmp)
			fileImg, err2 := os.Open(tmp)
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

			currentTrustedCompaniesLogo.ImageLogo = base64.StdEncoding.EncodeToString(buffer.Bytes())
			allTrustedCompaniesLogo= append(allTrustedCompaniesLogo, currentTrustedCompaniesLogo)
		}
		return allTrustedCompaniesLogo, err
	}
	return allTrustedCompaniesLogo, err
}


//func (img1 *ImageJpg)WriteImage() []ImageJpg{
//	//tmp_path:="2.jpg"
//	//path :="C:\\Users\\VovaGlh\\Downloads\\"+tmp_path
//	//m, err := os.Open(path)
//	//if err != nil {
//	//	log.Println(err)
//	//}
//	//imeg, _, err := image.Decode(m)
//	//if err != nil {
//	//	log.Println(err)
//	//}
//	//
//	//var img image.Image = imeg
//	//buffer := new(bytes.Buffer)
//	//if err := jpeg.Encode(buffer, img, nil)
//	//	err != nil {
//	//	log.Println("unable to encode image.")
//	//}
//	//
//	//return img