package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusConde int, message string)  {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusConde, errorResponse{message})
}

