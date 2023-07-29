// controller.go

package controller

import (
	"fmt"
	"sync"
	"tasktracker/model"
)

var mu sync.Mutex // Mutex untuk mengamankan akses konkurensi pada data tasks

// Interface untuk mengelola tugas
type TaskManager interface {
	AddTask(judul, deskripsi, prioritas, tanggalTenggat string) error
	UpdateTaskStatus(id int) error
	GetTasks() []model.Task
	GetTaskByID(id int) (model.Task, error)
}

// Implementasi TaskManager
type TaskController struct {
	tasks  []model.Task
	lastID int
}

// Implementasi method AddTask untuk TaskController
func (c *TaskController) AddTask(judul, deskripsi, prioritas, tanggalTenggat string) error {
	mu.Lock()
	defer mu.Unlock()

	c.lastID++
	task := model.Task{
		ID:             c.lastID,
		Judul:          judul,
		Deskripsi:      deskripsi,
		Prioritas:      prioritas,
		TanggalTenggat: tanggalTenggat,
		Status:         "To Do", // Status default adalah "To Do"
	}

	c.tasks = append(c.tasks, task)
	return nil
}

// Implementasi method DeleteTask untuk TaskController
func (c *TaskController) DeleteTask(id int) error {
	mu.Lock()
	defer mu.Unlock()

	for i, task := range c.tasks {
		if task.ID == id {
			// Menghapus tugas dari slice
			c.tasks = append(c.tasks[:i], c.tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Tugas dengan ID %d tidak ditemukan", id)
}

// Implementasi method UpdateTaskStatus untuk TaskController
func (c *TaskController) UpdateTaskStatus(id int) error {
	mu.Lock()
	defer mu.Unlock()

	for i, task := range c.tasks {
		if task.ID == id {
			c.tasks[i].Status = "Done"
			return nil
		}
	}
	return fmt.Errorf("Tugas dengan ID %d tidak ditemukan", id)
}

// Implementasi method GetTasks untuk TaskController
func (c *TaskController) GetTasks() []model.Task {
	mu.Lock()
	defer mu.Unlock()

	return c.tasks
}

// Implementasi method GetTaskByID untuk TaskController
func (c *TaskController) GetTaskByID(id int) (model.Task, error) {
	mu.Lock()
	defer mu.Unlock()

	for _, task := range c.tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return model.Task{}, fmt.Errorf("Tugas dengan ID %d tidak ditemukan", id)
}

// Fungsi untuk membuat instance TaskController sebagai implementasi TaskManager
func NewTaskController() *TaskController {
	return &TaskController{
		tasks:  make([]model.Task, 0),
		lastID: 0,
	}
}

// Fungsi ini akan dipanggil di main untuk menginisialisasi data awal
func init() {
	// Inisialisasi data awal
	taskController := NewTaskController()

	taskController.AddTask("Memperbaiki Masalah Jaringan", "Mengatasi masalah koneksi jaringan di lantai 2.", "Tinggi", "2023-07-31")
	taskController.AddTask("Menginstal Perangkat Lunak Baru", "Menginstal perangkat lunak terbaru di setiap komputer.", "Sedang", "2023-08-15")
}
