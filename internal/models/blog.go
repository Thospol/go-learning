package models

// Blog blog info
type Blog struct {
	Model
	Author  Author `gorm:"embedded;embeddedPrefix:author_" json:"author"`
	Title   string `json:"title"`
	Message string `json:"message"`
}
