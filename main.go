package main

import (
	"fmt"
)

func main() {
	fmt.Println("Récupération des informations système...\n")

	GetOSInfo()


	GetCPUInfo()


	GetMemoryInfo()


	GetDiskInfo()

	
	GetNetworkInfo()
}
