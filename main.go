package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

type BlogPost struct {
	gorm.Model
	Title   string `gorm:"type:text;default:'Post title'"`
	Content string `gorm:"type:longtext;default:'Post content'"`
}

func (p BlogPost) String() string {
	return fmt.Sprintf("%v", p.Title)
}

func handleErrors(errtxt string) {
	panic(errtxt)
}

func pageHome(w http.ResponseWriter, r *http.Request) {
	var PageData struct {
		PageTitle string
		Posts     []BlogPost
	}
	Data := PageData
	Data.PageTitle = "Posts"
	db, err := gorm.Open(
		"mysql",
		fmt.Sprintf(
			"%v:%v/%v?parseTime=true",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_NAME"),
		),
	)
	if err != nil {
		handleErrors(err.Error())
	}
	defer db.Close()
	db.AutoMigrate(&BlogPost{})
	db.Find(&Data.Posts)
	template_files := []string{
		"templates/post_list.html",
		"templates/base.html",
	}
	ts, err := template.ParseFiles(template_files...)
	if err != nil {
		handleErrors(err.Error())
	}
	err = ts.Execute(w, Data)
	if err != nil {
		handleErrors(err.Error())
	}
}

func pageNewPost(w http.ResponseWriter, r *http.Request) {
	if is_new_post := r.FormValue("new_post"); is_new_post != "" {
		db, err := gorm.Open(
			"mysql",
			fmt.Sprintf(
				"%v:%v/%v?parseTime=true",
				os.Getenv("DB_USER"),
				os.Getenv("DB_PASS"),
				os.Getenv("DB_NAME"),
			),
		)
		if err != nil {
			handleErrors(err.Error())
		}
		defer db.Close()
		db.AutoMigrate(&BlogPost{})
		post := BlogPost{Title: r.FormValue("title"), Content: r.FormValue("content")}
		db.Create(&post)
		http.Redirect(w, r, "/", 301)
	}
	template_files := []string{
		"templates/post_new.html",
		"templates/base.html",
	}
	ts, err := template.ParseFiles(template_files...)
	if err != nil {
		handleErrors(err.Error())
	}
	err = ts.Execute(w, nil)
	if err != nil {
		handleErrors(err.Error())
	}
}

func loadENV() {
	err := godotenv.Load()
	if err != nil {
		handleErrors("Error loading .env file")
	}
}

func handleRequests(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		pageHome(w, r)
	case "/new/":
		pageNewPost(w, r)
	default:
		http.NotFound(w, r)
		return
	}

}

func handleFavicon(w http.ResponseWriter, r *http.Request) {}

func main() {
	loadENV()
	http.HandleFunc("/favicon.ico", handleFavicon)
	http.HandleFunc("/", handleRequests)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
