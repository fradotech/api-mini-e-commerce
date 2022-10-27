package gorm_mysql

import (
	"mini_e_commerce/src/server/model"
	"mini_e_commerce/src/server/repository"
	"time"

	"gorm.io/gorm"
)

type categoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) repository.CategoryRepo {
	return &categoryRepo{
		db: db,
	}
}

func (c *categoryRepo) GetCategories() (*[]model.Category, error) {
	var categories []model.Category

	err := c.db.Where("deleted_at IS NULL").Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return &categories, nil
}

func (c *categoryRepo) CreateCategory(category *model.Category) error {
	return c.db.Create(category).Error
}

func (p *categoryRepo) UpdateCategory(id string, category *model.Category) error {
	return p.db.Where("id = ?", id).Updates(category).Error
}

func (p *categoryRepo) DeleteCategory(id string, category *model.Category) error {
	return p.db.Model(&category).Where("id = ?", id).Update("deleted_at", time.Now()).Error
}

func (c *categoryRepo) GetCategoryById(categoryId string) (*model.Category, error) {
	var category model.Category
	err := c.db.Where("categories.id=?", categoryId).Where("categories.deleted_at IS NULL").First(&category).Error
	if err != nil {
		return nil, err
	}

	return &category, nil
}
