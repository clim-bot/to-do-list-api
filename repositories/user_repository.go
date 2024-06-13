package repositories

import (
    "gorm.io/gorm"
    "github.com/clim-bot/todo-list-api/models"
)

func CreateUser(db *gorm.DB, user *models.User) error {
    return db.Create(user).Error
}

func GetUserByEmail(db *gorm.DB, email string) (*models.User, error) {
    var user models.User
    if err := db.Where("email = ?", email).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}
