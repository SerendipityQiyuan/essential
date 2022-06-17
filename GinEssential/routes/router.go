package routes

import (
	"awesomeProject1/CategoryController"
	"awesomeProject1/PostController"
	"awesomeProject1/UserController"
	"awesomeProject1/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) {
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())
	r.POST("/api/register", UserController.Register)
	r.POST("/api/login", UserController.Login)
	r.GET("/api/info", middleware.AuthMiddleware(), UserController.Info)

	//创建路由分组
	categoryRoutes := r.Group("api/categories")
	categoryController := CategoryController.NewCategoryController()
	categoryRoutes.POST("", categoryController.Create)
	categoryRoutes.PUT("/:id", categoryController.Update)
	categoryRoutes.GET("/:id", categoryController.Show)
	categoryRoutes.DELETE("/:id", categoryController.Delete)

	postRoutes := r.Group("api/posts")
	postRoutes.Use(middleware.AuthMiddleware())
	postController := PostController.NewPostController()
	postRoutes.POST("", postController.Create)
	postRoutes.PUT("", postController.Update)
	postRoutes.GET("", postController.Show)
	postRoutes.DELETE("", postController.Delete)
	postRoutes.POST("/page/list", postController.PageList)
}
