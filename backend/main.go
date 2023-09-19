package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
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
	router.GET("/listFields", listFields)
	router.Run("localhost:8080")
}

type Activity struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type ActivityRequestBody struct {
	Type string `json:"type"`
}

type activityFields struct {
	ID          string   `json:"id"`
	Type        string   `json:"type"`
	FieldType   string   `json:"fieldType"`
	Description string   `json:"description"`
	Response    Response `json:"response"`
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
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

	json := jsonResp(ctx)

	jsonToYaml(json)

	//fmt.Println(string(yaml))

}

func jsonResp(ctx *gin.Context) []activityFields {
	file, err := os.Open("listFields.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	// Create a JSON decoder
	decoder := json.NewDecoder(file)
	var activities []activityFields
	err = decoder.Decode(&activities)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil
	}
	ctx.JSON(http.StatusOK, activities)
	return activities
}

func jsonToYaml(data []activityFields) {

	yamlFile, err := os.Create("output.yaml")
	if err != nil {
		log.Fatalf("Error creating YAML file: %v", err)
	}
	defer yamlFile.Close()

	// Marshal the slice of maps into YAML format
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		log.Fatalf("Error marshaling to YAML: %v", err)
	}

	// Write the YAML data to the file
	_, err = yamlFile.Write(yamlData)
	if err != nil {
		log.Fatalf("Error writing to YAML file: %v", err)
	}

	fmt.Println("YAML data has been written to output.yaml")
}
