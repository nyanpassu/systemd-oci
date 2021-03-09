package meta

// Status .
type Status interface {
	SetStatus(id string, status string) error
	GetStatus(id string) (string, error)
}
