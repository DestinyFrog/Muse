package controllers

import (
	"fmt"
	"net/http"

	"app/config"
	"app/db"
	"app/models"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func ServeSong(r *gin.Engine) {

	r.GET("/song", func(c *gin.Context) {
		var data []models.Song
		err := db.Database.Model(&models.Song{}).Preload("Autor").Find(&data).Error
		if err != nil {
			c.String(http.StatusInternalServerError, "Data not Found")
			return
		}

		c.JSON(http.StatusOK, data)
	})

	r.GET("/song/:id", func(c *gin.Context) {
		id := c.Param("id")

		var data models.Song
		err := db.Database.Model(&models.Song{}).Preload("Autor").First(&data, "id = ?", id).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Data not Found",
			})
			return
		}

		c.JSON(http.StatusOK, data)
	})

	r.GET("/song/:id/file", func(c *gin.Context) {
		id := c.Param("id")

		var data models.Song
		err := db.Database.Model(&models.Song{}).Preload("Autor").First(&data, "id = ?", id).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Data not Found",
			})
			return
		}

		c.Header("Content-Disposition", "attachment; filename="+data.Name)
		c.Header("Content-Type", "audio/mpeg")
		c.Header("Content-Length", "0")

		c.File(config.Config.Data_Files_Folder + data.File_path)
	})

	r.POST("/song", func(c *gin.Context) {
		var bodyReq models.SongRequest

		bodyReq.Name = c.PostForm("name")
		bodyReq.Autor_id = c.PostForm("autor_id")
		bodyReq.Year_launch = c.PostForm("year_launch")

		data, err := bodyReq.ToSong()
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		id := uuid.NewV4()
		data.Id = id
		data.File_path = fmt.Sprintf("%s.mp3", data.Id)

		err = db.Database.Create(&data).Error
		if err != nil {
			c.String(http.StatusInternalServerError, "Data not Saved")
			return
		}

		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		c.SaveUploadedFile(file, config.Config.Data_Files_Folder+data.File_path)

		c.Redirect(http.StatusOK, "/song")
	})

	r.PUT("/song/:id", func(c *gin.Context) {
		id := c.Param("id")

		var bodyReq models.SongRequest

		bodyReq.Name = c.PostForm("name")
		bodyReq.Autor_id = c.PostForm("autor_id")
		bodyReq.Year_launch = c.PostForm("year_launch")

		var data models.Song
		err := db.Database.First(&data, "id = ?", id).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Data not Found",
			})
			return
		}

		// db.Database.Save(&models.Song{
		// id: id,
		// })

		c.JSON(http.StatusOK, data)
	})

	r.DELETE("/song/:id", func(c *gin.Context) {
		id := c.Param("id")

		err := db.Database.Delete(&models.Song{}, "WHERE id = ?", id).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Deletado com Sucesso",
		})
	})
}
