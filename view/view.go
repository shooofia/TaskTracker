package view

import (
	"fmt"
	"tasktracker/model"
)

type View struct{}

func (v View) DisplayTasks(tasks []model.Task) {
	fmt.Println("Daftar Tugas:")
	for _, task := range tasks {
		fmt.Printf("ID: %d\n", task.ID)
		fmt.Printf("Judul: %s\n", task.Judul)
		fmt.Printf("Deskripsi: %s\n", task.Deskripsi)
		fmt.Printf("Prioritas: %s\n", task.Prioritas)
		fmt.Printf("Tanggal Tenggat: %s\n", task.TanggalTenggat)
		fmt.Printf("Status: %s\n\n", task.Status)
	}
}

func (v View) DisplayTaskDetail(task model.Task) {
	fmt.Println("Detail Tugas:")
	fmt.Printf("ID: %d\n", task.ID)
	fmt.Printf("Judul: %s\n", task.Judul)
	fmt.Printf("Deskripsi: %s\n", task.Deskripsi)
	fmt.Printf("Prioritas: %s\n", task.Prioritas)
	fmt.Printf("Tanggal Tenggat: %s\n", task.TanggalTenggat)
	fmt.Printf("Status: %s\n\n", task.Status)
}

func (v View) DisplayMessage(message string) {
	fmt.Println(message)
}
