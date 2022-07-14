package CategoryController

import (
	"awesomeProject1/model"
	"awesomeProject1/repository"
	"awesomeProject1/response"
	"awesomeProject1/vo"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ICategoryController interface {
	RestController
	SelectCategoryList(ctx *gin.Context)
}

type CategoryController struct {
	Repository repository.CategoryRepository
}

func NewCategoryController() ICategoryController {
	repository := repository.NewCategoryRepository()
	err := repository.DB.AutoMigrate(model.Category{})
	if err != nil {
		fmt.Println("table create filed,err:", err)
		return nil
	}

	return CategoryController{Repository: repository}
}

func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory vo.CreateCategoryRequest
	err := ctx.ShouldBind(&requestCategory)
	if err != nil {
		response.Fail(ctx, nil, "数据验证错误，分类名称必填")
		return
	}

	var category *model.Category
	category, err = c.Repository.Create(requestCategory.Name)
	if err != nil {
		panic(err)
		return

	}

	response.Success(ctx, gin.H{"category": category}, "创建成功")
}

func (c CategoryController) Update(ctx *gin.Context) {
	var requestCategory vo.CreateCategoryRequest

	err := ctx.ShouldBind(&requestCategory)
	if err != nil {
		response.Fail(ctx, nil, "数据验证错误，分类名称必填")
		return
	}

	//获取path中的参数
	categoryID, _ := strconv.Atoi(ctx.Params.ByName("id"))
	println("categoryID:", categoryID)
	updateCategory, err := c.Repository.SelectById(categoryID)
	if err != nil {
		response.Fail(ctx, nil, "分类不存在")
		return
	}

	//更新分类
	//map struct name value
	category, err := c.Repository.Update(*updateCategory, requestCategory.Name)
	if err != nil {
		panic(err)
		return
	}

	response.Success(ctx, gin.H{"category": category}, "修改成功")
}

func (c CategoryController) Show(ctx *gin.Context) {
	categoryID, _ := strconv.Atoi(ctx.Params.ByName("id"))

	category, err := c.Repository.SelectById(categoryID)
	if err != nil {
		response.Fail(ctx, nil, "分类不存在")
		return
	}
	response.Success(ctx, gin.H{"category": category}, "查看成功")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	categoryID, _ := strconv.Atoi(ctx.Params.ByName("id"))

	err := c.Repository.DeleteById(categoryID)
	if err != nil {
		response.Fail(ctx, nil, "删除失败")
		return
	}
	response.Success(ctx, nil, "删除成功")
}

func (c CategoryController) SelectCategoryList(ctx *gin.Context) {
	category, err := c.Repository.SelectList()
	if err != nil {
		response.Fail(ctx, nil, "分类列表查询失败")
		return
	}
	response.Success(ctx, gin.H{"data": category}, "分类列表查询成功")
}
