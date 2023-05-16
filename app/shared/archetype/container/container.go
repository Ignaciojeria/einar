package container

import (
	"errors"
	"os"

	"github.com/rs/zerolog/log"
)

var InstallationsContainer map[string]DependencyContainer = map[string]DependencyContainer{}
var ComponentsContainer map[string]DependencyContainer = map[string]DependencyContainer{}
var HTTPServerContainer DependencyContainer = DependencyContainer{}

type DependencyContainer struct {
	InjectionProps InjectionProps
	LoadDependency LoadDependency
	isPresent      bool
}

type InjectionProps struct {
	DependencyID string // name of injected dependency should be unique and required
	Paralel      bool   // if true injected dependency should be executed as a go routine
}

type LoadDependency func() error

func InjectComponent(dependency LoadDependency, props InjectionProps) error {
	if props.DependencyID == "" {
		err := errors.New("container injector error on InjectionProps. DependencyID can't be empty")
		log.Error().Err(err).Send()
		return err
	}
	if ComponentsContainer[props.DependencyID].isPresent {
		err := errors.New("container injector error. Next dependency already exits : " + props.DependencyID)
		log.Error().Err(err).Send()
		return err
	}
	ComponentsContainer[props.DependencyID] = DependencyContainer{LoadDependency: dependency, InjectionProps: props, isPresent: true}
	return nil
}

func InjectInstallation(dependency LoadDependency, props InjectionProps) error {
	if props.DependencyID == "" {
		err := errors.New("container injector error on InjectionProps. DependencyID can't be empty")
		log.Error().Err(err).Send()
		return err
	}
	if InstallationsContainer[props.DependencyID].isPresent {
		err := errors.New("container injector error. Next dependency already exits : " + props.DependencyID)
		log.Error().Err(err).Send()
		return err
	}
	InstallationsContainer[props.DependencyID] = DependencyContainer{LoadDependency: dependency, InjectionProps: props, isPresent: true}
	return nil
}

func InjectHTTPServer(dependency LoadDependency, props InjectionProps) error {
	HTTPServerContainer = DependencyContainer{LoadDependency: dependency, InjectionProps: props, isPresent: true}
	return nil
}

type IExit func()

var Exit IExit = func() {
	os.Exit(0)
}
