package PostController

import (
	"awesomeProject1/Controller/CategoryController"
	"awesomeProject1/common"
	"awesomeProject1/model"
	"awesomeProject1/response"
	"awesomeProject1/vo"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"strconv"
	"strings"
)

type IPostController interface {
	CategoryController.RestController
	AddLike(ctx *gin.Context)
	PageList(ctx *gin.Context)
	AllPageList(ctx *gin.Context)
}

type PostController struct {
	DB *gorm.DB
}

func NewPostController() IPostController {
	db := common.GetDB()
	err := db.AutoMigrate(model.Post{})
	if err != nil {
		panic(err)
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

// PageList 查询个人的文章记录
func (p PostController) PageList(ctx *gin.Context) {
	//获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	//分页
	var posts []model.Post
	user, _ := ctx.Get("user")
	var total int64

	p.DB.
		Where("user_id = ?", user.(model.User).ID).
		Preload("Category").
		Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Order("created_at desc").
		Find(&posts)

	p.DB.Model(&posts).Where("user_id = ?", user.(model.User).ID).Count(&total)
	//查询记录总条数
	response.Success(ctx, gin.H{"data": &posts, "total": total}, "成功")
}

// AllPageList 查询全部文章记录
func (p PostController) AllPageList(ctx *gin.Context) {
	//获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))
	fmt.Println(pageNum, pageSize)

	//分页
	var posts []model.Post
	var total int64
	p.DB.
		Order("created_at desc").
		Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&posts)
	p.DB.Model(&posts).Count(&total)
	println(posts)
	//查询记录总条数
	response.Success(ctx, gin.H{"data": &posts, "total": total}, "成功")
}

func (p PostController) AddLike(ctx *gin.Context) {
	var requestEvaluation vo.RequestEvaluation

	err := ctx.Bind(&requestEvaluation)

	if err != nil {
		response.Fail(ctx, nil, "数据错误")
		return
	}

	var post model.Post
	id := requestEvaluation.PostId

	//判断该用户是否点过赞
	err = p.DB.Where("id = ?", id).First(&post).Error
	users := post.LikeUsers

	//将字符串转换为slice
	likeUserSlice := strings.Split(users, ",")
	likeUserMap := make(map[string]struct{}, len(likeUserSlice))

	//将slice转换为map
	for _, item := range likeUserSlice {
		likeUserMap[item] = struct{}{}
	}
	fmt.Println(likeUserMap)

	user, _ := ctx.Get("user")

	//判断用户id是否存在于likeUserMap
	_, ok := likeUserMap[strconv.Itoa(int(user.(model.User).ID))]

	//存在
	if ok {
		response.Fail(ctx, nil, "该用户已点过赞")
		return
	}

	//不存在
	//将用户加入likeUserSlice
	likeUserSlice = append(likeUserSlice, strconv.Itoa(int(user.(model.User).ID)))
	//将slice拼接回字符串
	users = strings.Join(likeUserSlice, ",")
	//将新的like_users字符串写入post表
	p.DB.Model(&post).Update("like_users", users)

	//更新点赞记录
	err = p.DB.Model(post).Where("id", id).First(&post).Updates(requestEvaluation).Error
	if err != nil {
		response.Fail(ctx, nil, "点赞失败")
		return
	}
	response.Success(ctx, gin.H{"data": post}, "点赞成功")
}
