package handler

import (
	"github.com/andrejtad/final/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler{
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine  {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	auth := router.Group("/api/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.UserIdentity)
	{
		dataOwners := api.Group("/dataowners")
		{
			dataOwners.POST("/", h.createDataOwner)
			dataOwners.GET("/", h.getAllDataOwners)
			dataOwners.PUT("/:id", h.updateDataOwner)
			dataOwners.GET("/:id", h.getDataOwnerById)
			dataOwners.DELETE("/:id", h.DeleteDataOwner)

			dataSets := dataOwners.Group(":id/datasets")
			{
				dataSets.POST("/", h.createDataset)
				dataSets.GET("/", h.getAllDatasets)
			}
		}
		dataSets := api.Group("/datasets")
		{
			dataSets.PUT("/:id", h.updateDataset)
			dataSets.GET("/:id", h.getDatasetById)
			dataSets.DELETE("/:id", h.deleteDataset)

			specifications := dataSets.Group(":id/specifications")
			{
				specifications.POST("/", h.createSpecification)
				specifications.GET("/", h.getAllSpecifications)
			}
		}
		specifications := api.Group("/specifications")
		{
			specifications.PUT("/:id", h.updateSpecification)
			specifications.GET("/:id", h.getSpecificationById)
			specifications.DELETE("/:id", h.deleteSpecification)
		}
		tags := api.Group("/tags")
		{
			tags.POST("/", h.createTag)
			tags.GET("/", h.getAllTags)
			tags.PUT("/:id", h.updateTag)
			tags.GET("/:id", h.getTagById)
			tags.DELETE("/:id", h.DeleteTag)
		}
		linkType := api.Group("/linktype")
		{
			linkType.GET("/", h.getAllLinkTypes)
		}
	}
	return router
}
