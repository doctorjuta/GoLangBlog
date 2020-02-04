package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type BlogPost struct{}

type Post struct {
	gorm.Model
	Title   string `gorm:"type:text"`
	Content string `gorm:"type:longtext"`
}

func (p Post) String() string {
	return fmt.Sprintf("%v", p.Title)
}

func (p *BlogPost) GetArchive() []Post {
	var Posts []Post
	DB.Connect()
	defer DB.Close()
	DB.cnct.Find(&Posts)
	return Posts
}

func (p *BlogPost) NewPost(post Post) uint {
	DB.Connect()
	defer DB.Close()
	DB.cnct.Create(&post)
	return post.ID
}

func (p *BlogPost) GetPostByID(id string) Post {
	var target_post Post
	DB.Connect()
	defer DB.Close()
	DB.cnct.First(&target_post, id)
	return target_post
}

func (p *BlogPost) RemovePost(id string) {
	var target_post Post
	DB.Connect()
	defer DB.Close()
	target_post = p.GetPostByID(id)
	if target_post.Title != "" {
		DB.Connect()
		defer DB.Close()
		DB.cnct.Delete(&target_post)
	}
}
