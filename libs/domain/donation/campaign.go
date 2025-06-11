package donation

import domain "shareddomain"

var _ domain.Entity = (*Campaign)(nil)

type Campaign struct {
	id          string
	Title       string
	Description string
	Active      bool
}

func (c *Campaign) Id() string {
	return c.id
}

func (c *Campaign) Type() domain.EntityType {
	return "campaign"
}

func NewCampaign(state Campaign) *Campaign {
	return &Campaign{
		id:          state.id,
		Title:       state.Title,
		Description: state.Description,
		Active:      state.Active,
	}
}

func CreateCampaign(title, description string) (*Campaign, error) {
	return &Campaign{
		Title:       title,
		Description: description,
		Active:      true,
	}, nil
}
