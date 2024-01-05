package controllers

import (
	"net/http"
	"time"
	"waizly-tech-test/models"
	"waizly-tech-test/utils/token"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TaskInput struct {
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	IsDone      bool      `json:"is_done type:boolean"`
}

// Create a Task
// @Summary Create Task
// @Description create new Task Task
// @Tags Task
// @Param Body body TaskInput true "the body to create new Task"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Task
// @Router /task [post]
func CreateTask(c *gin.Context) {
	var input TaskInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer_id, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := models.Task{CustomerID: customer_id, Description: input.Description, DueDate: input.DueDate, IsDone: input.IsDone}

	db := c.MustGet("db").(*gorm.DB)

	db.Create(&task)
	c.JSON(http.StatusOK, gin.H{"data": task})

}

// GetAllTask
// @Summary Get all Task.
// @Description Get a list of Task.
// @Tags Task
// @Produce json
// @Success 200 {object} []models.Task
// @Router /task [get]
func GetAllTask(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var task []models.Task
	db.Find(&task)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

// GetTaskByCustomerId Waizly
// @Summary Get Task.
// @Description Get an Task by id.
// @Tags Task
// @Produce json
// @Param id path string true "Task id"
// @Success 200 {object} models.Task
// @Router /task/{customer_id} [get]
func GetTaskByCustomerId(c *gin.Context) { // Get model if exist
	var task models.Task

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("customer_id = ?", c.Param("customer_id")).Find(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

// UpdateTask
// @Summary Update Task.
// @Description Update Task by id.
// @Tags Task
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Task id"
// @Param Body body ProductInput true "the body to update Task"
// @Success 200 {object} models.Task
// @Router /task/{id} [patch]
func UpdateTask(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var task models.Task
	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input TaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Task
	updatedInput.Description = input.Description
	updatedInput.DueDate = input.DueDate
	updatedInput.IsDone = input.IsDone
	updatedInput.UpdatedAt = time.Now()

	db.Model(&task).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

// DeleteTask
// @Summary Delete one Task.
// @Description Delete a Task by id.
// @Tags Product
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Task id"
// @Success 200 {object} map[string]boolean
// @Router /task/{id} [delete]
func DeleteTask(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var task models.Task
	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&task)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
