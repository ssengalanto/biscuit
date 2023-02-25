package app

import (
	"os"

	"github.com/ssengalanto/biscuit/pkg/constants"
)

type defaults struct {
	env string
	cfg string
}

func getDefaults() defaults {
	var d defaults

	d.env = os.Getenv(constants.AppEnv)
	d.cfg = os.Getenv(constants.ConfigType)

	if d.env == "" {
		d.env = *flEnv
	}

	if d.cfg == "" {
		d.cfg = *flCfg
	}

	return d
}
