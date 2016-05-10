package use

import (
	"github.com/elastic/beats/metricbeat/mb"
)

func init() {
	if err := mb.Registry.AddModule("use", New); err != nil {
		panic(err)
	}
}

func New(base mb.BaseModule) (mb.Module, error) {
	return &UseModule {
		BaseModule:base,
	}, nil
}

type UseModule struct {
	mb.BaseModule
}
