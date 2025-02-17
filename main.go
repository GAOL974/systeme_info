package main

import (
	"fmt"
	"systeme_info/cpuinfo"
	"systeme_info/diskinfo"
	"systeme_info/memoryinfo"
	"systeme_info/networkinfo"
	"systeme_info/osinfo"
)

func main() {
	fmt.Println("Récupération des informations système...")

	osinfo.GetOSInfo()

	cpuinfo.GetCPUInfo()

	memoryinfo.GetMemoryInfo()

	diskinfo.GetDiskInfo()

	networkinfo.GetNetworkInfo()
}
