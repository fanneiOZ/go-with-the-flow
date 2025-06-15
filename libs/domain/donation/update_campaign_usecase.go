package donation

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
	return nil, nil
}
