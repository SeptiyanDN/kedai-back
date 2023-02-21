package handler

import (
	"kedaiprogrammer/categories"
	"kedaiprogrammer/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type categoryHandler struct {
	categoryServices categories.Services
}

func NewCategoryHandler(categoryServices categories.Services) *categoryHandler {
	return &categoryHandler{categoryServices}
}

func (h *categoryHandler) SaveCategory(c *gin.Context) {
	var input categories.AddCategoryInput
	err := c.ShouldBind(&input)

	if err != nil {
		response := helper.APIResponse("Create New Category Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newCategory, err := h.categoryServices.SaveCategory(input)
	if err != nil {
		response := helper.APIResponse("Create New Category Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := categories.FormatCategory(newCategory)
	response := helper.APIResponse("Create Category Success", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *categoryHandler) GetAllCategory(c *gin.Context) {
	category, err := h.categoryServices.FindAll()
	if err != nil {
		response := helper.APIResponse("Failed to Get All Categories", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := categories.FormatCategories(category)
	response := helper.APIResponse("Get All Categories Successfully", http.StatusOK, "Success", formatter)
	c.JSON(http.StatusOK, response)
}
