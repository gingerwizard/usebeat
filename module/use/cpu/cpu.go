package cpu

import (
	"github.com/elastic/beats/metricbeat/mb"
	"github.com/elastic/beats/libbeat/common"
)

func init() {
	if err := mb.Registry.AddMetricSet("use", "cpu", cpuMetricSetFactory); err != nil {
		panic(err)
	}
}

//Seems wrong to return the module from the BaseMetricSet
var cpuMetricSetFactory = func(base mb.BaseMetricSet) (mb.MetricSet, error) {
	return CPUMetricSet{name:"cpu",module:base.Module(),host:base.Host()},nil;
}


type CPUMetricSet struct {
	name   string
	module mb.Module
	host   string
}


func (b *CPUMetricSet) Name() string {
	return b.name
}

// Module returns the parent Module for the MetricSet.
func (b *CPUMetricSet) Module() mb.Module {
	return b.module
}

// Host returns the hostname or other module specific value that identifies a
// specific host or service instance from which to collect metrics.
func (b *CPUMetricSet) Host() string {
	return b.host
}


func (m *CPUMetricSet) Fetch() (event common.MapStr, err error) {
	return common.MapStr{"hello": "world"}
}
