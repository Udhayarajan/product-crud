package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ph ProductHandler) GetProductsHandler(c *gin.Context) {
	products, err := ph.domain.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}
