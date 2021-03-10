package command

import (
	"os"

	"github.com/juju/errors"
	specsGo "github.com/opencontainers/runtime-spec/specs-go"
	cli "github.com/urfave/cli/v2"

	"github.com/nyanpassu/systemd-oci/common"
	"github.com/nyanpassu/systemd-oci/container"
	"github.com/nyanpassu/systemd-oci/meta"
	"github.com/nyanpassu/systemd-oci/spec"
	"github.com/nyanpassu/systemd-oci/systemd"
)

func getContainerFactory() (container.Factory, error) {
	var (
		err error
		m   meta.Meta
		um  systemd.UnitManager
		f   container.Factory
	)
	if m, err = meta.NewMeta(meta.Config{}); err != nil {
		return nil, err
	}
	if um, err = systemd.NewUnitManager(); err != nil {
		return nil, err
	}
	if f, err = container.NewFactory(m, meta.NewStatus(um), um); err != nil {
		return nil, err
	}
	return f, nil
}

func createContainer(context *cli.Context) (container.Container, error) {
	return getOrCreateContainer(context, true)
}

func getContainer(context *cli.Context) (container.Container, error) {
	return getOrCreateContainer(context, false)
}

func getOrCreateContainer(context *cli.Context, create bool) (container.Container, error) {
	id := context.Args().First()
	if id == "" {
		return nil, errors.New("container id is empty")
	}
	var (
		f   container.Factory
		err error
	)
	if f, err = getContainerFactory(); err != nil {
		return nil, err
	}
	if create {
		var spec *specsGo.Spec
		if spec, err = setupSpec(context); err != nil {
			return nil, err
		}
		return f.CreateContainer(id, spec)
	}
	return f.GetContainer(id)
}

// setupSpec performs initial setup based on the cli.Context for the container
func setupSpec(context *cli.Context) (*specsGo.Spec, error) {
	bundle := context.String("bundle")
	if bundle != "" {
		if err := os.Chdir(bundle); err != nil {
			return nil, err
		}
	}
	spec, err := spec.LoadSpec(common.SpecConfig)
	if err != nil {
		return nil, err
	}
	return spec, nil
}
