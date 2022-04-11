package main

import (
	"cors/app"
	"cors/controller"
	"cors/helper"
	"fmt"
	"net/http"
)

func main() {

	tagsController := controller.NewsTagsController()

	router := app.NewRouter(tagsController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	fmt.Println()
	fmt.Println("************************")
	fmt.Println("**** Service Ready! ****")
	fmt.Println("************************")
	fmt.Println()

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
