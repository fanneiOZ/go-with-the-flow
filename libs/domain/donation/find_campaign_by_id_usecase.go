package donation

type FindCampaignByIdUseCase struct {
	campaignRepo *CampaignRepo
}

func NewFindCampaignByIdUseCase(campaignRepo *CampaignRepo) *FindCampaignByIdUseCase {
	return &FindCampaignByIdUseCase{campaignRepo}
}

func (uc *FindCampaignByIdUseCase) Execute(id int) (*Campaign, error) {
	campaign, err := uc.campaignRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	if campaign == nil {
		return nil, ErrCampaignNotFound
	}

	return campaign, nil
}
