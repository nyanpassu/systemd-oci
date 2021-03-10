package main

import (
	"os"
	"os/exec"

	"github.com/juju/errors"
	log "github.com/sirupsen/logrus"

	"github.com/nyanpassu/systemd-oci/systemd"
)

func main() {
	id := os.Args[1]
	cmdFile := os.Args[2]
	arg := os.Args[3:]

	cmd := exec.Command(cmdFile, arg...)
	if err := cmd.Start(); err != nil {
		log.Fatalln(err)
	}

	pid := cmd.Process.Pid

	if err := systemd.WritePid(id, pid); err != nil {
		if errKill := cmd.Process.Kill(); errKill != nil {
			log.Fatalln(errors.Wrap(err, errKill))
		}
		log.Fatalln(err)
	}

	status, err := cmd.Process.Wait()
	if err != nil {
		log.Fatalln(err)
	}
	os.Exit(status.ExitCode())

}
