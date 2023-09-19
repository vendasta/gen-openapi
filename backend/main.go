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

	//router.Use(ginBodyLogMiddleware)

	router.GET("/listActivity", listActivity)
	router.DELETE("/listFields", listFields)
	router.Run("localhost:8080")
}

type Activity struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type ActivityRequestBody struct {
	Type string `json:"type"`
}

func listActivity(ctx *gin.Context) {
	var requestBody ActivityRequestBody
	var response []Activity
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filterType := requestBody.Type

	file, err := os.Open("listActivity.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var activities []Activity
	err = decoder.Decode(&activities)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	for _, activity := range activities {
		if activity.Type == filterType {
			response = append(response, activity)
		}

		if activity.Type == "" {
			response = append(response, activity)
		}
	}
	ctx.JSON(http.StatusOK, response)
}

func listFields(ctx *gin.Context) {

}
