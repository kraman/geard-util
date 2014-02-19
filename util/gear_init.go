package util

import (
	"github.com/smarterclayton/geard"
	"github.com/smarterclayton/go-systemd/account"
	"os"
)

func InitializeGear(name string) error {
	var accountService *account.AccountService
	var err error
	var gearId geard.Identifier

	if gearId, err = geard.NewIdentifier(name); err != nil {
		return err
	}

	if accountService, err = account.NewAccountService(); err != nil {
		return err
	}

	_, err = accountService.GetUserByName("gear-" + name)
	if err != account.ErrNoSuchUser {
		if err != nil {
			return err
		}
		return nil
	}

	os.RemoveAll(gearId.HomePath())
	gearId.HomePath()
	if _, err = accountService.CreateUser("gear-" + name, "gear-" + name, gearId.HomePath(), "/bin/bash"); err != nil {
		return err
	}

	return nil
}
