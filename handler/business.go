package handler

import (
	"kedaiprogrammer/businesses"
	"kedaiprogrammer/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type businessHandler struct {
	businessServices businesses.Services
}

func NewBusinessHandler(businessServices businesses.Services) *businessHandler {
	return &businessHandler{businessServices}
}

func (h *businessHandler) SaveBusiness(c *gin.Context) {
	var input businesses.AddBusinessInput
	err := c.ShouldBind(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Create Business Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newBusiness, err := h.businessServices.SaveBusiness(input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Create Business Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := businesses.FormatBusiness(newBusiness)
	response := helper.APIResponse("Create Business Success", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *businessHandler) GetAllBusiness(c *gin.Context) {
	business, err := h.businessServices.FindAll()
	if err != nil {
		response := helper.APIResponse("Get All Product Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := businesses.FormatBusinesses(business)
	response := helper.APIResponse("Get All Business Success", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
