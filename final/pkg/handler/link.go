package handler

import (
	"github.com/andrejtad/final"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type getAllLinkResponse struct {
	Data []final.Link `json:"data"`
}

func (h *Handler) getAllLink(c *gin.Context)  {

	childId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	links, err := h.services.Link.GetAll(childId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllLinkResponse{
		Data: links,
	})
}
