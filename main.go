package main

//go:generate swag init

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"tasktracker/controller"
	"tasktracker/model"

	"github.com/labstack/echo/v4"

	echoSwagger "github.com/swaggo/echo-swagger"
)

var tasks []model.Task

func getTasksHandler(c echo.Context) error {
	tasks := taskController.GetTasks()
	return c.JSON(http.StatusOK, tasks)
}

func addTaskHandler(c echo.Context) error {
	var task model.Task
	err := c.Bind(&task)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
	}
	tasks = append(tasks, task)

	taskController.AddTask(task.Judul, task.Deskripsi, task.Prioritas, task.TanggalTenggat)
	return c.JSON(http.StatusCreated, map[string]string{"message": "Tugas baru berhasil ditambahkan."})
}

func deleteTaskHandler(c echo.Context) error {
	taskID := c.Param("id")
	id, err := strconv.Atoi(taskID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid task ID"})
	}

	err = taskController.DeleteTask(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to delete task"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Task successfully deleted"})
}

func updateTaskHandler(c echo.Context) error {
	taskID := c.Param("id")
	id, err := strconv.Atoi(taskID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid task ID"})
	}

	var updatedTask model.Task
	err = c.Bind(&updatedTask)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
	}

	var existingTask *model.Task
	for i, task := range tasks {
		if task.ID == id {
			existingTask = &tasks[i]
			break
		}
	}

	if existingTask == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Task not found"})
	}

	// Update tugas dengan data yang baru
	existingTask.Judul = updatedTask.Judul
	existingTask.Deskripsi = updatedTask.Deskripsi
	existingTask.Prioritas = updatedTask.Prioritas
	existingTask.TanggalTenggat = updatedTask.TanggalTenggat

	return c.JSON(http.StatusOK, map[string]string{"message": "Task updated successfully"})
}

var taskController *controller.TaskController

func main() {
	taskController = controller.NewTaskController()

	// Inisialisasi data awal menggunakan taskController.AddTask()
	taskController.AddTask("Memperbaiki Masalah Jaringan", "Mengatasi masalah koneksi jaringan di lantai 2.", "Tinggi", "2023-07-31")
	taskController.AddTask("Menginstal Perangkat Lunak Baru", "Menginstal perangkat lunak terbaru di setiap komputer.", "Sedang", "2023-08-15")

	e := echo.New()

	// Endpoint untuk mendapatkan daftar tugas
	e.GET("/tasks", getTasksHandler)

	// Endpoint untuk menambahkan tugas baru
	e.POST("/tasks", addTaskHandler)

	// Endpoint untuk menghapus tugas berdasarkan ID
	e.DELETE("/tasks/:id", deleteTaskHandler)

	// Endpoint untuk mendapatkan dokumentasi Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Endpoint untuk mengedit tugas berdasarkan ID
	e.PUT("/tasks/:id", updateTaskHandler)

	port := 8080
	fmt.Printf("Server berjalan di http://localhost:%d\n", port)
	log.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
