package database

import (
	"jerrybook/ent"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var Client *ent.Client

func Connect() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db_host := os.Getenv("DB_HOST")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_port := os.Getenv("DB_PORT")
	db_name := os.Getenv("DB_NAME")

	Client, err = ent.Open("mysql", db_host+":"+db_password+"@tcp("+db_user+":"+db_port+")/"+db_name+"?parseTime=True")
	if err != nil {
		log.Fatal(err)
	}
	//defer Client.Close()
	/*
		if err := Client.Schema.Create(context.Background()); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
		*/
}
