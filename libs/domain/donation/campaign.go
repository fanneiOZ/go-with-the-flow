package donation

import (
	"shareddomain/entity"
	"strconv"
	"time"
)

var _ entity.Entity = (*Campaign)(nil)

type Campaign struct {
	id          int
	Title       string
	Description string
	Active      bool
	createdAt   time.Time
	updatedAt   time.Time
	*entity.Version
}

func (c *Campaign) Id() string {
	return strconv.Itoa(c.id)
}

func (c *Campaign) Type() entity.Type {
	return "campaign"
}

func NewCampaign(state Campaign) *Campaign {
	return &Campaign{
		id:          state.id,
		Title:       state.Title,
		Description: state.Description,
		Active:      state.Active,
		Version:     state.Version,
		createdAt:   state.createdAt,
		updatedAt:   state.updatedAt,
	}
}

func CreateCampaign(title, description string) (*Campaign, error) {
	return &Campaign{
		Title:       title,
		Description: description,
		Active:      true,
		Version:     entity.CreateNewVersion(),
	}, nil
}
