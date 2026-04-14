package zone

type CreateZoneRequest struct {
	Year uint `json:"year" validate:"required,min=4"`
	Name string `json:"name" validate:"required,min=6"`
}

type CreateZoneResponse struct {
	Message string `json:"message"`
}

type ZoneResponse struct {
	ZoneID uint `json:"zoneId"`
	ZoneName string  `json:"zoneName"`
}

type GetAllZoneResponse struct {
	Zones []ZoneResponse `json:"zones"`
}

type GetAllZoneForm struct {
	Year int `json:"year" validate:"required,min=4"`
}

