package handler

import (
	"lightban/api/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type createTaskRequest struct {
	Name        string `json:"name" binding:"required,min=1,max=100"`
	Description string `json:"description" binding:"min=0,max=1000"`
}

type updateTaskRequest struct {
	Name        string `json:"name" binding:"required,min=1,max=100"`
	Description string `json:"description" binding:"min=0,max=1000"`
	ListID      uint   `json:"list_id" binding:"required"`
}

func (h *Handler) CreateTask(c *gin.Context) {

	id := c.Param("id")

	listId, err := strconv.Atoi(id)

	u, ok := c.Get("user")
	if !ok {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	user := u.(*model.User)
	var req createTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	list, err := h.db.GetList(uint(listId))
	if err != nil {
		c.JSON(404, gin.H{"error": "Not found"})
		return
	}

	if list.Project.UserID != user.ID {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	task := &model.Task{
		Name:        req.Name,
		Description: req.Description,
		ListID:      list.ID,
	}

	if err := h.db.CreateTask(task); err != nil {
		c.JSON(500, gin.H{"error": "Internal error"})
		return
	}

	c.JSON(200, task)
}

func (h *Handler) UpdateTask(c *gin.Context) {

	taskIdParam := c.Param("id")
	taskId, err := strconv.Atoi(taskIdParam)

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	var req updateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	task, err := h.db.GetTask(uint(taskId))
	if err != nil {
		c.JSON(404, gin.H{"error": "Not found"})
		return
	}

	list, err := h.db.GetList(req.ListID)

	if err != nil {
		c.JSON(404, gin.H{"error": "Not found"})
		return
	}

	if list.Project.UserID != task.List.Project.UserID {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	task.Name = req.Name
	task.Description = req.Description
	task.ListID = list.ID
	task.List = list

	if err := h.db.UpdateTask(task); err != nil {
		c.JSON(500, gin.H{"error": "Internal error"})
		return
	}

	c.JSON(200, task)
}

func (h *Handler) DeleteTask(c *gin.Context) {

	taskIdParam := c.Param("id")
	taskId, err := strconv.Atoi(taskIdParam)

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	task, err := h.db.GetTask(uint(taskId))

	if err != nil {
		c.JSON(404, gin.H{"error": "Not found"})
		return
	}

	if err := h.db.DeleteTaskByID(task.ID); err != nil {
		c.JSON(500, gin.H{"error": "Internal error"})
		return
	}

	c.JSON(200, task)
}
