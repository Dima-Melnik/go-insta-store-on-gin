package handlers

import (
	"net/http"

	"github.com/Dima-Melnik/go-insta-store-on-gin/internal/models"
	"github.com/Dima-Melnik/go-insta-store-on-gin/internal/models/dto"
	"github.com/Dima-Melnik/go-insta-store-on-gin/internal/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllProducts(c *gin.Context) {
	products, err := h.productService.GetAll()
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSON(c, http.StatusOK, products)
}

func (h *Handler) GetProductByID(c *gin.Context) {
	id, err := utils.GetID(c)
	if err != nil {
		return
	}

	product, err := h.productService.GetByID(id)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSON(c, http.StatusOK, product)
}

func (h *Handler) GetProductByName(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		utils.SendError(c, http.StatusBadRequest, map[string]string{
			"error": "missing 'name' query parameter",
		})
	}

	product, err := h.productService.GetByName(name)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSON(c, http.StatusOK, product)
}

func (h *Handler) CreateProduct(c *gin.Context) {
	var createProduct dto.CreateProduct

	if !utils.BindJSON(c, createProduct) {
		return
	}

	product := models.Product{
		Name:       createProduct.Name,
		Desc:       createProduct.Desc,
		ImageURL:   createProduct.ImageURL,
		Price:      createProduct.Price,
		CategoryID: createProduct.CategoryID,
		Category:   models.Category{},
		Attributes: []models.Attribute{},
	}

	createdProduct, err := h.productService.Create(product)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSON(c, http.StatusCreated, dto.ToProductResponse(createdProduct))
}

func (h *Handler) UpdateProduct(c *gin.Context) {
	var updateProduct dto.UpdateProduct

	if !utils.BindJSON(c, updateProduct) {
		return
	}

	id, err := utils.GetID(c)
	if err != nil {
		return
	}

	product := models.Product{
		Name:       *updateProduct.Name,
		Desc:       *updateProduct.Desc,
		ImageURL:   *updateProduct.ImageURL,
		Price:      *updateProduct.Price,
		CategoryID: *updateProduct.CategoryID,
		Category:   models.Category{},
		Attributes: []models.Attribute{},
	}

	updatedProduct, err := h.productService.Update(id, product)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSON(c, http.StatusOK, dto.ToProductResponse(updatedProduct))
}

func (h *Handler) Delete(c *gin.Context) {
	id, err := utils.GetID(c)
	if err != nil {
		return
	}

	err = h.productService.Delete(id)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSON(c, http.StatusOK, "successfully deleted")
}
