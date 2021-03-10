package main

import (
	"log"

	// "github.com/nyanpassu/systemd-oci/utils"

	"github.com/nyanpassu/systemd-oci/systemd/systemctl"
)

func main() {
	// exists, err := utils.FileExists("/var/run/systemd-runc")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// log.Println(exists)

	unitName := generateSystemdUnitName("mybusybox")

	if err := systemctl.Enable(unitName); err != nil {
		log.Println("failed")
		log.Fatalln(err)
	}
	log.Println("success")
}

func generateSystemdUnitName(id string) string {
	return "systemdoci-" + id
}
