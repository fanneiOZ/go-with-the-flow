package handler

import (
	"context"
	"domain/donation"
	"donationapi/response"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"sharedinfra/db/postgres"
	"strconv"
)

const (
	GroupPathCampaign = "/campaigns"
)

func CampaignRouterGroup(engine *gin.Engine) *gin.RouterGroup {
	group := engine.Group(GroupPathCampaign)
	group.POST("", createNewCampaign)
	group.GET(":id", findById)

	return group
}

func createNewCampaign(c *gin.Context) {
	pgContext, err := postgres.NewContext(context.Background())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	repo := donation.NewCampaignRepo(pgContext)
	uc := donation.NewCreateCampaignUseCase(repo)

	var input donation.CreateCampaignInput
	err = c.BindJSON(&input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	campaign, err := uc.Execute(input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := response.Campaign{
		Metadata:    response.CreateMetadata(campaign),
		Title:       campaign.Title,
		Description: campaign.Description,
	}

	c.JSON(200, res)
}

func findById(c *gin.Context) {
	var err error
	pgContext, err := postgres.NewContext(context.Background())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	repo := donation.NewCampaignRepo(pgContext)
	uc := donation.NewFindCampaignByIdUseCase(repo)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	campaign, err := uc.Execute(id)
	if err != nil {
		var errStatus int
		switch {
		case errors.Is(err, donation.ErrCampaignNotFound):
			errStatus = http.StatusNotFound
			break
		default:
			errStatus = http.StatusBadRequest
		}

		c.AbortWithStatusJSON(errStatus, gin.H{"error": err.Error()})
		return
	}

	res := response.Campaign{
		Metadata:    response.CreateMetadata(campaign),
		Title:       campaign.Title,
		Description: campaign.Description,
	}

	c.JSON(200, res)
}
