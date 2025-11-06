package handlers

import (
	"net/http"

	"github.com/Dima-Melnik/go-insta-store-on-gin/internal/models"
	"github.com/Dima-Melnik/go-insta-store-on-gin/internal/models/dto"
	"github.com/Dima-Melnik/go-insta-store-on-gin/internal/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllCategories(c *gin.Context) {
	categories, err := h.categoryService.GetAll()
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err.Error())
	}

	utils.SendJSON(c, http.StatusOK, categories)
}

func (h *Handler) GetCategoryByID(c *gin.Context) {
	id, err := utils.GetID(c)
	if err != nil {
		return
	}

	category, err := h.categoryService.GetByID(id)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSON(c, http.StatusOK, category)
}

func (h *Handler) GetCategoryByName(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		utils.SendError(c, http.StatusBadRequest, map[string]string{
			"error": "missing 'name' query parameter",
		})
	}

	category, err := h.categoryService.GetByName(name)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSON(c, http.StatusOK, category)
}

func (h *Handler) CreateCategory(c *gin.Context) {
	var createCategory dto.CreateCategory

	if !utils.BindJSON(c, createCategory) {
		return
	}

	category := models.Category{
		Name:     createCategory.Name,
		Products: []models.Product{},
	}

	createdCategory, err := h.categoryService.Create(category)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSON(c, http.StatusCreated, dto.ToCategoryResponse(createdCategory))
}

func (h *Handler) UpdateCategory(c *gin.Context) {
	var updateCategory dto.UpdateCategory

	if !utils.BindJSON(c, updateCategory) {
		return
	}

	id, err := utils.GetID(c)
	if err != nil {
		return
	}

	category := models.Category{
		Name:     *updateCategory.Name,
		Products: []models.Product{},
	}

	updatedCategory, err := h.categoryService.Update(id, category)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSON(c, http.StatusOK, dto.ToCategoryResponse(updatedCategory))
}

func (h *Handler) DeleteCategory(c *gin.Context) {
	id, err := utils.GetID(c)
	if err != nil {
		return
	}

	err = h.categoryService.Delete(id)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSON(c, http.StatusOK, "successfully deleted")
}
