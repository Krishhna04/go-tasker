package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krishna/go-tasker/internal/models"
)

func GetTasks(c *gin.Context) {
	var tasks []models.Task
	userId := c.MustGet("user_id").(uint)
	models.DB.Where("user_id = ?", userId).Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

func CreateTask(c *gin.Context) {
	var input models.Task
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := c.MustGet("user_id").(uint)
	task := models.Task{Title: input.Title, Description: input.Description, UserID: userId}
	models.DB.Create(&task)
	c.JSON(http.StatusOK, task)
}

func GetTask(c *gin.Context) {
	var task models.Task
	if err := models.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
	var task models.Task
	if err := models.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	var input models.Task
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.Title = input.Title
	task.Description = input.Description
	models.DB.Save(&task)

	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	var task models.Task
	if err := models.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	models.DB.Delete(&task)
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
