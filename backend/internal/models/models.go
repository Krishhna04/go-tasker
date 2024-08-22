package models

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Username string `json:"username"`
    Password string `json:"-"`
    Tasks    []Task
}

type Task struct {
    gorm.Model
    Title       string `json:"title"`
    Description string `json:"description"`
    UserID      uint   `json:"user_id"`
}
