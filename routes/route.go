package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"belajar/controller"
	"belajar/middlewares"

	swaggerFiles "github.com/swaggo/files" //swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.CorsMiddleware())
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	r.GET("/menus", controller.GetAllMenu)
	r.GET("/menus/:id", controller.GetMenuByID)
	r.GET("/menus/:id/cart-items", controller.GetCartItemsByMenuID)
	r.GET("/menus/:id/carts", controller.GetCartsByMenuID)

	adminMiddlewareRoutes := r.Group("/menus")
	adminMiddlewareRoutes.Use(middlewares.AdminCheckMiddleware())
	adminMiddlewareRoutes.POST("/", controller.CreateMenu)
	adminMiddlewareRoutes.PUT("/:id", controller.UpdateMenu)
	adminMiddlewareRoutes.DELETE("/:id", controller.DeleteMenu)

	r.GET("/carts", controller.GetCart)
	r.GET("/carts/:id/menus", controller.GetMenusByCartID)
	r.POST("/carts/:cart_id/menus/:menu_id", controller.AddMenuToCart)
	r.PUT("/carts/:cart_id/menus/:menu_id", controller.UpdateMenuInCart)
	r.DELETE("/carts/:cart_id/menus/:menu_id", controller.DeleteMenuFromCart)
	r.DELETE("/carts/:cart_id/empty", controller.EmptyCart)
	r.DELETE("/carts/:cart_id", controller.DeleteCart)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger/doc.json")))
	r.Static("/uploads", "./uploads")

	return r
}