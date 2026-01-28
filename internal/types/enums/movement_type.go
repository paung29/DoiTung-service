package enums

type MovementType string

const (
	MovementCarryOver MovementType = "CARRY_OVER"
	MovementIncoming  MovementType = "INCOMING"
	MovementIssued     MovementType = "ISSUED"
)