package models

import "time"

type Cart struct {
	Id         uint      `json:"id" gorm:"primary_key;auto_increment"`
	TotalPrice int     `json:"total_price"`
	UserID     uint      `json:"user_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	CartItems []CartItem `json:"-"`
	User User `json:"-"`
}
