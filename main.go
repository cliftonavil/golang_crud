package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
	Units uint
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	var template = template.Must(template.ParseFiles("index.html"))
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Activity := []Product{}
		db.Find(&Activity)
		template.ExecuteTemplate(w, "index.html", Activity)
	})
	router.HandleFunc("/add/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "add.html")
	})
	router.HandleFunc("/insert/", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		price := r.FormValue("price")
		unit := r.FormValue("units")
		fmt.Println(name, price, unit)
		// Create
		db.Create(&Product{Code: name, Price: 100, Units: 5000})
	})
	router.HandleFunc("/delete/{id}/", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		var product Product
		d := db.Where("id = ?", id).Delete(&product)

		fmt.Println("checkdata : ", d)
		http.Redirect(w, r, "/", http.StatusOK)
	})

	router.HandleFunc("/edit/{id}/", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		fmt.Println("***", id)
		Activity := []Product{}
		fmt.Println(Activity)
		db.Where("id = ?", id).First(&Activity)
		fmt.Println(Activity[0])
		template.ExecuteTemplate(w, "edit.html", Activity)
	})

	// Migrate the schema
	// db.AutoMigrate(&Product{})

	// Create
	// db.Create(&Product{Code: "Laptop", Price: 2000, Units: 100})
	// db.Create(&Product{Code: "TV", Price: 5000, Units: 500})

	servererror := http.ListenAndServe(":8080", router)
	if servererror != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
