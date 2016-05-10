package cpu

import (
	"github.com/elastic/beats/metricbeat/mb"
	"github.com/elastic/beats/libbeat/common"
)

func init() {
	if err := mb.Registry.AddMetricSet("use", "cpu", New); err != nil {
		panic(err)
	}
}

func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	return &CPUMetricSet{
		BaseMetricSet: base,
	},nil

}

type CPUMetricSet struct {
	mb.BaseMetricSet
}


func (m *CPUMetricSet) Fetch() (event common.MapStr, err error) {
	return common.MapStr{"hello": "world"},nil
}