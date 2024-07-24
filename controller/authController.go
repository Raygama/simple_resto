package controller

import (
	"belajar/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

// LoginUser godoc
// @Summary Login as a user
// @Description login with username and password
// @Tags Auth
// @Param Body body LoginInput true "the body to login a user"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /login [post]
func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	u := models.User{}
	u.Username = input.Username
	u.Password = input.Password
	tokenLogin, err := models.LoginCheck(u.Username, u.Password, db)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username or password is incorrect"})
		return
	}

	if err := db.Where("username = ?", u.Username).First(&u).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	user := map[string]string {
		"user_id" : strconv.Itoa(int(u.Id)),
		"username" : u.Username,
		"role" : u.Role,
	}

	c.JSON(http.StatusOK, gin.H{"message" : "login success", "user": user, "token" : tokenLogin})
}

// Register godoc
// @Summary register a new user
// @Description registering a user from public access
// @Tags Auth
// @Param Body body RegisterInput true "body to register a new user"
// @Success 200 {object} map[string]interface{}
// @Router /register [post]
func Register(c *gin.Context) {
	var input RegisterInput
	u := models.User{}
	db := c.MustGet("db").(*gorm.DB)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var existingUser models.User
	if err := db.Where("username = ?", input.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already taken"})
		return
	}

	u.Username = input.Username
	u.Password = input.Password
	u.Role = input.Role

	_, err := u.SaveUser(db)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user := map[string]string {
		"username" : u.Username,
		"role" : u.Role,
	}

	c.JSON(http.StatusOK, gin.H{"message" : "register success", "user": user})
}
