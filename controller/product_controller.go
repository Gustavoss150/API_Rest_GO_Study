package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) ProductController {
	return ProductController{
		productUsecase: usecase,
	}
}

func (p *ProductController) GetProducts(ctxt *gin.Context) {
	products, err := p.productUsecase.GetProducts()
	if err != nil {
		ctxt.JSON(http.StatusInternalServerError, err)
	}

	ctxt.JSON(http.StatusOK, products)
}

func (p *ProductController) CreateProduct(ctxt *gin.Context) {
	var product model.Product
	err := ctxt.BindJSON(&product)

	if err != nil {
		ctxt.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.productUsecase.CreateProduct(product)

	if err != nil {
		ctxt.JSON(http.StatusInternalServerError, err)
		return
	}

	ctxt.JSON(http.StatusCreated, insertedProduct)
}
