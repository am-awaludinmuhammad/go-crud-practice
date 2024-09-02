package categories

import (
	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAllCategories() ([]Category, error)
	CreateCategory(category *Category) error
	GetCategoryByID(id uint) (*Category, error)
	UpdateCategory(id uint, updatedCategory *Category) error
	DestroyCategory(id uint) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) GetAllCategories() ([]Category, error) {
	var categoories []Category
	result := r.db.Find(&categoories)
	if result.Error != nil {
		return nil, result.Error
	}

	return categoories, nil
}

func (r *categoryRepository) CreateCategory(category *Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) GetCategoryByID(id uint) (*Category, error) {
	var category Category
	if err := r.db.First(&category, id).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *categoryRepository) DestroyCategory(id uint) error {
	result := r.db.Delete(&Category{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *categoryRepository) UpdateCategory(id uint, updatedCategory *Category) error {
	var category Category
	if err := r.db.First(&category, id).Error; err != nil {
		return err
	}

	category.Name = updatedCategory.Name

	if err := r.db.Save(&category).Error; err != nil {
		return err
	}

	return nil
}
