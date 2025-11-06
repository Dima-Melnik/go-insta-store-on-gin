package handlers

import (
	"net/http"

	"github.com/Dima-Melnik/go-insta-store-on-gin/internal/models"
	"github.com/Dima-Melnik/go-insta-store-on-gin/internal/models/dto"
	"github.com/Dima-Melnik/go-insta-store-on-gin/internal/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllAttributes(c *gin.Context) {
	attributes, err := h.attributeService.GetAll()
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSON(c, http.StatusOK, attributes)
}

func (h *Handler) GetAttributeByID(c *gin.Context) {
	id, err := utils.GetID(c)
	if err != nil {
		return
	}

	attribute, err := h.attributeService.GetByID(id)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSON(c, http.StatusOK, attribute)
}

func (h *Handler) GetAttributeByName(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		utils.SendError(c, http.StatusBadRequest, map[string]string{
			"error": "missing 'name' query paeameter",
		})
	}

	category, err := h.attributeService.GetByName(name)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSON(c, http.StatusOK, category)
}

func (h *Handler) CreateAttribute(c *gin.Context) {
	var createAttribute dto.CreateAttribute

	if !utils.BindJSON(c, createAttribute) {
		return
	}

	attribute := models.Attribute{
		Name:     createAttribute.Name,
		Value:    createAttribute.Value,
		Products: []models.Product{},
	}

	createdProduct, err := h.attributeService.Create(attribute)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSON(c, http.StatusCreated, dto.ToAttributeResponse(createdProduct))
}

func (h *Handler) UpdateAttribute(c *gin.Context) {
	var updateAttribute dto.UpdateAttribute

	if !utils.BindJSON(c, updateAttribute) {
		return
	}

	id, err := utils.GetID(c)
	if err != nil {
		return
	}

	attribute := models.Attribute{
		Name:     *updateAttribute.Name,
		Value:    *updateAttribute.Value,
		Products: []models.Product{},
	}

	updatedAttribute, err := h.attributeService.Update(id, attribute)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSON(c, http.StatusOK, dto.ToAttributeResponse(updatedAttribute))
}

func (h *Handler) DeleteAttribute(c *gin.Context) {
	id, err := utils.GetID(c)
	if err != nil {
		return
	}

	err = h.attributeService.Delete(id)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSON(c, http.StatusOK, "successfully deleted")
}
