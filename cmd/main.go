package main

import (
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/mochammadshenna/saza-aluminium/app"
	"github.com/mochammadshenna/saza-aluminium/util/helper"

	"github.com/go-playground/validator"
)

func main() {
	validate := validator.New()
	db := app.NewDB()

	fmt.Println(db)
	fmt.Println(validate)


	host := "localhost:8080"
	server := http.Server{
		Addr:    host,
	}

	fmt.Printf("Apps running on %s \n", host)
	err := server.ListenAndServe()
	helper.PanicError(err)
}
