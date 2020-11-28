package handler

import (
	"github.com/andrejtad/final"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createSpecification(c *gin.Context)  {

	dataOwnerId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	var input final.Specification
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Specification.Create(dataOwnerId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":id,
	})
}

type getAllSpecificationsResponse struct {
	Data []final.Specification `json:"data"`
}

func (h *Handler) getAllSpecifications(c *gin.Context)  {

	datasetId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	datasets, err := h.services.Specification.GetAll(datasetId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllSpecificationsResponse{
		Data: datasets,
	})

}

func (h *Handler) getSpecificationById(c *gin.Context)  {

	specificationId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	dataset, err := h.services.Specification.GetById(specificationId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dataset)
}

func (h *Handler) updateSpecification(c *gin.Context)  {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var dateset final.UpdateSpecificationInput
	if err := c.BindJSON(&dateset); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Specification.Update(id, dateset); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteSpecification(c *gin.Context)  {

	specificationId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	err = h.services.Specification.Delete(specificationId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

