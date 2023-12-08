//Dear programmer :
// When I wrote this code, only God and I knew how it worked.
// Now only God knows it!

// Therefore, if you are trying to optimize it,
// This Routine and it fails(most surely),
// please increase this counter as a warning for the next person:

//Total_hours_wasted_here = 96


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
