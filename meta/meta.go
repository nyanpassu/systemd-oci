package meta

import (
	"encoding/json"
	"os"

	"github.com/juju/errors"
)

// Meta .
type Meta interface {
	CreateContainer(Container) error
	GetContainer(id string) (Container, error)
	DeleteContainer(id string) error
}

// NewMeta .
func NewMeta(config Config) (Meta, error) {
	return &meta{}, nil
}

type meta struct {
}

func (m *meta) CreateContainer(container Container) error {
	dirPath := m.locateDir(container.ID)
	if _, err := os.Stat(dirPath); !os.IsNotExist(err) {
		if err != nil {
			return err
		}
		return errors.New("container exists")
	}
	if err := os.Mkdir(dirPath, 0644); err != nil {
		return err
	}
	filePath := m.locateFile(container.ID)
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		if err != nil {
			return err
		}
		return errors.New("container exists")
	}
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_EXCL|os.O_SYNC, 0644)
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(container, "", "  ")
	if err != nil {
		return err
	}
	_, err = f.WriteString(string(data))
	f.Close()
	return nil
}

func (m *meta) GetContainer(id string) (Container, error) {
	var (
		c Container
		content []byte
		err error
	)
	filePath := m.locateFile(id)
	if content, err = os.ReadFile(filePath); err != nil {
		return Container{}, err
	}
	if err = json.Unmarshal(content, &c); err != nil {
		return Container{}, err
	}
	return c, nil
}

func (m *meta) DeleteContainer(id string) error {
	if err := os.RemoveAll(m.locateDir(id)); err != nil {
		return err
	}
	if 
}

func (m *meta) locateDir(id string) string {
	return "/var/run/systemd-runc/" + id
}

func (m *meta) locateFile(id string) string {
	return "/var/run/systemd-runc/" + id + "/container.json"
}

func (m *meta) write(container Container) string {

}
