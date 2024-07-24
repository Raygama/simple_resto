package controller

import (
	"belajar/models"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// GetAllMenu adalah fungsi mendapatkan menu
// @Summary Get all menu
// @Description Get all list of menus
// @Tags Menu
// @Produce json
// @Success 200 {object} []models.Menu
// @Router /menus [get]
func GetAllMenu(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var menus []models.Menu
	db.Find(&menus)

	c.JSON(http.StatusOK, gin.H{"data" : menus})
}

// GetMenuByID adalah fungsi membuat menu berdasarkan ID
// @Summary Get menu by ID
// @Description Get a menu by ID
// @Tags Menu
// @Produce json
// @Param id path string true "Menu ID"
// @Success 200 {object} models.Menu
// @Router /menus/{id} [get]
func GetMenuByID(c *gin.Context) {
	var menu models.Menu
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&menu).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data" : menu})
}

// GetCartItemsByMenuID adalah fungsi mendapatkan cartItems berdasarkan ID Menu, cartItems adalah tabel join cart dan menu
// @Summary get cart items by menu ID
// @Description get cart items by menu ID
// @Tags Menu
// @Produce json
// @Param id path string true "Menu ID"
// @Success 200 {object} []models.CartItem
// @Router /menus/{id}/cart-items [get]
func GetCartItemsByMenuID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var cartItems []models.CartItem
	if err := db.Where("id = ?", c.Param("id")).Find(&cartItems).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data" : cartItems})
}

// GetCartByMenuID adalah fungsi mendapatkan Cart berdasarkan ID menu
// @Summary get cart by menu ID
// @Description get cart by menu ID
// @Tags Menu
// @Produce json
// @Param id path string true "Menu ID"
// @Success 200 {object} []models.Cart
// @Router /menus/{id}/carts [get]
func GetCartsByMenuID(c *gin.Context) {
	var carts []models.Cart

	db := c.MustGet("db").(*gorm.DB)
	menuID := c.Param("id")

	if err := db.Joins("JOIN cart_items ON carts.id = cart_items.cart_id").
		Where("cart_items.menu_id = ?", menuID).
		Preload("CartItems").
		Find(&carts).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(carts) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No cart associated with this menu id"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": carts})	
}

// CreateMenu adalah fungsi insert menu
// @Summary Create new menu
// @Description Create new menu
// @Tags Menu
// @Accept multipart/form-data
// @Param name formData string true "Menu name"
// @Param price formData string true "Menu price"
// @Param image formData file true "Menu image"
// @Param Authorization header string true "Authorization, how to input in swagger: 'Bearer <token>' "
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Menu
// @Router /menus [post]
func CreateMenu(c *gin.Context) {
	var menu models.Menu
	image, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
	}

	filename := uuid.New().String() + filepath.Ext(image.Filename)
	filepath := filepath.Join("uploads", filename)

	menu.Image = filename
	menu.Name = c.PostForm("name")
	menu.Price, err = strconv.Atoi(c.PostForm("price"))
	menu.ImageUrl = fmt.Sprintf("http://localhost:8080/uploads/%s", filename)
	menu.CreatedAt = time.Now()
	menu.UpdatedAt = time.Now()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	if err = c.SaveUploadedFile(image, filepath); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	if err = db.Create(&menu).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data" : menu})
}

// UpdateMenu adalah fungsi update menu yg sudah ada
// @Summary update menu
// @Description update menu
// @Tags Menu
// @Accept multipart/form-data
// @Param id formData string true "id menu"
// @Param nama formData string false "nama menu"
// @Param price formData string false "harga menu"
// @Param image formData file false "gambar menu"
// @Param Authorization header string true "Authorization, how to input in swagger: 'Bearer <token>' "
// @Security BearerToken
// @Success 200 {object} models.Menu
// @Router /menus/{id} [put]
func UpdateMenu(c *gin.Context) {
	var menu models.Menu
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.PostForm("id")).First(&menu).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Niggas"})
		return
	}

	if c.PostForm("nama") != "" {
		menu.Name = c.PostForm("nama")
	}
	if c.PostForm("price") != "" {
		menu.Price, _ = strconv.Atoi(c.PostForm("price"))
	}
	menu.UpdatedAt = time.Now()
	
	file, err := c.FormFile("image")
	if err != nil && err != http.ErrMissingFile {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	if file != nil {
		fileName := uuid.New().String() + filepath.Ext(file.Filename)
		filePath := filepath.Join("uploads", fileName)

		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
			return
		}

		oldFilepath := filepath.Join("uploads", menu.Image)
		if err := os.Remove(oldFilepath); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
			return
		}

		menu.Image = fileName
		menu.ImageUrl = fmt.Sprintf("http://localhost:8080/uploads/%s", fileName)
	}

	if err := db.Save(&menu).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data" : menu})
}

// DeleteMenu adalah fungsi delete sebuah menu
// @Summary Delete a menu
// @Description Delete a single menu by id
// @Tags Menu
// @Produce json
// @Param id path string true "id menu"
// @Param Authorization header string true "Authorization, how to input in swagger: 'Bearer <token>' "
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /menus/{id} [delete]
func DeleteMenu(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var menu models.Menu
	if err := db.Where("id = ?", c.Param("id")).First(&menu).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data" : err.Error()})
		return
	}

	if err := os.Remove(filepath.Join("uploads", menu.Image)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data" : err.Error()})
		return
	}
	db.Delete(&menu)
	c.JSON(http.StatusOK, gin.H{"data" : true})
}