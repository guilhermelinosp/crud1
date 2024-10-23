package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/guilhermelinosp/crud1/config/logs"
	"github.com/guilhermelinosp/crud1/http"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logs.Error("Error loading .env file: %v", err)
	}

	gin.SetMode(os.Getenv("GIN_MODE"))

	r := gin.Default()

	http.InitUserHandler(&r.RouterGroup)

	if err := r.Run(); err != nil {
		logs.Error("Error running server: %v", err)
	}
}
