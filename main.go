package main

import (
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
		// for _, s := range Activity {
		// 	fmt.Println(s.ID, s.Code, s.Price, s.Units)
		// }
		template.ExecuteTemplate(w, "index.html", Activity)
		// http.ServeFile(w, r, "index.html")
		// template.ExecuteTemplate(w, "index.html", myProducts)
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
