package PostController

import (
	"awesomeProject1/CategoryController"
	"awesomeProject1/common"
	"awesomeProject1/model"
	"awesomeProject1/response"
	"awesomeProject1/vo"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"strconv"
)

type IPostController interface {
	CategoryController.RestController
	PageList(ctx *gin.Context)
}

type PostController struct {
	DB *gorm.DB
}

func NewPostController() IPostController {
	db := common.GetDB()
	err := db.AutoMigrate(model.Post{})
	if err != nil {
		return nil
	}
	return PostController{DB: db}
}

func (p PostController) Create(ctx *gin.Context) {
	var requestPost vo.CreatePostRequest
	//数据验证
	err := ctx.ShouldBind(&requestPost)
	log.Println("requestPost1:", requestPost)

	if err != nil {
		response.Fail(ctx, nil, "数据验证错误")
		return
	}

	//获取登录用户 user
	user, _ := ctx.Get("user")

	//创建文章
	post := model.Post{
		UserId:     user.(model.User).ID,
		CategoryId: requestPost.CategoryId,
		Title:      requestPost.Title,
		HeadImg:    requestPost.HeadImg,
		Content:    requestPost.Content,
	}

	err = p.DB.Create(&post).Error
	if err != nil {
		panic(err)
		return
	}

	response.Success(ctx, gin.H{"post": post}, "文章创建成功")
}

func (p PostController) Update(ctx *gin.Context) {
	var requestPost vo.CreatePostRequest
	//数据验证
	err := ctx.ShouldBind(&requestPost)

	if err != nil {
		println(err)
		response.Fail(ctx, nil, "数据验证错误")
		return
	}

	//获取path中的postId
	postId := ctx.Query("id")

	var post model.Post
	err = p.DB.Where("id = ?", postId).First(&post).Error
	if err != nil {
		response.Fail(ctx, nil, "文章不存在")
		return
	}
	//判断当前用户是否为文章作者
	user, _ := ctx.Get("user")
	userId := user.(model.User).ID
	if userId != post.UserId {
		response.Fail(ctx, nil, "文章不属于你，请勿非法操作")
		return
	}
	err = p.DB.Model(post).Updates(requestPost).Error
	if err != nil {
		response.Fail(ctx, nil, "更新失败")
		return
	}
	response.Success(ctx, gin.H{"post": post}, "文章更新成功")
}

func (p PostController) Show(ctx *gin.Context) {
	//获取path中的postId
	postId := ctx.Query("id")

	var post model.Post
	err := p.DB.Preload("Category").Where("id = ?", postId).First(&post).Error
	if err != nil {
		response.Fail(ctx, nil, "文章不存在")
		return
	}
	response.Success(ctx, gin.H{"post": post}, "查看成功")
}

func (p PostController) Delete(ctx *gin.Context) {
	//获取path中的postId
	postId := ctx.Query("id")

	var post model.Post
	err := p.DB.Where("id = ?", postId).First(&post).Error
	if err != nil {
		response.Fail(ctx, nil, "文章不存在")
		return
	}

	//判断当前用户是否为文章作者
	user, _ := ctx.Get("user")
	userId := user.(model.User).ID
	if userId != post.UserId {
		response.Fail(ctx, nil, "文章不属于你，请勿非法操作")
		return
	}
	err = p.DB.Delete(&post).Error
	if err != nil {
		response.Fail(ctx, nil, "删除失败")
		return
	}
	response.Success(ctx, gin.H{"post": post}, "删除成功")

}

func (p PostController) PageList(ctx *gin.Context) {
	//获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	//分页
	var posts []model.Post
	user, _ := ctx.Get("user")
	var total int64

	p.DB.Where("user_id = ?", user.(model.User).ID).Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&posts).Count(&total)
	println(posts)
	//查询记录总条数
	response.Success(ctx, gin.H{"data": posts, "total": total}, "成功")
}
