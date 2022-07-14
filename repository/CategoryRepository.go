package repository

import (
	"awesomeProject1/common"
	"awesomeProject1/model"
	"fmt"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository() CategoryRepository {
	return CategoryRepository{DB: common.GetDB()}
}

func (c CategoryRepository) Create(name string) (*model.Category, error) {
	category := model.Category{Name: name}

	if err := c.DB.Create(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c CategoryRepository) Update(category model.Category, name string) (*model.Category, error) {
	if err := c.DB.Model(&category).Update("name", name).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c CategoryRepository) SelectById(id int) (*model.Category, error) {
	category := model.Category{}
	if err := c.DB.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c CategoryRepository) DeleteById(id int) error {
	category := model.Category{}
	if err := c.DB.First(&category, id).Error; err != nil {
		return err
	}
	c.DB.Delete(&category, id)
	return nil
}

func (c CategoryRepository) SelectList() ([]model.Category, error) {
	var category []model.Category
	err := c.DB.Raw("select * from categories").Find(&category).Error
	fmt.Println("category:", category)
	if err != nil {
		return nil, err
	}
	return category, nil
}
