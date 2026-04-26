package pole

type PoleService interface {
	GetPoleByZone(year int, zoneNo int) (PolesByZoneResponse, error)
}
