package systemd

// UnitManager .
type UnitManager interface {
	GetUnit(string) (Unit, bool, error)
	Create(UnitFile) error
}
