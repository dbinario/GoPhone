package controllers

import (
	"GoPhone/app"
	"GoPhone/interfaces"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostContacto(c *gin.Context) {

	var db *sql.DB

	//se obtiene la conexion de la base de datos

	app.Setup()
	db = app.GetDB()
	defer db.Close() //este defer cierra la conexion cuando acaba de usarla

	var contacto interfaces.Contacto

	c.ShouldBindJSON(&contacto)

	var count int

	db.QueryRow("SELECT COUNT(*) FROM tbl_contactos WHERE telefono = LOWER(?)", contacto.Telefono).Scan(&count)

	if count > 0 {

		c.IndentedJSON(http.StatusConflict, gin.H{
			"message": "El telefono ya esta registrado",
		})

	} else {

		// Insertar datos en la tabla
		stmt, err := db.Prepare("INSERT INTO tbl_contactos(nombre, telefono) VALUES(UPPER(?),LOWER(?))")
		if err != nil {
			panic(err.Error())
		}
		defer stmt.Close()

		// Ejecutar la consulta con los parámetros
		res, err := stmt.Exec(contacto.Nombre, contacto.Telefono)
		if err != nil {
			panic(err.Error())
		}

		id, _ := res.LastInsertId()

		var conres interfaces.ContactoRes

		db.QueryRow("SELECT id,nombre,telefono FROM tbl_contactos WHERE id =?", id).Scan(&conres.ID, &conres.Nombre, &conres.Telefono)

		fmt.Println(conres)

		c.JSON(http.StatusOK, gin.H{
			"message": conres,
		})

	}

}
