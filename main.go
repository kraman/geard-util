package main

import (
	"fmt"
	"os"
	"github.com/docopt/docopt.go"
	"github.com/kraman/geard-util/util"
)

func main() {
	usage := `Geard Utilities

Usage:
	geard-util gen-authorized-keys <gear name>
`
	var arguments map[string]interface{}
	var err error
	if arguments, err = docopt.Parse(usage, nil, true, "GearD Utilities", false); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	gearName := arguments["<gear name>"].(string)

	switch {
	case arguments["gen-authorized-keys"]:
		err = util.GenerateAuthorizedKeys(gearName)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(2)		
	}
}
