package main

import (
	"html/template"
	"net/http"
	"strings"
)

type BlogController struct {
	Data struct {
		PageTitle string
		Posts     []Post
		Post      Post
	}
}

func (it *BlogController) Info() {
	if it.Data.PageTitle == "" {
		it.Data.PageTitle = "Blog page"
	}
}

func (it *BlogController) processTemplates(w http.ResponseWriter, template_files []string) {
	ts, err := template.ParseFiles(template_files...)
	if err != nil {
		BG.handleErrors(err.Error())
	}
	err = ts.Execute(w, it.Data)
	if err != nil {
		BG.handleErrors(err.Error())
	}
}

func (it *BlogController) handleRequests(w http.ResponseWriter, r *http.Request) {
	url_pattern := strings.Split(r.URL.Path, "/")
	switch url_pattern[1] {
	case "":
		it.HandleHome(w, r)
	case "new":
		it.HandleNewPost(w, r)
	case "remove":
		it.HandleRemovePost(w, r, &url_pattern)
	case "dashboard":
		it.HandleDashboard(w, r)
	default:
		http.NotFound(w, r)
		return
	}

}

func (it *BlogController) handleFavicon(w http.ResponseWriter, r *http.Request) {}

func (it *BlogController) HandleHome(w http.ResponseWriter, r *http.Request) {
	var bp BlogPost
	it.Data.PageTitle = "Posts"
	it.Data.Posts = bp.GetArchive()
	is_user_auth := BU.IsUserLogin(w, r)
	template_files := []string{
		"./templates/post_list.html",
		"./templates/base.html",
	}
	if is_user_auth {
		template_files = append(template_files, "./templates/admin_link.html")
	} else {
		template_files = append(template_files, "./templates/admin_link_nonlogin.html")
	}
	it.processTemplates(w, template_files)
}

func (it *BlogController) HandleDashboard(w http.ResponseWriter, r *http.Request) {
	if is_user_auth := BU.IsUserLogin(w, r); !is_user_auth {
		http.Redirect(w, r, "/", 301)
		return
	}
	template_files := []string{
		"./templates/dahsboard.html",
		"./templates/base.html",
		"./templates/site_link.html",
	}
	it.processTemplates(w, template_files)
}

func (it *BlogController) HandleNewPost(w http.ResponseWriter, r *http.Request) {
	if is_user_auth := BU.IsUserLogin(w, r); !is_user_auth {
		http.Redirect(w, r, "/", 301)
		return
	}
	if is_new_post := r.FormValue("new_post"); is_new_post != "" {
		var bp BlogPost
		post := Post{Title: r.FormValue("title"), Content: r.FormValue("content")}
		bp.NewPost(post)
		http.Redirect(w, r, "/", 301)
		return
	}
	template_files := []string{
		"templates/post_new.html",
		"templates/base.html",
	}
	it.processTemplates(w, template_files)
}

func (it *BlogController) HandleRemovePost(w http.ResponseWriter, r *http.Request, args *[]string) {
	if is_user_auth := BU.IsUserLogin(w, r); !is_user_auth {
		http.Redirect(w, r, "/", 301)
		return
	}
	var bp BlogPost
	var post_id string
	if post_id = r.FormValue("post_id"); post_id != "" {
		bp.RemovePost(post_id)
		http.Redirect(w, r, "/", 301)
		return
	}
	it.Data.PageTitle = "Remove post"
	if len(*args) < 3 {
		http.Redirect(w, r, "/", 301)
		return
	}
	post_id = (*args)[2]
	it.Data.Post = bp.GetPostByID(post_id)
	if it.Data.Post.Title == "" {
		http.Redirect(w, r, "/", 301)
		return
	}
	template_files := []string{
		"templates/post_remove.html",
		"templates/base.html",
	}
	it.processTemplates(w, template_files)
}
