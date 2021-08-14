package campaign

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	Slug             string `json:"slug"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	campaignFrommater := CampaignFormatter{
		ID:               campaign.ID,
		UserID:           campaign.UserID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		ImageURL:         "",
		Slug:             campaign.Slug,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
	}

	if len(campaign.CampaignImages) > 0 {
		campaignFrommater.ImageURL = campaign.CampaignImages[0].FileName
	}

	return campaignFrommater
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	campaignsFromatter := []CampaignFormatter{}

	for _, campaign := range campaigns {
		campaignFromatter := FormatCampaign(campaign)
		campaignsFromatter = append(campaignsFromatter, campaignFromatter)
	}

	return campaignsFromatter
}
