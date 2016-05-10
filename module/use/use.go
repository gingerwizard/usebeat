package use

import (
	"github.com/elastic/beats/metricbeat/mb"
	"github.com/elastic/beats/libbeat/common"
)

func init() {
	if err := mb.Registry.AddModule("use", useModuleFactory); err != nil {
		panic(err)
	}
}

//Nic - what about the rawConfig?
var useModuleFactory = func(base mb.BaseModule) (mb.Module, error) {
	return UseModule{ name:"use", config: base.Config() }, nil
}

type UseModule struct {
	name      string
	config    mb.ModuleConfig
	rawConfig *common.Config
}

// Name returns the name of the Module.
func (m *UseModule) Name() string { return m.name }

// Config returns the ModuleConfig used to create the Module.
func (m *UseModule) Config() mb.ModuleConfig { return m.config }

// UnpackConfig unpacks the raw module config to the given object.
func (m *UseModule) UnpackConfig(to interface{}) error {
	return m.rawConfig.Unpack(to)
}