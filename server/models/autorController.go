package models

import (
	"fmt"
	"net/http"

	"app/db"

	"github.com/gin-gonic/gin"
)

func (Autor) SetupServer(r *gin.Engine) {

	r.GET("/autor", func(c *gin.Context) {
		var data []Autor
		err := db.Database.Model(&Autor{}).Preload("Songs").Find(&data).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Data not Found",
			})
		}

		c.JSON(http.StatusOK, data)
	})

	r.GET("/autor/:id", func(c *gin.Context) {
		id := c.Param("id")

		var data Autor
		err := db.Database.Model(&Autor{}).Preload("Songs").First(&data, "id = ?", id).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Data not Found",
			})
			return
		}

		c.JSON(http.StatusOK, data)
	})

	r.GET("/autor/:id/file", func(c *gin.Context) {
		id := c.Param("id")

		var data Song
		err := db.Database.Model(&Song{}).Preload("Autor").First(&data, "id = ?", id).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Data not Found",
			})
			return
		}

		c.Header("Content-Disposition", "attachment; filename="+data.Name)
		c.Header("Content-Type", "audio/mpeg")
		c.Header("Content-Length", "0")

		c.File(data.File_path)
	})

	r.POST("/autor", func(c *gin.Context) {
		var bodyReq AutorRequest

		bodyReq.Name = c.PostForm("name")
		bodyReq.Profile_path = c.PostForm("profile_path")
		bodyReq.Nationality = c.PostForm("nationality")

		data, err := bodyReq.ToAutor()
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		err = db.Database.Create(&data).Error
		if err != nil {
			c.String(http.StatusInternalServerError, "Data not Saved")
			return
		}

		c.Redirect(http.StatusOK, "/autor")
	})

	r.DELETE("/autor/:id", func(c *gin.Context) {
		id, found := c.Params.Get("id")
		if !found {
			c.String(http.StatusBadRequest, "ID is strange")
		}

		err := db.Database.Delete(&Song{}, "id = ?", id).Error
		if err != nil {
			c.String(http.StatusInternalServerError, "Data not Found")
			return
		}

		c.String(http.StatusOK, fmt.Sprintf("Object with Id '%s' deleted", id))
	})

}
