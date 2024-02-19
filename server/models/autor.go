package models

import (
	"time"
	
	uuid "github.com/satori/go.uuid"
)

type songExtension struct {
	Id				uuid.UUID		`json:"id"`
	Name			string			`json:"name"`
	Year_launch 	*time.Time		`json:"year_launch"`
	Autor_id		uint			`json:"-"`
}

func (songExtension) TableName() string {
	return "song"
}

type Autor struct {
	Id				uint   			`json:"id"`
	Name			string 			`json:"name"`
	Profile_path	string 			`json:"profile_path"`
	Nationality		string 			`json:"nationality"`
	Songs			[]songExtension `json:"songs"`
}

func (Autor) TableName() string {
	return "autor"
}