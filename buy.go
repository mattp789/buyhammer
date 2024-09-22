package main

import (
	// "buyhammer/modules"
	"buyhammer/modules"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)



func main() {
    r := gin.Default()
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"}, 
        AllowMethods:     []string{"GET", "POST"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge: 86400,
    }))

    r.POST("/calculate", func(c *gin.Context) {
        var json struct {
            DataItems []string `json:"dataitems"`
        }

        if err := c.ShouldBindJSON(&json); err != nil {
            log.Printf("Error binding JSON: %v", err)
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
         var response string = modules.Csviterate(json.DataItems)
        c.JSON(http.StatusOK, gin.H{"message": response})
    })

    r.Run()
}