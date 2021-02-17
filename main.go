package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Tweet struct {
	ID    int    `gorm:"primaryKey"`
	Title string `gorm:"column:title"`
}

func (Tweet) TableName() string {
	return "tweet"
}

func main() {
	db, err := gorm.Open(sqlite.Open("sample.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to open a database: %s", err.Error())
	}
	db.AutoMigrate(&Tweet{})

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		posts := []Tweet{}
		db.Find(&posts)
		c.JSON(http.StatusOK, map[string]interface{}{
			"posts": posts,
		})
	})
	r.POST("/post", func(c *gin.Context) {
		title := c.PostForm("title")
		db.Create(&Tweet{
			Title: title,
		})
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "ok",
		})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to start a server: %s", err.Error())
	}
}
