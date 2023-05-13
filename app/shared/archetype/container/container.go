package container

import (
	"errors"
	"os"

	"github.com/rs/zerolog/log"
)

var Dependencies map[string]DependencyContainer = map[string]DependencyContainer{}

type DependencyContainer struct {
	InjectionProps InjectionProps
	Dependency     Dependency
	isPresent      bool
}

type InjectionProps struct {
	DependencyID string // name of injected dependency should be unique and required
	Paralel      bool   // if true injected dependency should be executed as a go routine
}

type Dependency func() error

func Inject(dependency Dependency, props InjectionProps) error {
	if props.DependencyID == "" {
		err := errors.New("container injector error on InjectionProps. DependencyID can't be empty")
		log.Error().Err(err).Send()
		return err
	}
	if Dependencies[props.DependencyID].isPresent {
		err := errors.New("container injector error. Next dependency already exits : " + props.DependencyID)
		log.Error().Err(err).Send()
		return err
	}
	Dependencies[props.DependencyID] = DependencyContainer{Dependency: dependency, InjectionProps: props, isPresent: true}
	return nil
}

type IExit func()

var Exit IExit = func() {
	os.Exit(0)
}
