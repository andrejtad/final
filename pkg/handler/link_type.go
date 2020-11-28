package handler

import (
	"github.com/andrejtad/final"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getAllLinkTypeResponse struct {
	Data []final.LinkType `json:"data"`
}

func (h *Handler) getAllLinkTypes(c *gin.Context)  {

	linkTypes, err := h.services.LinkType.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllLinkTypeResponse{
		Data: linkTypes,
	})
}
