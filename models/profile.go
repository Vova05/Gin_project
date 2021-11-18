package models

import (
	"Gin_project/config"
)

//type ProfileClient struct {
//	Name string `json:"name"`
//	Surname string `json:"surname"`
//	Mail string `json:"mail"`
//	Telephone string `json:"telephone"`
//	Password string `json:"password"`
//	TimeData time.Time `json:"time_data"`
//}
//type DataProfileClient struct {
//	Name string
//	Surname string
//	Mail string
//	Telephone string
//	Password string
//	TimeData time.Time
//}


func getMail(mailName string){

}
func (profile *ProfileClient) GetProfile(mail string) (*ProfileClient, error) {
	err := config.DB.QueryRow(
		"SELECT name,surname,mail,telephone FROM client_profile WHERE mail=?", mail).Scan(
		&profile.Name,&profile.Surname, &profile.Mail, &profile.Telephone)
	return profile, err
}