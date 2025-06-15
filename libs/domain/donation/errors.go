package donation

import "errors"

var (
	ErrCampaignNotFound     = errors.New("campaign not found")
	ErrRequiredFieldMissing = errors.New("required field missing")
	ErrInvalidDonateAmount  = errors.New("donate amount must be greater than zero")
)
