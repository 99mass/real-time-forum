package main

import (
	"fmt"
	"net/http"
	"forum/controller"
	"forum/helper"
	"forum/models"
	"forum/routes"
)

var PORT = ":8080"
var Category = []models.Category{}

var CategoryName = []string{"Education", "Sport", "Art", "Culture", "Religion"}

func main() {

	// Openning the database
	db, _ := helper.CreateDatabase()

	// Create tables
	err := helper.CreateTables(db)
	if err != nil {
		fmt.Println(err)
	}
	static := http.FileServer(http.Dir("./static"))
	fs := http.FileServer(http.Dir("./css"))
	js := http.FileServer(http.Dir("./js"))
	assets := http.FileServer(http.Dir("./assets"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	http.Handle("/js/", http.StripPrefix("/js/", js))
	http.Handle("/assets/", http.StripPrefix("/assets/", assets))
	http.Handle("/static/", http.StripPrefix("/static/", static))


	// Run Handlers
	routes.Route(db)

	fmt.Println("Listening in http://localhost" + PORT)
	for _, v := range CategoryName {
		cat := models.Category{
			NameCategory: v,
		}

		controller.CreateCategory(db, cat)
	}

	http.ListenAndServe(PORT, nil)
}
