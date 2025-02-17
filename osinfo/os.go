package osinfo

import (
	"fmt"
	"runtime"

	"github.com/shirou/gopsutil/v3/host"
)

func GetOSInfo() {
	info, err := host.Info()
	if err != nil {
		fmt.Println("Erreur lors de la récupération des infos de l'OS :", err)
		return
	}

	arch := "Inconnu"
	if runtime.GOARCH == "amd64" {
		arch = "64 bits"
	} else if runtime.GOARCH == "386" {
		arch = "32 bits"
	}

	fmt.Println("Informations Système")
	fmt.Printf("OS : %s\n", info.Platform)
	fmt.Printf("Architecture : %s\n", arch)
	fmt.Printf("Uptime : %d secondes\n", info.Uptime)
	fmt.Printf("Hôte : %s\n\n", info.Hostname)
}
