package uploadImagesController

import (
	"awesomeProject1/common"
	"awesomeProject1/model"
	"awesomeProject1/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"os"
	"path"
	"strings"
)

type IUploadPortraitImages interface {
	UploadPortraitImage(ctx *gin.Context)
}

type UploadPortraitImages struct {
	DB *gorm.DB
}

func NewUploadPortraitImagesController() IUploadPortraitImages {
	db := common.GetDB()
	return UploadPortraitImages{DB: db}
}

// UploadPortraitImage 上传单张图片
func (u UploadPortraitImages) UploadPortraitImage(ctx *gin.Context) {
	file, err := ctx.FormFile("portrait_image")
	if err != nil {
		fmt.Println("err:", err)
		response.Fail(ctx, nil, "图片上传失败")
		return
	}
	//将图片保存路径设为 upload/image
	dst := path.Join("./upload/portrait_image", file.Filename)
	err = ctx.SaveUploadedFile(file, dst)
	if err != nil {
		response.Fail(ctx, gin.H{"err": err}, "图片保存失败")
		return
	}
	var users model.User
	user, _ := ctx.Get("user")
	imageUrl := "http://localhost:3000/" + dst

	//先取出旧的图片路径
	var oldUsersInfo model.User
	u.DB.Where("id = ?", user.(model.User).ID).First(&oldUsersInfo)
	oldImageUrl := oldUsersInfo.PortraitImage
	oldImagePosition := strings.Index(oldImageUrl, "upload")
	fmt.Println("oldImagePosition:", oldImagePosition)
	//"http://localhost:3000/... ("/"为第20个字符)
	oldImageUrl = "./" + oldImageUrl[oldImagePosition:]
	fmt.Println("oldUrl", oldImageUrl)
	//删除旧的头像
	os.Remove(oldImageUrl)

	//向数据库写入新的头像路径
	err = u.DB.Model(&users).Where("id = ?", user.(model.User).ID).Update("portrait_image", imageUrl).Error
	if err != nil {
		response.Fail(ctx, gin.H{"err": err}, "图片上传失败")
		return
	}

	response.Success(ctx, gin.H{"url": imageUrl}, "图片上传成功")
}
