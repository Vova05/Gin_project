package models

import (
	"Gin_project/config"
	_ "image"
	"log"
	"time"
	//"github.com/username/notes_api_layered/config"
)
type Note struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type NoteParams struct {
	Title string
	Body  string
}

func (note *Note) GetAll() ([]Note, error) {
	rows, err := config.DB.Query("SELECT * FROM notes")
	allNotes := []Note{}
	if err == nil {
		for rows.Next() {
			var currentNote Note
			rows.Scan(
				&currentNote.Id,
				&currentNote.Title,
				&currentNote.Body,
				&currentNote.CreatedAt,
				&currentNote.UpdatedAt)
			allNotes = append(allNotes, currentNote)
		}
		return allNotes, err
	}
	return allNotes, err
}
func (note *Note) Fetch(id string) (*Note, error) {
	err := config.DB.QueryRow(
		"SELECT id, title, body, created_at, updated_at FROM notes WHERE id=?", id).Scan(
		&note.Id, &note.Title, &note.Body, &note.CreatedAt, &note.UpdatedAt)
	return note, err
}

type CommercialOffers struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Price	  float64	`json:"price"`
	Link 	  string	`json:"link"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (commercialOffers *CommercialOffers) GetAllCommercialOffers() ([]CommercialOffers, error) {
	rows, err := config.DB.Query("SELECT * FROM CommercialOffers")
	allCommercialOffers := []CommercialOffers{}
	if err == nil {
		for rows.Next() {
			var currentCommercialOffers CommercialOffers
			rows.Scan(
				&currentCommercialOffers.Id,
				&currentCommercialOffers.Title,
				&currentCommercialOffers.Body,
				&currentCommercialOffers.Price,
				&currentCommercialOffers.CreatedAt,
				&currentCommercialOffers.UpdatedAt,
				&currentCommercialOffers.Link)
			allCommercialOffers = append(allCommercialOffers, currentCommercialOffers)
		}
		return allCommercialOffers, err
	}
	return allCommercialOffers, err
}

type DataConsultation struct {
	Name    string
	Phone      string
	Message	  string
	EmployeeName	string
}



func (data *DataConsultation) PostConsultation(dataConsultation DataConsultation) (*DataConsultation,error){

	var created_at = time.Now().UTC()
	var updated_at = time.Now().UTC()
	statement, _ := config.DB.Prepare("INSERT INTO consultations (client_name, client_telephone, client_message,employee_name,created_at, updated_at) VALUES (?, ?, ?, ?,?,?)")
	_ , err := statement.Exec(dataConsultation.Name,dataConsultation.Phone,dataConsultation.Message,dataConsultation.EmployeeName, created_at, updated_at)

	//log.Println("Unable to create note", err.Error())
	return &dataConsultation, err

}
func (note *Note) Create(data NoteParams) (*Note, error) {
	var created_at = time.Now().UTC()
	var updated_at = time.Now().UTC()
	statement, _ := config.DB.Prepare("INSERT INTO notes (title, body, created_at, updated_at) VALUES (?, ?, ?, ?)")
	result, err := statement.Exec(data.Title, data.Body, created_at, updated_at)
	if err == nil {
		id, _ := result.LastInsertId()
		note.Id = int(id)
		note.Title = data.Title
		note.Body = data.Body
		note.CreatedAt = created_at
		note.UpdatedAt = updated_at
		return note, err
	}
	log.Println("Unable to create note", err.Error())
	return note, err
}

//image

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
//	//buffer := new(bytes.Buffer)
//	//if err := jpeg.Encode(buffer, imeg, nil)
//	//	err != nil {
//	//	log.Println("unable to encode image.")
//	//}
//	//
//	//return imeg
//
//
//
//	//w.Set("Content-Type", "image/jpeg")
//	//w.Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
//	//w.Header().Set("Content-Type", "image/jpeg")
//	//w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
//	//if _, err := w.Writer.Write(buffer.Bytes()); err != nil {
//	//	log.Println("unable to write image.")
//	//}
//	var str[] ImageJpg
//	for i:=2;i<3;i++{
//		tmp :=strconv.Itoa(i)
//		tmp_path:=tmp+".jpg"
//		log.Println(tmp_path)
//		path :="C:\\Users\\VovaGlh\\Downloads\\"+tmp_path
//		m, err := os.Open(path)
//		if err != nil {
//			log.Println(err)
//		}
//		img, _, err := image.Decode(m)
//		if err != nil {
//			log.Println(err)
//		}
//
//		//var img image.Image = imeg
//		buffer := new(bytes.Buffer)
//		if err := jpeg.Encode(buffer, img, nil); err != nil {
//			log.Fatalln("unable to encode image.")
//		}
//
//		ImageTmp:=ImageJpg{ImageData: base64.StdEncoding.EncodeToString(buffer.Bytes())}
//		str =append(str,ImageTmp)
//	}
//
//	return str
//}
//
//type ImageJpg struct {
//	ImageData string
//}
