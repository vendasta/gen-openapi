package backend

import (
	"context"
	"fmt"
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
}

func listActivity(ctx *gin.Context) {

}

func listFields(ctx *gin.Context) {

}
