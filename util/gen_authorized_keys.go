package util

import (
	"bufio"
	"fmt"
	"github.com/smarterclayton/geard"
	"github.com/smarterclayton/geard/selinux"	
	"io"
	"os"
	"path"
	"path/filepath"
	"os/user"
	"strconv"
)

func GenerateAuthorizedKeys(name string) error {
	var err error
	var gearId geard.Identifier
	var sshKeys []string
	var destFile *os.File
	var srcFile *os.File
	var u *user.User
	
	if u, err = user.Lookup("gear-" + name) ; err != nil {
		return err
	}

	if gearId, err = geard.NewIdentifier(name); err != nil {
		return err
	}

	sshKeys, err = filepath.Glob(path.Join(gearId.SshAccessBasePath(), "*"))
	fmt.Println(path.Join(gearId.SshAccessBasePath(), "*"))

	os.MkdirAll(gearId.HomePath(), 0700)
	os.Mkdir(path.Join(gearId.HomePath(), ".ssh"), 0700)
	authKeysPath := path.Join(gearId.HomePath(), ".ssh", "authorized_keys")
	if _, err = os.Stat(authKeysPath); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	} else {
		return nil
	}

	if destFile, err = os.Create(authKeysPath); err != nil {
		return err
	}
	defer destFile.Close()
	w := bufio.NewWriter(destFile)

	for _, keyFile := range sshKeys {
		s, _ := os.Stat(keyFile)
		if s.IsDir() {
			continue
		}

		fmt.Println("k", keyFile)
		srcFile, err = os.Open(keyFile)
		defer srcFile.Close()
		fmt.Println(w.WriteString("command=\"" + geard.GearBasePath() + "/bin/geard-switchns\",no-port-forwarding,no-agent-forwarding,no-X11-forwarding "))
		fmt.Println(io.Copy(w, srcFile))
		fmt.Println(w.WriteString("\n"))
	}
	w.Flush()
	
	uid, _ := strconv.Atoi(u.Uid)
	gid, _ := strconv.Atoi(u.Gid)
	os.Chown(gearId.HomePath(), uid, gid)
	os.Chown(path.Join(gearId.HomePath(), ".ssh"), uid, gid)
	os.Chown(path.Join(gearId.HomePath(), ".ssh", "authorized_keys"), uid, gid)
	selinux.RestoreConRecursive(gearId.BaseHomePath())
	return nil
}
