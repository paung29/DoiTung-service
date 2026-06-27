package pole

type PolesByZoneResponse struct {
	Poles []PoleResponse `json:"poles"`
}

type PoleResponse struct {
	PoleId                 uint   `json:"poleId"`
	ZoneId                 uint   `json:"zoneId"`
	Location               string `json:"location"`
	PoleNo                 uint   `json:"poleNo"`
	HarvestGradingFormDone bool   `json:"harvestGradingFormDone"`
	CreatedAt              string `json:"createdAt"`
	UpdatedAt              string `json:"updatedAt"`
}

type PoleFilterResponse struct {
	Poles []PoleResponse `json:"poles"`
}
