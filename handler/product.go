package handler

import (
	"crud-golang/helper"
	"crud-golang/product"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	product product.Service
}

func NewHandler(serviceProduct product.Service) *productHandler {
	return &productHandler{serviceProduct}
}

func (h *productHandler) CreateProductHandler(c *gin.Context) {
	var input product.CreateProductInput

	err := c.ShouldBind(&input)

	if err != nil {
		errorMessage := helper.FormatValidationError(err)
		response := helper.APIResponse("Product Error Create", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	product, err := h.product.CreateProduct(input)

	if err != nil {
		response := helper.APIResponse("Product Error Create", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIResponse("Product Success Created", http.StatusCreated, "success", product)
	c.JSON(http.StatusCreated, response)
}

func (h *productHandler) GetAllProductHandler(c *gin.Context) {
	product, err := h.product.GetAllProduct()
	if err != nil {
		response := helper.APIResponse("Error to Get Data Products", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("List of products", http.StatusOK, "success", product)
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) GetProductByIDHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	product, err := h.product.GetByIDProduct(id)
	if err != nil {
		response := helper.APIResponse("Error to Get Data Product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Data product", http.StatusOK, "success", product)
	c.JSON(http.StatusOK, response)

}

func (h *productHandler) UpdateProductHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input product.CreateProductInput

	err := c.ShouldBind(&input)
	if err != nil {
		errorMessage := helper.FormatValidationError(err)
		response := helper.APIResponse("Product Error Update", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	product, err := h.product.UpdateProduct(id, input)
	if err != nil {
		response := helper.APIResponse("Product Error Update", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIResponse("Product Success Updated", http.StatusOK, "success", product)
	c.JSON(http.StatusCreated, response)
}

func (h *productHandler) DeleteProductHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	product, err := h.product.DeleteProduct(id)
	if err != nil {
		response := helper.APIResponse("Error to Delete Data Product", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Product Success Deleted", http.StatusOK, "success", product)
	c.JSON(http.StatusOK, response)

}
