package models

type Autor struct {
	Id           uint   	`json:"id"`
	Name         string 	`json:"name"`
	Profile_path string 	`json:"profile_path"`
	Nationality  string 	`json:"nationality"`
	Songs        []Song 	`json:"songs"`
}

func (Autor) TableName() string {
	return "autor"
}