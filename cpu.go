package main

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/cpu"
)


func GetCPUInfo() {
	info, err := cpu.Info()
	if err != nil {
		fmt.Println("Erreur lors de la récupération des infos CPU :", err)
		return
	}

	physicalCores, err := cpu.Counts(false)
	if err != nil {
		fmt.Println("Erreur lors de la récupération des cœurs physiques :", err)
		return
	}

	logicalCores, err := cpu.Counts(true)
	if err != nil {
		fmt.Println("Erreur lors de la récupération des threads CPU :", err)
		return
	}

	fmt.Println("Informations CPU")
	fmt.Printf("Modèle : %s\n", info[0].ModelName)
	fmt.Printf("Cœurs : %d\n", physicalCores)
	fmt.Printf("Threads : %d\n", logicalCores)
	fmt.Printf("Fréquence : %.2f GHz\n\n", info[0].Mhz/1000)
}