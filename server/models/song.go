package models

import (
	"fmt"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Song struct {
	Id          uuid.UUID   `json:"id"`
	Name        string     `json:"name"`
	Year_launch *time.Time `json:"year_launch"`
	File_path   string     `json:"file_path"`
	Autor_id    uint
	Autor       Autor      `json:"autor"`
}

func (Song) TableName() string {
	return "song"
}

func (o *Song) ToString() string {
	return fmt.Sprintf("%d - %s (%s)", o.Id, o.Name, o.Autor.Name)
}

type SongRequest struct {
	Name        string     `json:"name"`
	Year_launch string		`json:"year_launch"`
	Autor_id	string		`json:"autor_id"`
}

func (req *SongRequest) ToSong() (data Song, err error) {
	data.Name = req.Name
	// date, err := time.Parse(time.RFC3339, req.Year_launch)
	// if err != nil {
		// return
	// }

	// data.Year_launch = &date

	autor_id, err := strconv.Atoi(req.Autor_id)
	if err != nil {
		return
	}

	data.Autor_id = uint(autor_id)
	return
}