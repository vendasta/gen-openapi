package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
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
	router.StaticFS("/file", gin.Dir("static", false))
	router.Use(CORSMiddleware())

	//router.Use(ginBodyLogMiddleware)

	router.GET("/listActivity", listActivity)
	router.GET("/listFields", listFields)
	router.POST("/saveFields", saveFileds)
	router.POST("/generateYaml", generateYaml)
	router.GET("/listIntegrations", listIntegrations)
	router.GET("/viewIntegration", viewIntegrations)
	router.GET("/connectedCards", connectedCards)
	router.Run("localhost:8080")
}

type Activity struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type ActivityRequestBody struct {
	Type string `json:"type"`
}

type ActivityField struct {
	ID          string   `json:"id"`
	Type        string   `json:"type"`
	FieldType   string   `json:"fieldType"`
	Description string   `json:"description"`
	Response    Response `json:"response"`
}

type IntegrationField struct {
	CardId          string `json:"cardId"`
	CardName        string `json:"cardName"`
	CardLogo        string `json:"cardLogo"`
	CardDescription string `json:"cardDescription"`
	Connected       string `json:"connected"`
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type APIConfig struct {
	Title          string
	Version        string
	ActivityFields []ActivityField
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
	jsonResp(ctx)
	//	json := jsonResp(ctx)
	//jsonToYaml(json)
}

func listIntegrations(ctx *gin.Context) {
	jsonRespIntegrations(ctx)
}

func connectedCards(ctx *gin.Context) {
	jsonRespConnectedCards(ctx)
}

func jsonRespConnectedCards(ctx *gin.Context) []IntegrationField {
	file, err := os.Open("listIntegrations.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	// Create a JSON decoder
	decoder := json.NewDecoder(file)
	var integrations []IntegrationField
	err = decoder.Decode(&integrations)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil
	}
	var int []IntegrationField
	for _, integration := range integrations {
		if integration.Connected == "true" {

			int = append(int, integration)
		}
	}
	ctx.JSON(http.StatusOK, int)
	return int
}

func viewIntegrations(ctx *gin.Context) {
	id := ctx.Query("id")
	fmt.Println(id)
	viewData(ctx, id)

}

func viewData(ctx *gin.Context, id string) IntegrationField {
	file, err := os.Open("listIntegrations.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return IntegrationField{}
	}
	defer file.Close()

	// Create a JSON decoder
	decoder := json.NewDecoder(file)
	var integrations []IntegrationField
	err = decoder.Decode(&integrations)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return IntegrationField{}
	}

	for _, integration := range integrations {
		if integration.CardId == id {
			ctx.JSON(http.StatusOK, integration)
			return integration
		}
	}
	return IntegrationField{}
}

func jsonRespIntegrations(ctx *gin.Context) []IntegrationField {
	file, err := os.Open("listIntegrations.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	// Create a JSON decoder
	decoder := json.NewDecoder(file)
	var integrations []IntegrationField
	err = decoder.Decode(&integrations)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil
	}
	ctx.JSON(http.StatusOK, integrations)
	return integrations
}

func jsonResp(ctx *gin.Context) []ActivityField {
	file, err := os.Open("listFields.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	// Create a JSON decoder
	decoder := json.NewDecoder(file)
	var activities []ActivityField
	err = decoder.Decode(&activities)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil
	}
	ctx.JSON(http.StatusOK, activities)
	return activities
}

func jsonToYaml(data []ActivityField) {

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

func generateYaml(ctx *gin.Context) {
	ymalData, _ := generateYamlData(ctx)
	fmt.Println(ymalData)
}

func generateYamlData(ctx *gin.Context) ([]ActivityField, error) {
	//resp, _ := HttpRequest("GET", "/generateYaml")
	//var activities []ActivityField
	//errr := json.Unmarshal(resp, &activities)
	//
	//if errr != nil {
	//	fmt.Println("Cant unmarshal the byte array")
	//	return activities, nil
	//}
	//ctx.JSON(http.StatusOK, activities)
	//return activities, errr
	templateFilePath := "template.yaml"
	templateBytes, err := ioutil.ReadFile(templateFilePath)
	if err != nil {
		panic(err)
	}

	templateStr := string(templateBytes)

	t := template.New("yamlTemplate")
	parsedTemplate, err := t.Parse(templateStr)
	if err != nil {
		panic(err)
	}
	values := jsonResp(ctx)

	data := TemplateData{
		ActivityField: values,
	}
	var populatedYAML bytes.Buffer
	err = parsedTemplate.Execute(&populatedYAML, data)
	if err != nil {
		panic(err)
	}

	outputFilePath := "output.yaml"
	err = ioutil.WriteFile(outputFilePath, []byte(populatedYAML.String()), 0644)
	if err != nil {
		panic(err)
	}
	copyFile()
	return values, err
}

func saveFileds(ctx *gin.Context) {
	var activity []ActivityField

	if err := ctx.ShouldBindJSON(&activity); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jsonData, err := json.Marshal(activity)
	if err != nil {
		panic(err)
	}

	filePath := "listFields.json"

	err = ioutil.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Data inserted successfully"})

}

func HttpRequest(method string, url string) ([]byte, error) {
	req, _ := http.NewRequest(method, url, nil)
	log.Println("URL:", url)
	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return body, nil
}

type TemplateData struct {
	ActivityField []ActivityField
}

func copyFile() {
	// Source and destination file paths
	sourceFilePath := "output.yaml"
	destinationFilePath := "../src/social.yaml"

	// Open the source file for reading
	sourceFile, err := os.Open(sourceFilePath)
	if err != nil {
		fmt.Println("Error opening source file:", err)
		return
	}
	defer sourceFile.Close()

	// Create or open the destination file for writing
	destinationFile, err := os.Create(destinationFilePath)
	if err != nil {
		fmt.Println("Error creating destination file:", err)
		return
	}
	defer destinationFile.Close()

	// Copy the content from source to destination
	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		fmt.Println("Error copying file:", err)
		return
	}

	fmt.Println("File copied successfully!")
}
