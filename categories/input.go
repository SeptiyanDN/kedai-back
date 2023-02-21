package categories

type FindCategoryInput struct {
	BusinessID int `form:"business_id" binding:"required"`
}
type AddCategoryInput struct {
	Category_name string `form:"category_name" binding:"required"`
	Business_id   int    `form:"business_id" binding:"required"`
}
