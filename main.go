package main

import (
	"fmt"
	"github.com/docopt/docopt.go"
	"github.com/kraman/geard-util/util"
)

func main() {
	usage := `Geard Utilities

Usage:
	geard-util gear-init <gear name>
	geard-util gen-authorized-keys <gear name>
`
	var arguments map[string]interface{}
	var err error
	if arguments, err = docopt.Parse(usage, nil, true, "GearD Utilities", false); err != nil {
		fmt.Println(err)
	}

	gearName := arguments["<gear name>"].(string)

	switch {
	case arguments["gear-init"]:
		err = util.InitializeGear(gearName)
	case arguments["gen-authorized-keys"]:
		err = util.GenerateAuthorizedKeys(gearName)
	}

	if err != nil {
		fmt.Println(err)
	}
}
