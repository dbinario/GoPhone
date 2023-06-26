package main

import (
	"GoPhone/controllers"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando archivo .env")
	}

	router := gin.Default()
	router.Use(cors.Default())

	router.POST("/contacto", controllers.PostContacto)

	router.Run("localhost:8080")
}
