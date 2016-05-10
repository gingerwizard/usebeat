package main

import (
	"os"
	"github.com/elastic/beats/metricbeat/beater"

	_"github.com/elastic/beats/metricbeat/include"

	_"github.com/gingerwizard/usebeat/module/use"
	_"github.com/gingerwizard/usebeat/module/use/cpu"

	"github.com/elastic/beats/libbeat/beat"
)

var Name = "usebeat"

func main() {
	if err := beat.Run(Name, "", beater.New()); err != nil {
		os.Exit(1)
	}
}