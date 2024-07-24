package controller

import (
	"belajar/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MenuWithQuantity struct {
	models.Menu
	Quantity int `json:"quantity"`
}

// GetMenusByCartID untuk mendapat menu dari cart ID
// @Summary Get menus by cart id
// @Description Get menus from a specific cart using cart id
// @Tags Cart
// @Produce json
// @Param id path string true "id cart"
// @Success 200 {object} []MenuWithQuantity
// @Router /carts/{id}/menus [get]
func GetMenusByCartID(c *gin.Context) {
	var cartItems []models.CartItem
	var menusWithQuantity []MenuWithQuantity
	
	db := c.MustGet("db").(*gorm.DB)
	cartID := c.Param("id")
	
	if err := db.Where("cart_id = ?", cartID).Find(&cartItems).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if len(cartItems) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No menus associated with this cart id"})
		return
	}

	for _, cartItem := range cartItems {
		var menu models.Menu
		if err := db.Where("id = ?", cartItem.MenuID).First(&menu).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Menu not found!"})
			return
		}

		menuWithQuantity := MenuWithQuantity{
			Menu:     menu,
			Quantity: cartItem.Qty,
		}
		menusWithQuantity = append(menusWithQuantity, menuWithQuantity)
	}

	c.JSON(http.StatusOK, gin.H{"data": menusWithQuantity})
}

// GetCart adalah fungsi untuk menginisialisasi cart kosong yang baru
// @Summary Get empty cart
// @Description Menginisialisasi cart kosong yang nanti akan di isi oleh user id tertentu
// @Tags Cart
// @Produce json
// @Param user_id query string true "user id"
// @Success 200 {object} models.Cart
// @Router /carts [get]
func GetCart(c *gin.Context) {
	var cart models.Cart
	db := c.MustGet("db").(*gorm.DB)

	cart.TotalPrice = 0
	userIDStr := c.Query("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id!"})
		return
	}
	cart.UserID = uint(userID)
	cart.CreatedAt = time.Now()
	cart.UpdatedAt = time.Now()
	if err := db.Create(&cart).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, cart)
}


// AddMenuToCart adalah fungsi untuk menambahkan menu ke cart yang sudah ada
// @Summary add menu to cart
// @Description add menu to an existing cart
// @Param cart_id path string true "cart_id"
// @Param menu_id path string true "menu_id"
// @Tags Cart
// @Produce json
// @Success 200 {object} models.Cart
// @Router /carts/{cart_id}/menus/{menu_id} [post]
func AddMenuToCart(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	cartID := c.Param("cart_id")
	menuID := c.Param("menu_id")

	var cart models.Cart
	if err := db.Where("id = ?", cartID).First(&cart).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cart not found!"})
		return
	}

	var menu models.Menu
	if err := db.Where("id = ?", menuID).First(&menu).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Menu not found!"})
		return
	}

	var cartItem models.CartItem
	if err := db.Where("cart_id = ? AND menu_id = ?", cartID, menuID).First(&cartItem).Error; err == nil {
		// Jika CartItem sudah ada, tambahkan kuantitasnya
		cartItem.Qty += 1
	} else {
		// Jika CartItem belum ada, buat yang baru
		cartItem = models.CartItem{
			CartID: cart.Id,
			MenuID: menu.Id,
			Qty:    1,
		}
		db.Create(&cartItem)
	}

	if err := db.Save(&cartItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var totalPrice int
	db.Model(&models.CartItem{}).Where("cart_id = ?", cart.Id).Select("SUM(qty * menus.price)").Joins("JOIN menus ON cart_items.menu_id = menus.id").Scan(&totalPrice)
	cart.TotalPrice = totalPrice

	if err := db.Save(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cart})
}

