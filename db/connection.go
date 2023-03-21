package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*
  conexion a db, yo estoy usando docker con postgres, aqui dejo el comando para crear el contenedor:
 docker run --name some-postgres -e POSTGRES_USER=seba000 -e POSTGRES_PASSWORD=testpassword -p 5432:5432 -d postgres
*/

 
var DSN = "host=localhost user=seba000 password=testpassword dbname=gorm port=5432"
var DB *gorm.DB

func DBConnection() {
	//conexion a db postgress
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	}else{
		log.Println("DB connected")
	}
}