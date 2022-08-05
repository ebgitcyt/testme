package main

import (
	"fmt"
	"importes/complicaciones"

	_ "github.com/mattn/go-sqlite3"
)

/*************MAIN****************/

func main() {
	fmt.Println("apriete 1 para crear tablas Predeterminadas")
	var option int
	fmt.Scanln(&option)

	if option == 1 {
		complicaciones.CrearBD()
	}
}
