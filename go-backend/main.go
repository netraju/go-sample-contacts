// main.go

package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	a := App{}
	err := godotenv.Load("../.env")

	if err != nil {
		fmt.Println(err)
	}
	a.Initialize(
		os.Getenv("APP_DB_SERVER"),
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	a.Run(":8010")
}
