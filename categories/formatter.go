package categories

import (
	"kedaiprogrammer/businesses"
	"time"
)

type CategoriesFormatter struct {
	ID            int                 `json:"id"`
	Category_name string              `json:"Category_name"`
	Business_id   int                 `json:"business_id" binding:"required"`
	Business      businesses.Business `json:"business"`
	Created_at    time.Time           `json:"created_at"`
	Updated_at    time.Time           `json:"updated_at"`
}

func FormatCategory(Category Category) CategoriesFormatter {
	CategoriesFormatter := CategoriesFormatter{}
	CategoriesFormatter.ID = Category.ID
	CategoriesFormatter.Category_name = Category.Category_name
	CategoriesFormatter.Business_id = Category.Business_id
	CategoriesFormatter.Business = Category.Business
	return CategoriesFormatter
}

func FormatCategories(Categories []Category) []CategoriesFormatter {
	CategoriesFormatter := []CategoriesFormatter{}
	for _, Category := range Categories {
		CategoryFormatter := FormatCategory(Category)
		CategoriesFormatter = append(CategoriesFormatter, CategoryFormatter)
	}
	return CategoriesFormatter
}
