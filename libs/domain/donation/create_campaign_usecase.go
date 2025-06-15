package donation

type CreateCampaignUseCase struct {
	campaignRepo *CampaignRepo
}

type CreateCampaignInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewCreateCampaignUseCase(campaignRepo *CampaignRepo) *CreateCampaignUseCase {
	return &CreateCampaignUseCase{campaignRepo}
}

func (uc *CreateCampaignUseCase) Execute(input CreateCampaignInput) (*Campaign, error) {
	var err error
	var campaign, _ = CreateCampaign(input.Title, input.Description)
	campaign, err = uc.campaignRepo.Save(campaign)
	if err != nil {
		return nil, err
	}

	return campaign, nil
}
