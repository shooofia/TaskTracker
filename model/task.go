// model/task.go

package model

import "gorm.io/gorm"

type Task struct {
	ID             int    `json:"id" gorm:"primaryKey"`
	Judul          string `json:"judul"`
	Deskripsi      string `json:"deskripsi"`
	Prioritas      string `json:"prioritas"`
	TanggalTenggat string `json:"tanggal_tenggat"`
	Status         string `json:"status"`
}

type TaskManager interface {
	AddTask(task *Task) error
	UpdateTaskStatus(id int) error
	GetTasks() ([]Task, error)
	GetTaskByID(id int) (Task, error)
}

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskManager(db *gorm.DB) TaskManager {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) AddTask(task *Task) error {
	return r.db.Create(task).Error
}

func (r *TaskRepository) UpdateTaskStatus(id int) error {
	task := Task{}
	err := r.db.First(&task, id).Error
	if err != nil {
		return err
	}

	task.Status = "Done"
	return r.db.Save(&task).Error
}

func (r *TaskRepository) GetTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepository) GetTaskByID(id int) (Task, error) {
	task := Task{}
	err := r.db.First(&task, id).Error
	return task, err
}
