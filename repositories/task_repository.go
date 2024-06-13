package repositories

import (
    "gorm.io/gorm"
    "github.com/clim-bot/todo-list-api/models"
)

func CreateTask(db *gorm.DB, task *models.Task) error {
    return db.Create(task).Error
}

func GetTasksByUserID(db *gorm.DB, userID uint) ([]models.Task, error) {
    var tasks []models.Task
    if err := db.Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
        return nil, err
    }
    return tasks, nil
}

func UpdateTask(db *gorm.DB, task *models.Task) error {
    return db.Save(task).Error
}

func DeleteTask(db *gorm.DB, task *models.Task) error {
    return db.Delete(task).Error
}
