package main

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/disk"
)

func GetDiskInfo() {
	partitions, err := disk.Partitions(false)
	if err != nil {
		fmt.Println("Erreur lors de la récupération des partitions :", err)
		return
	}

	fmt.Println("Informations Disques")
	for _, p := range partitions {
		usage, err := disk.Usage(p.Mountpoint)
		if err != nil {
			continue
		}
		fmt.Printf("Partition : %s\n", p.Mountpoint)
		fmt.Printf("Total : %.2f GB\n", float64(usage.Total)/1e9)
		fmt.Printf("Utilisé : %.2f GB\n", float64(usage.Used)/1e9)
		fmt.Printf("Libre : %.2f GB\n", float64(usage.Free)/1e9)
		fmt.Printf("Utilisation : %.2f%%\n\n", usage.UsedPercent)
	}
}
