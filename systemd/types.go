package systemd

// UnitFile .
type UnitFile struct {
	ContainerID string
	Name        string
	Args        []string
}

// Status .
type Status int
