package models

import (
	"errors"
)

type AutorRequest struct {
	Name         string `json:"name"`
	Profile_path string `json:"profile_path"`
	Nationality  string `json:"nationality"`
}

func (req *AutorRequest) ToAutor() (data Autor, err error) {
	if req.Name == "" || req.Profile_path == "" || req.Nationality == "" {
		err = errors.New("any of the values is null")
		return
	}

	data.Name = req.Name
	data.Profile_path = req.Profile_path
	data.Nationality = req.Nationality
	return
}
