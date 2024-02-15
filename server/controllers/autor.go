package controllers

import (
	"net/http"

	"app/db"
	"app/models"

	"github.com/gin-gonic/gin"
)

func ServeAutor(r *gin.Engine) {

	r.GET("/autor", func(c *gin.Context) {
		var data []models.Autor
		err := db.Database.Model(&models.Autor{}).Preload("Songs").Find(&data).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Data not Found",
			})
		}

		c.JSON(http.StatusOK, data)
	})

	r.GET("/autor/:id", func(c *gin.Context) {
		id := c.Param("id")

		var data models.Autor
		err := db.Database.Model(&models.Autor{}).Preload("Songs").First(&data, "id = ?", id).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Data not Found",
			})
			return
		}

		c.JSON(http.StatusOK, data)
	})

	r.POST("/autor", func(c *gin.Context) {
		var bodyReq models.AutorRequest

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

}