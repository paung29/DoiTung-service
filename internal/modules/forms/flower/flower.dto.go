package flower

type FlowerFormRequest struct {
	ClusterId    uint   `json:"clusterId" validate:"required"`
	TotalFlowers *uint  `json:"totalFlowers" validate:"required,number"`
	Condition    string `json:"condition" validate:"required,oneof=GOOD INSECT ROTTEN"`
}

type FlowerFormResponse struct {
	Message string `json:"message"`
}

type FlowerFormDetails struct {
	No             int    `json:"no,omitempty"`
	ClusterId      uint   `json:"clusterId"`
	Location       string `json:"location"`
	PoleNo         int    `json:"poleNo"`
	ClusterNo      int    `json:"clusterNo"`
	TotalFlowers   uint   `json:"totalFlowers"`
	Condition      string `json:"condition"`
	FlowerFormDone bool   `json:"flowerFormDone,omitempty"`
	RecordedBy     string `json:"recordedBy,omitempty"`
	Date           string `json:"date,omitempty"`
}

type FlowerFormHistoriesResponse struct {
	FlowerFormHistories []FlowerFormHistory `json:"flowerFormHistories"`
}

type FlowerFormHistory struct {
	ClusterId    uint   `json:"clusterId"`
	Location     string `json:"location"`
	PoleNo       int    `json:"poleNo"`
	ClusterNo    int    `json:"clusterNo"`
	ProgressDone uint   `json:"progressDone"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}

type FlowerFormLists struct {
	FlowerForms []FlowerFormDetails `json:"flowerForms"`
}
