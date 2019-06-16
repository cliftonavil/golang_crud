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
	Price string
	Units string
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var template = template.Must(template.ParseFiles("index.html"))
		Activity := []Product{}
		db.Find(&Activity)
		template.ExecuteTemplate(w, "index.html", Activity)
	})
	router.HandleFunc("/add/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "add.html")
	})
	router.HandleFunc("/insert/", func(w http.ResponseWriter, r *http.Request) {
		var template = template.Must(template.ParseFiles("index.html"))
		if r.Method == "POST"{
			name := r.FormValue("name")
			price := r.FormValue("price")
			unit := r.FormValue("units")
			fmt.Println(name, price, unit)
			// Create
			db.Create(&Product{Code: name, Price: price, Units: unit})
			Activity := []Product{}
			db.Find(&Activity)
			template.ExecuteTemplate(w, "index.html", Activity)
		}else{
			http.ServeFile(w, r, "add.html")
		}
	})
	router.HandleFunc("/delete/{id}/", func(w http.ResponseWriter, r *http.Request) {
		var template = template.Must(template.ParseFiles("index.html"))
		vars := mux.Vars(r)
		id := vars["id"]
		var product Product
		db.Where("id = ?", id).Delete(&product)

		Activity := []Product{}
		db.Find(&Activity)
		template.ExecuteTemplate(w, "index.html", Activity)
	})

	router.HandleFunc("/edit/{id}/", func(w http.ResponseWriter, r *http.Request) {
		var template = template.Must(template.ParseFiles("edit.html"))
		vars := mux.Vars(r)
		id := vars["id"]
		Activity := []Product{}
		db.First(&Activity,id)
		template.ExecuteTemplate(w, "edit.html", Activity)
	})
	router.HandleFunc("/update/", func(w http.ResponseWriter, r *http.Request) {
		// var template = template.Must(template.ParseFiles("index.html"))
		id := r.FormValue("id")
		// name := r.FormValue("name")
		// price := r.FormValue("price")
		// units := r.FormValue("units")
		fmt.Println("ID :",id)
		// Activity := []Product{}
		val:=db.First("id = ?", id)
		fmt.Println("Updated!!!!",val)
		// db.Find(&Activity)
		// template.ExecuteTemplate(w, "index.html", Activity)
	})
	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	// db.Create(&Product{Code: "Laptop", Price: 2000, Units: 100})


	servererror := http.ListenAndServe(":8080", router)
	if servererror != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
