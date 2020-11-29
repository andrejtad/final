package handler

import (
	"github.com/andrejtad/final"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createDataset(c *gin.Context)  {

	dataOwnerId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	var input final.Dataset
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Dataset.Create(dataOwnerId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":id,
	})
}

type getAllDatasetsResponse struct {
	Data []final.Dataset `json:"data"`
}

func (h *Handler) getAllDatasets(c *gin.Context)  {

	dataOwnerId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	datasets, err := h.services.Dataset.GetAll(dataOwnerId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllDatasetsResponse{
		Data: datasets,
	})

}

func (h *Handler) getDatasetById(c *gin.Context)  {

	datasetId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	dataset, err := h.services.Dataset.GetById(datasetId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dataset)
}

func (h *Handler) updateDataset(c *gin.Context)  {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var dateset final.UpdateDatasetInput
	if err := c.BindJSON(&dateset); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Dataset.Update(id, dateset); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteDataset(c *gin.Context)  {

	datasetId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	err = h.services.Dataset.Delete(datasetId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

