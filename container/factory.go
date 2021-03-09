package container

import (
	"github.com/juju/errors"
	specsGo "github.com/opencontainers/runtime-spec/specs-go"

	"github.com/nyanpassu/systemd-oci/meta"
	"github.com/nyanpassu/systemd-oci/systemd"
)

// Factory .
type Factory interface {
	CreateContainer(id string, spec *specsGo.Spec) (Container, error)
	GetContainer(id string) (Container, error)
}

// NewFactory .
func NewFactory(m meta.Meta) (Factory, error) {
	return &factory{}, nil
}

type factory struct {
	meta    meta.Meta
	status  meta.Status
	systemd systemd.UnitManager
}

func (f *factory) CreateContainer(id string, spec *specsGo.Spec) (Container, error) {
	systemdUnitName := generateSystemdUnitName(id)
	if err := f.meta.CreateContainer(meta.Container{
		ID:              id,
		SystemdUnitName: systemdUnitName,
	}); err != nil {
		return nil, err
	}
	if err := f.systemd.Create(systemd.UnitFile{
		ContainerID: id,
		Name:        systemdUnitName,
		Args:        spec.Process.Args,
	}); err != nil {
		if errDel := f.meta.DeleteContainer(id); errDel != nil {
			return nil, errors.Wrap(err, errDel)
		}
		return nil, err
	}
	return newContainer(id, f.meta, f.status, f.systemd), nil
}

func (f *factory) GetContainer(id string) (Container, error) {
	_, err := f.meta.GetContainer(id)
	if err != nil {
		return nil, err
	}
	return newContainer(id, f.meta, f.status, f.systemd), nil
}

func generateSystemdUnitName(id string) string {
	return "systemdoci-" + id
}
