package models

type CartItem struct {
	Id     uint `json:"id" gorm:"primary_key;auto_increment"`
	CartID uint `json:"cart_id"`
	MenuID uint `json:"menu_id"`
	Qty    int  `json:"qty"`
	Cart   Cart `json:"-"`
	Menu   Menu `json:"-"`
}