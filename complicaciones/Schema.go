package complicaciones

/***************Constantes*****************/

const (
	moduloSchema = `CREATE TABLE IF NOT EXISTS Modulo (
	ID INTEGER,
	Nombre TEXT NOT NULL UNIQUE,
	Descripcion TEXT,
	PRIMARY KEY (ID AUTOINCREMENT))`

	tokenSchema = `CREATE TABLE IF NOT EXISTS Tokens (
	ID	INTEGER NOT NULL,
	id_usuario	INTEGER NOT NULL,
	modulo	INTEGER NOT NULL,
	perfil	INTEGER NOT NULL,
	token	TEXT NOT NULL UNIQUE,
	fecha_de_creacion	TEXT NOT NULL,
	PRIMARY KEY(ID AUTOINCREMENT))`

	urlSchema = `CREATE TABLE IF NOT EXISTS URL (
	ID	INTEGER,
	URI	TEXT NOT NULL UNIQUE,
	IP	TEXT NOT NULL,
	Modulo	INTEGER NOT NULL,
	PRIMARY KEY(ID AUTOINCREMENT))`

	usuarioSchema = `CREATE TABLE IF NOT EXISTS Usuario (
	ID	INTEGER NOT NULL,
	Name	TEXT NOT NULL UNIQUE,
	Password	TEXT NOT NULL,
	Rol	INTEGER,
	ID_Datos_Peronales	INTEGER UNIQUE,
	PRIMARY KEY(ID AUTOINCREMENT))`
)

func initS() []string {

	schm := []string{moduloSchema, tokenSchema, urlSchema, usuarioSchema}
	return schm

}
