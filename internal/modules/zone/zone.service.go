package zone

type ZoneService interface {
	CreateZone(form CreateZoneRequest) (CreateZoneResponse, error)
}
