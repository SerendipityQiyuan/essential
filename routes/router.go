package routes

import (
	"awesomeProject1/Controller/CategoryController"
	"awesomeProject1/Controller/PostController"
	"awesomeProject1/Controller/UserController"
	"awesomeProject1/Controller/uploadImagesController"
	"awesomeProject1/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
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
	categoryRoutes.GET("/categoryList", categoryController.SelectCategoryList)

	//文章上传路由
	postRoutes := r.Group("api/posts")
	postController := PostController.NewPostController()
	postRoutes.GET("/page/allList", postController.AllPageList)
	postRoutes.Use(middleware.AuthMiddleware())
	postRoutes.POST("", postController.Create)
	postRoutes.PUT("", postController.Update)
	postRoutes.GET("", postController.Show)
	postRoutes.DELETE("", postController.Delete)
	postRoutes.GET("/page/list", postController.PageList)
	postRoutes.POST("/addLike", postController.AddLike)

	//文件上传路由
	uploadImageRoutes := r.Group("api/upload")
	uploadImageController := uploadImagesController.NewUploadPortraitImagesController()
	//图片上传
	uploadImageRoutes.POST("/uploadPortraitImage", middleware.AuthMiddleware(), uploadImageController.UploadPortraitImage)

	r.StaticFS("/upload/portrait_image", http.Dir("./upload/portrait_image"))
	r.StaticFS("/upload/image", http.Dir("./upload/image"))
}