// UpdateMenuInCart adalah fungsi untuk memperbarui jumlah menu dalam cart
// @Summary update menu in cart
// @Description update menu quantity in an existing cart or remove if quantity is 0
// @Param cart_id path string true "cart_id"
// @Param menu_id path string true "menu_id"
// @Param quantity formData int true "quantity"
// @Tags Cart
// @Produce json
// @Success 200 {object} models.Cart
// @Router /carts/{cart_id}/menus/{menu_id} [put]
func UpdateMenuInCart(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	cartID := c.Param("cart_id")
	menuID := c.Param("menu_id")

	Quantity, err := strconv.Atoi(c.PostForm("quantity"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	var cart models.Cart
	if err := db.Where("id = ?", cartID).First(&cart).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cart not found!"})
		return
	}

	var menu models.Menu
	if err := db.Where("id = ?", menuID).First(&menu).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Menu not found!"})
		return
	}

	var cartItem models.CartItem
	if err := db.Where("cart_id = ? AND menu_id = ?", cartID, menuID).First(&cartItem).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CartItem not found!"})
		return
	}

	if Quantity <= 0 {
		if err := db.Delete(&cartItem).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		cartItem.Qty = Quantity
		if err := db.Save(&cartItem).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	// Hitung total harga
	var totalPrice int
	db.Model(&models.CartItem{}).Where("cart_id = ?", cart.Id).Select("SUM(qty * menus.price)").Joins("JOIN menus ON cart_items.menu_id = menus.id").Scan(&totalPrice)
	cart.TotalPrice = totalPrice

	if err := db.Save(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cart})
}

// DeleteMenuFromCart adalah fungsi untuk menghapus menu dari cart
// @Summary delete menu from cart
// @Description delete menu from an existing cart
// @Param cart_id path string true "cart_id"
// @Param menu_id path string true "menu_id"
// @Tags Cart
// @Produce json
// @Success 200 {object} models.Cart
// @Router /carts/{cart_id}/menus/{menu_id} [delete]
func DeleteMenuFromCart(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	cartID := c.Param("cart_id")
	menuID := c.Param("menu_id")

	var cart models.Cart
	if err := db.Where("id = ?", cartID).First(&cart).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cart not found!"})
		return
	}

	var cartItem models.CartItem
	if err := db.Where("cart_id = ? AND menu_id = ?", cartID, menuID).First(&cartItem).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CartItem not found!"})
		return
	}

	if err := db.Delete(&cartItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Hitung total harga
	var totalPrice int
	db.Model(&models.CartItem{}).Where("cart_id = ?", cart.Id).Select("SUM(qty * menus.price)").Joins("JOIN menus ON cart_items.menu_id = menus.id").Scan(&totalPrice)
	cart.TotalPrice = totalPrice

	if err := db.Save(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cart})
}

// EmptyCart adalah fungsi untuk menghapus semua menu dari cart yang ada
// @Summary empty cart
// @Description remove all menus from an existing cart
// @Param cart_id path string true "cart_id"
// @Tags Cart
// @Produce json
// @Success 200 {object} models.Cart
// @Router /carts/{cart_id}/empty [delete]
func EmptyCart(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	cartID := c.Param("cart_id")

	var cart models.Cart
	if err := db.Where("id = ?", cartID).First(&cart).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cart not found!"})
		return
	}

	if err := db.Where("cart_id = ?", cartID).Delete(&models.CartItem{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Set total harga cart menjadi 0
	cart.TotalPrice = 0
	if err := db.Save(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cart})
}

// DeleteCart adalah fungsi untuk menghapus seluruh cart beserta semua menu yang ada di dalamnya
// @Summary delete cart
// @Description delete an entire cart along with all its menus
// @Param cart_id path string true "cart_id"
// @Tags Cart
// @Produce json
// @Success 200 {object} map[string]any
// @Router /carts/{cart_id} [delete]
func DeleteCart(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	cartID := c.Param("cart_id")

	// Hapus semua CartItem yang terkait dengan Cart
	if err := db.Where("cart_id = ?", cartID).Delete(&models.CartItem{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Hapus Cart itu sendiri
	if err := db.Where("id = ?", cartID).Delete(&models.Cart{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart and its items deleted successfully"})
}
