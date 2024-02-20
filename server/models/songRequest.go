package models

// "strconv"

type SongRequest struct {
	Name        string `json:"name"`
	Year_launch string `json:"year_launch"`
	Autor_id    int32  `json:"autor_id"`
}

func (req *SongRequest) ToSong() (data Song, err error) {
	data.Name = req.Name
	// date, err := time.Parse(time.RFC3339, req.Year_launch)
	// if err != nil {
	// return
	// }

	// data.Year_launch = &date

	// autor_id, err := strconv.Atoi(req.Autor_id)
	// if err != nil {
	// return
	// }

	data.Autor_id = uint(req.Autor_id)
	return
}
