package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"product-crud/apperror"
	"product-crud/domain"
)

func (ph ProductHandler) GetProductHandler(c *gin.Context) {
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

	product, err := ph.domain.GetProduct(reqParams.ID)
	if err != nil {
		if errors.Is(err, domain.ErrProductNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}
