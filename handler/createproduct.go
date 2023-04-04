package handler

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"product-crud/apperror"
	"product-crud/domain"
	"product-crud/model"
)

func (ph ProductHandler) InsertProductHandler(c *gin.Context) {
	product := model.Product{}
	err := c.ShouldBindJSON(&product)
	err = validate.StructCtx(context.Background(), product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": apperror.CustomValidationError(product, err)})
		return
	}
	err = ph.domain.InsertProduct(product)

	if err != nil {
		if errors.Is(err, domain.ErrProductAlreadyExists) {
			c.JSON(http.StatusConflict, gin.H{
				"error": err.Error(),
			})
			return
		}
		if errors.Is(err, domain.ErrProductNameTooShort) || errors.Is(err, domain.ErrProductPriceNegative) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "product created successfully",
	})
}
