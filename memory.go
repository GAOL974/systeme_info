package main

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/mem"
)

func GetMemoryInfo() {
	v, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Erreur lors de la récupération des infos mémoire :", err)
		return
	}

	fmt.Println("Informations Mémoire")
	fmt.Printf("Total : %.2f GB\n", float64(v.Total)/1e9)
	fmt.Printf("Utilisé : %.2f GB\n", float64(v.Used)/1e9)
	fmt.Printf("Disponible : %.2f GB\n", float64(v.Available)/1e9)
	fmt.Printf("Utilisation : %.2f%%\n\n", v.UsedPercent)
}
