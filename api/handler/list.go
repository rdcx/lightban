package handler

import (
	"lightban/api/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type createListRequest struct {
	Name      string `json:"name" binding:"required,min=1,max=100"`
	ProjectID uint   `json:"project_id"`
}

func (h *Handler) CreateList(c *gin.Context) {
	u, ok := c.Get("user")
	if !ok {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	user := u.(*model.User)
	var req createListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	project, err := h.db.GetProject(req.ProjectID)
	if err != nil {
		c.JSON(404, gin.H{"error": "Not found"})
		return
	}

	if project.UserID != user.ID {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	list := &model.List{
		Name:      req.Name,
		ProjectID: req.ProjectID,
	}

	if err := h.db.CreateList(list); err != nil {
		c.JSON(500, gin.H{"error": "Internal error"})
		return
	}

	c.JSON(200, list)
}

func (h *Handler) UpdateList(c *gin.Context) {

	listIdParam := c.Param("id")
	listId, err := strconv.Atoi(listIdParam)

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	u, ok := c.Get("user")
	if !ok {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	user := u.(*model.User)
	var req createListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	project, err := h.db.GetProject(req.ProjectID)
	if err != nil {
		c.JSON(404, gin.H{"error": "Not found"})
		return
	}

	if project.UserID != user.ID {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	list, err := h.db.GetList(uint(listId))

	if err != nil {
		c.JSON(404, gin.H{"error": "Not found"})
		return
	}

	list.Name = req.Name

	if err := h.db.UpdateList(list); err != nil {
		c.JSON(500, gin.H{"error": "Internal error"})
		return
	}

	c.JSON(200, list)
}

func (h *Handler) DeleteList(c *gin.Context) {

	listIdParam := c.Param("id")
	listId, err := strconv.Atoi(listIdParam)

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	list, err := h.db.GetList(uint(listId))

	if err != nil {
		c.JSON(404, gin.H{"error": "Not found"})
		return
	}

	if err := h.db.DeleteListByID(list.ID); err != nil {
		c.JSON(500, gin.H{"error": "Internal error"})
		return
	}

	c.JSON(200, list)
}
