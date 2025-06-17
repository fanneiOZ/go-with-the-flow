package donation_test

import (
	"domain/donation"
	"testing"
)

func TestCampaign(t *testing.T) {

	t.Run("should create a new Campaign", func(t *testing.T) {
		campaign, _ := donation.CreateCampaign("title", "description")

		if campaign.Id() != "0" {
			t.Error("Id should be 0")
		}

		if campaign.Type() != "campaign" {
			t.Errorf("expect value of type to be campaign, received: %s", campaign.Type())
		}

		if campaign.EntityVersion() != 1 {
			t.Errorf("expect value of entity version to be 1, received: %d", campaign.EntityVersion())
		}

		if campaign.Title != "title" {
			t.Errorf("expect value of title to be title, received: %s", campaign.Type())
		}

		if campaign.Description != "description" {
			t.Errorf("expect value of description to be description, received: %s", campaign.Type())
		}
	})
}
