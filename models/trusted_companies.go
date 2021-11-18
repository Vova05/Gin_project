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
