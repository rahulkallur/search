package controllers

import (
	model "Search/models"
	services "Search/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SearchRequestController interface {
	SearchRequestMapper(ctx *gin.Context)
}

type searchRequestcontroller struct {
	repo services.SearchRepository
}

func NewSearchRequestController(repo services.SearchRepository) SearchRequestController {
	return &searchRequestcontroller{
		repo: repo,
	}
}

// SearchRequestMapper handles incoming hotel search requests.
// It binds the JSON request body to the HotelSearchRequest model, validates it,
// and delegates processing to the repository layer.
// Parameters:
//   - ctx: The Gin context that holds the HTTP request and response objects.
func (c *searchRequestcontroller) SearchRequestMapper(ctx *gin.Context) {
	rq := model.HotelSearchRequest{}
	if err := ctx.BindJSON(&rq); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	// Pass the mapped request to the repository for further processing.
	request := c.repo.SearchRequestMapper(rq)
	ctx.JSON(http.StatusOK, request)
}
