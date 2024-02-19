package models

import (
	"errors"
)

type AutorRequest struct {
	Name				string		`json:"name"`
	Profile_path		string		`json:"profile_path"`
	Nationality			string		`json:"nationality"`
}

func (req *AutorRequest) ToAutor() (data Autor, err error) {
	if req.Name == "" {
		err = errors.New("'Name' is null")
		return
	} else {
		data.Name = req.Name
	}

	if req.Profile_path != "" {
		data.Profile_path = req.Profile_path
	}

	if req.Nationality != "" {
		data.Nationality = req.Nationality
	}

	return
}