package complicaciones

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

/*************DATABASE**************/

var Dab *sql.DB

/********Revisar Directorio********/
func revisarDirectorio() {
	_, err := os.Stat("./pkg")
	fmt.Println(err)
	if os.IsNotExist(err) {
		err = os.MkdirAll("./pkg/db", 0755)
		if err != nil {
			log.Panicln(err)
		}

	}
}

/*********Abrir Base de datos***********/
func Open() {
	revisarDirectorio()
	db, err := sql.Open("sqlite3", "./pkg/db/Prueba.db")
	fmt.Println(err)
	if err != nil {
		log.Panicln(err)
	} else {
		Dab = db
		//fmt.Println("Todo anda bien")
	}

}

/*********cierra la base de dato*********/
func Close() {
	Dab.Close()
}

/***********crea las tablas************/
func CreaTabla() {
	schm := initS()

	for _, j := range schm {

		_, err := Dab.Exec(j)
		if err != nil {

			fmt.Println(err)
		}
	}
}

/*******Ping***********/
func Ping() {

	if err := Dab.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("Conexion exitosa")
	}
}

/****CRear BD****/

func CrearBD() {
	Open()
	Ping()
	CreaTabla()
	defer Close()
}
