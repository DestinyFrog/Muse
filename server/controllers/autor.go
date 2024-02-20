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
				"error": err.Error(),
			})
			return
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

		if err := c.ShouldBindJSON(&bodyReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		data, err := bodyReq.ToAutor()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = db.Database.Create(&data).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, data)
	})

	r.PUT("/autor/:id", func(c *gin.Context) {
		id := c.Param("id")

		var data models.Autor
		err := db.Database.First(&data, "id = ?", id).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		var bodyReq models.AutorRequest
		if err := c.ShouldBindJSON(&bodyReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
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

		c.JSON(http.StatusOK, gin.H{
			"message": "Succesfully Created",
		})
	})

	r.DELETE("/autor/:id", func(c *gin.Context) {
		id := c.Param("id")

		err := db.Database.Where("id = ?", id).Delete(&models.Autor{}).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Data not Found",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Succesfully Deleted '" + id + "'",
		})
	})
}
