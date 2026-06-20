package zone

type ZoneService interface {
	CreateZone(form CreateZoneRequest) (CreateZoneResponse, error)
	GetAllZone(yearID uint) (GetAllZoneResponse, error)
	GetZoneManagementTable(yearID uint) (GetZoneManagementTableResponse, error)
	UpdateZoneName(form UpdateZoneName) (UpdateZoneNameResponse, error)
}
