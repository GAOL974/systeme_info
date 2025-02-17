package networkinfo

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/net"
)

func GetNetworkInfo() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Erreur lors de la récupération des interfaces réseau :", err)
		return
	}

	fmt.Println("Informations Réseau")
	for _, i := range interfaces {
		fmt.Printf("Interface : %s\n", i.Name)
		for _, addr := range i.Addrs {
			fmt.Printf("  - Adresse IP : %s\n", addr.Addr)
		}
		fmt.Println()
	}
}
