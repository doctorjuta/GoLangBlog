package main

import (
	"html/template"
	"net/http"
)

type BlogController struct {
	Data struct {
		PageTitle string
		Posts     []Post
	}
}

func (obj *BlogController) Info() {
	if obj.Data.PageTitle == "" {
		obj.Data.PageTitle = "Blog page"
	}
}

func (obj *BlogController) processTemplates(w http.ResponseWriter, template_files []string) {
	ts, err := template.ParseFiles(template_files...)
	if err != nil {
		BG.handleErrors(err.Error())
	}
	err = ts.Execute(w, obj.Data)
	if err != nil {
		BG.handleErrors(err.Error())
	}
}

func (obj *BlogController) handleRequests(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		obj.HandleHome(w, r)
	case "/new/":
		obj.HandleNewPost(w, r)
	default:
		http.NotFound(w, r)
		return
	}

}

func (obj *BlogController) handleFavicon(w http.ResponseWriter, r *http.Request) {}

func (bcnt *BlogController) HandleHome(w http.ResponseWriter, r *http.Request) {
	var bp BlogPost
	bcnt.Data.PageTitle = "Posts"
	bcnt.Data.Posts = bp.GetArchive()
	template_files := []string{
		"./templates/post_list.html",
		"./templates/base.html",
	}
	bcnt.processTemplates(w, template_files)
}

func (bcnt *BlogController) HandleNewPost(w http.ResponseWriter, r *http.Request) {
	if is_new_post := r.FormValue("new_post"); is_new_post != "" {
		var bp BlogPost
		post := Post{Title: r.FormValue("title"), Content: r.FormValue("content")}
		bp.NewPost(post)
		http.Redirect(w, r, "/", 301)
	}
	template_files := []string{
		"templates/post_new.html",
		"templates/base.html",
	}
	bcnt.processTemplates(w, template_files)
}
