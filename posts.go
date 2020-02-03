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
	DB.cnct.AutoMigrate(&Post{})
	DB.cnct.Find(&Posts)
	return Posts
}

func (p *BlogPost) NewPost(post Post) int {
	DB.Connect()
	defer DB.Close()
	DB.cnct.AutoMigrate(&Post{})
	DB.cnct.Create(&post)
	return 1
}
