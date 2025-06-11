package handler

import (
	"domain/donation"
	"donationapi/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	GroupPathCampaign = "/campaigns"
)

func CampaignRouterGroup(engine *gin.Engine) *gin.RouterGroup {
	group := engine.Group(GroupPathCampaign)
	group.POST("", createNewCampaign)

	return group
}

func createNewCampaign(c *gin.Context) {
	uc := donation.NewCreateCampaignUseCase()
	input := donation.CreateCampaignInput{}
	err := c.BindJSON(&input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	campaign, err := uc.Execute(input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	res := response.Campaign{
		Metadata:    response.CreateMetadata(campaign),
		Title:       campaign.Title,
		Description: campaign.Description,
	}

	c.JSON(200, res)
}
