package donation_test

import (
	"context"
	"domain/donation"
	"sharedinfra/db/postgres"
	"sync"
	"testing"
)

var ctx = context.Background()

func TestCampaignRepo(t *testing.T) {
	t.Run("", func(t *testing.T) {
		dbContext, err := postgres.NewContext(ctx)
		if err != nil {
			t.Error(err)
		}

		repo := donation.NewCampaignRepo(dbContext)
		campaign, _ := donation.CreateCampaign("new-title", "new-desc")

		campaign, err = repo.Save(campaign)
		if err != nil {
			t.Error(err)
		}

		var c = make(chan *donation.Campaign, 1)
		var wg sync.WaitGroup
		wg.Add(2)

		go func(campaign *donation.Campaign) {
			defer wg.Done()
			campaign.Title = "title-v2"
			newCampaign, err := repo.Save(campaign)
			if err != nil {
				t.Logf("gr1: %v", err)
				return
			}
			campaign = newCampaign
			t.Logf("gr1: %v, address: %v", &campaign, campaign)
			c <- campaign
		}(campaign)

		go func(campaign *donation.Campaign) {
			defer wg.Done()
			campaign.Title = "title-v3"
			newCampaign, err := repo.Save(campaign)
			if err != nil {
				t.Logf("gr2: %v", err)
				return
			}
			campaign = newCampaign
			t.Logf("gr2: %v, address: %v", &campaign, campaign)
			c <- campaign
		}(campaign)

		wg.Wait()
		campaign = <-c
		t.Logf("after wait %v, %v", &campaign, campaign)
		campaign.Title = "title-v4"
		campaign, err = repo.Save(campaign)
		if err != nil {
			t.Error(err)
		}
		t.Logf("after new update %v, %v", &campaign, campaign)

	})
}
