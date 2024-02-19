package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type autorExtension struct {
	Id				uint   			`json:"id"`
	Name			string 			`json:"name"`
 // Profile_path	string 			`json:"profile_path"`
 // Nationality		string 			`json:"nationality"`
}

func (autorExtension) TableName() string {
	return "autor"
}

type Song struct {
	Id				uuid.UUID			`json:"id"`
	Name			string				`json:"name"`
	Year_launch		*time.Time			`json:"year_launch"`
	File_path		string				`json:"-"`
	Autor_id		uint				`json:"-"`
	Autor			autorExtension		`json:"autor"`
}

func (Song) TableName() string {
	return "song"
}