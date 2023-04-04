package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"product-crud/apperror"
)

func (ph ProductHandler) DeleteProductHandler(c *gin.Context) {
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

	err := ph.domain.DeleteProduct(reqParams.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Deletion successful",
	})
}
