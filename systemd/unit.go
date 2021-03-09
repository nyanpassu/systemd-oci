package systemd

// Unit .
type Unit interface {
	Status() (string, error)
	Start() error
	Stop() error
	Delete() error
}
