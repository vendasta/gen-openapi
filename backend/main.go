package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	context.Background()
	startGin()
}

func CORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "OPTIONS", "GET", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

func startGin() {

	ctx := context.Background()
	fmt.Println(ctx)

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(CORSMiddleware())

	router.Use(ginBodyLogMiddleware)

	router.GET("/listActivity", listActivity)
	router.DELETE("/listFields", listFields)
	router.Run("localhost:8080")
}

type Activity struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

func listActivity(ctx *gin.Context) {
	file, err := os.Open("listActivity.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a JSON decoder
	decoder := json.NewDecoder(file)

	// Create a variable to hold the decoded data
	var activities []Activity

	// Use the decoder to unmarshal the JSON data into the variable
	err = decoder.Decode(&activities)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	ctx.JSON(http.StatusOK, activities)
}

func listFields(ctx *gin.Context) {

}
