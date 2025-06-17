package donation

import "shareddomain/entity"

type UpdateCampaignUseCase struct {
	campaignRepo *CampaignRepo
}

type UpdateCampaignInput struct {
	Id          int
	Title       string `json:"title"`
	Description string `json:"description"`
	Version     uint   `json:"version"`
}

func NewUpdateCampaignUseCase(campaignRepo *CampaignRepo) *UpdateCampaignUseCase {
	return &UpdateCampaignUseCase{campaignRepo: campaignRepo}
}

func (uc *UpdateCampaignUseCase) Execute(input UpdateCampaignInput) (*Campaign, error) {
	campaign, err := uc.campaignRepo.FindById(input.Id)
	if err != nil {
		return nil, err
	}

	campaign.Title = input.Title
	campaign.Description = input.Description
	campaign.Version = entity.NewVersion(input.Version)
	campaign, err = uc.campaignRepo.Save(campaign)
	if err != nil {
		return nil, err
	}

	return campaign, nil
}
