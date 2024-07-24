package models

import "time"

type Menu struct {
	Id uint `json:"id" gorm:"primary_key;auto_increment"`
	Name string `json:"name" gorm:"type:varchar(255)"`
	ImageUrl string `json:"imageurl" gorm:"type:varchar(255)"`
	Image string `json:"image"`
	Price int `json:"price"`
	CartItems []CartItem `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}