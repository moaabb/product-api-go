package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moaabb/product-api/pkg/db"
	"github.com/moaabb/product-api/pkg/models"
)

type ProductHandler struct {
	H db.Handler
}

func NewProductHandler(h db.Handler) ProductHandler {
	return ProductHandler{h}
}

func (m *ProductHandler) GetProducts(ctx *gin.Context) {
	var products []models.Product

	m.H.DB.Find(&products)

	ctx.JSON(http.StatusOK, products)
}

func (m *ProductHandler) CreateProduct(ctx *gin.Context) {
	var product models.Product

	ctx.BindJSON(&product)

	result := m.H.DB.Create(&product)

	if err := result.Error; err != nil {
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusCreated, product)
}
