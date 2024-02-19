package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"app/config"
	"app/models"
)

// [POST] /song
func TestPost_Autor(t *testing.T) {
	host := fmt.Sprintf("http://localhost%s", config.Config.Port)

	req := models.AutorRequest{
		Name: "Test Song",
		Profile_path: "any profile",
		Nationality: "Brazilian",
	}

	json_data, err := json.Marshal(&req)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.Post( host+"/autor", "application/json" ,bytes.NewReader(json_data) )
	if err != nil {
		t.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println(body)
}

// // GET [Song]
// func TestGet(t *testing.T) {
// 	host := fmt.Sprintf("http://localhost%s", config.Config.Port)

// 	resp, err := http.Get(host + "/song")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer resp.Body.Close()

// 	var data *[]models.Song
// 	err = json.Unmarshal(body, data)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }