package models

import (
	"Gin_project/config"
	"time"
)

type ProfileClient struct {
	Name string `json:"name"`
	Surname string `json:"surname"`
	Mail string `json:"mail"`
	Telephone string `json:"telephone"`
	Password string `json:"password"`
	TimeData time.Time `json:"time_data"`
}
type DataProfileClient struct {
	Name string
	Surname string
	Mail string
	Telephone string
	Password string
	TimeData time.Time
}

type DataChange struct {
	Year int `json:"year"`
	Month time.Month `json:"month"`
	Day int `json:"day"`
	TimeData time.Time `json:"time_data"`
}
type DataChangeDate struct {
	Year int
	Month time.Month
	Day int
	TimeData time.Time
}
func (profile *DataProfileClient) CreateProfileClient(data DataProfileClient) (*DataProfileClient, error) {
	data.TimeData = time.Now().UTC()
	statement, _ := config.DB.Prepare("INSERT INTO client_profile (name, surname, mail,telephone , password, time) VALUES (?, ?, ?, ?, ?, ?)")
	_ , err := statement.Exec(data.Name, data.Surname, data.Mail,data.Telephone,data.Password,data.TimeData)

	return &data, err
}
