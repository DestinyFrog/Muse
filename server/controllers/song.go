package controllers

import (
	"fmt"
	"net/http"

	"app/db"
	"app/config"
	"app/models"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func ServeSong(r *gin.Engine) {

	r.GET("/song", func(c *gin.Context) {
		var data []models.Song
		err := db.Database.Model(&models.Song{}).Preload("Autor").Find(&data).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
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
				"error": err.Error(),
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

		c.File( fmt.Sprintf(config.Config.Data_Files_Folder, data.File_path) )
	})

	r.POST("/song", func(c *gin.Context) {
		var bodyReq models.SongRequest

		if err := c.ShouldBindJSON(&bodyReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		data, err := bodyReq.ToSong()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		data.Id = uuid.NewV4()
		data.File_path = fmt.Sprintf("%s.mp3", data.Id)

		err = db.Database.Create(&data).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, data)
	})

	r.POST("/song/:id/file", func(c *gin.Context) {
		id := c.Param("id")

		var data models.Song
		err := db.Database.Model(&models.Song{}).Preload("Autor").First(&data, "id = ?", id).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Data not Found",
			})
			return
		}

		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.SaveUploadedFile(file, config.Config.Data_Files_Folder+data.File_path)
		c.JSON(http.StatusOK, gin.H{
			"message": "Succesfully Uploaded File",
		})
	})

	r.PUT("/song/:id", func(c *gin.Context) {
		id := c.Param("id")

		var data models.Song
		err := db.Database.First(&data, "id = ?", id).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		var bodyReq models.SongRequest
		if err := c.ShouldBindJSON(&bodyReq); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := db.Database.Model(&data).Updates(bodyReq).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, data)
	})

	r.DELETE("/song/:id", func(c *gin.Context) {
		id := c.Param("id")

		err := db.Database.Where("id = ?", id).Delete(&models.Song{}).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Succesfully Deleted '" + id + "'",
		})
	})
}
