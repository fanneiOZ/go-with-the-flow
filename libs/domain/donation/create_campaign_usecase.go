package donation

type CreateCampaignUseCase struct {
}

type CreateCampaignInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewCreateCampaignUseCase() *CreateCampaignUseCase {
	return &CreateCampaignUseCase{}
}

func (uc *CreateCampaignUseCase) Execute(input CreateCampaignInput) (*Campaign, error) {
	campaign, _ := CreateCampaign(input.Title, input.Description)

	return campaign, nil
}
