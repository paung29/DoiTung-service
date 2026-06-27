package pole

type PoleService interface {
	GetPoleByZone(year int, zoneNo int) (PolesByZoneResponse, error)
	GetPoleFilter(zoneId uint, poleNo *uint, harvestGradingFormDone *bool) (PoleFilterResponse, error)
}
