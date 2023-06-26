package controllers

import (
	"GoPhone/app"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func PostPhone(c *gin.Context) {

	var db *sql.DB

	//se obtiene la conexion de la base de datos

	app.Setup()
	db = app.GetDB()
	defer db.Close()

	

}
