package categories

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Category, error)
	Save(category Category) (Category, error)
	FindCategoryByBusinessID(BusinessID int) ([]Category, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(category Category) (Category, error) {
	err := r.db.Preload("Business").Create(&category).Error
	if err != nil {
		return category, err
	}
	return category, nil
}

func (r *repository) FindAll() ([]Category, error) {
	var categories []Category
	err := r.db.Preload("Business").Find(&categories).Error
	if err != nil {
		return categories, err
	}
	return categories, nil
}

func (r *repository) FindCategoryByBusinessID(BusinessID int) ([]Category, error) {
	var categories []Category
	err := r.db.Preload("Business").Where("business_id = ?", BusinessID).Find(&categories).Error
	if err != nil {
		return categories, err
	}
	return categories, nil
}
