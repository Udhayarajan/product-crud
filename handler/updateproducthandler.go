package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"product-crud/apperror"
	"product-crud/model"
)

func (ph ProductHandler) UpdateProductHandler(c *gin.Context) {
	type RequestParams struct {
		ID string `json:"id" form:"id" validate:"required,uuid"`
	}
	reqParams := RequestParams{
		ID: c.Param("id"),
	}
	if err := validate.Struct(reqParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": apperror.CustomValidationError(reqParams, err),
		})
		return
	}
	product := model.Product{}
	err := c.ShouldBindJSON(&product)
	err = validate.StructCtx(context.Background(), product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": apperror.CustomValidationError(product, err)})
		return
	}
	err = ph.domain.UpdateProduct(reqParams.ID, product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Updated successfully",
	})
}
