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

type updateListRequest struct {
	Name     string `json:"name" binding:"required,min=1,max=100"`
	Position int    `json:"position"`
}

func (h *Handler) CreateList(c *gin.Context) {

	projectIdParam := c.Param("id")
	projectId, err := strconv.Atoi(projectIdParam)
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

	project, err := h.db.GetProject(uint(projectId))
	if err != nil {
		c.JSON(404, gin.H{"error": "Not found"})
		return
	}

	if project.UserID != user.ID {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	listPos := len(project.Lists) + 1

	list := &model.List{
		Name:      req.Name,
		ProjectID: uint(projectId),
		Position:  listPos,
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
	var req updateListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	list, err := h.db.GetList(uint(listId))

	if err != nil {
		c.JSON(404, gin.H{"error": "Not found"})
		return
	}

	project, err := h.db.GetProject(list.ProjectID)
	if err != nil {
		c.JSON(404, gin.H{"error": "Not found"})
		return
	}

	if project.UserID != user.ID {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	list.Name = req.Name

	// update positions
	if req.Position != 0 {
		if req.Position > list.Position {
			for i := 0; i < len(project.Lists); i++ {
				if project.Lists[i].Position > list.Position && project.Lists[i].Position <= req.Position {
					project.Lists[i].Position--
					if err := h.db.UpdateList(&project.Lists[i]); err != nil {
						c.JSON(500, gin.H{"error": "Internal error"})
						return
					}
				}
			}
		} else if req.Position < list.Position {
			for i := 0; i < len(project.Lists); i++ {
				if project.Lists[i].Position < list.Position && project.Lists[i].Position >= req.Position {
					project.Lists[i].Position++
					if err := h.db.UpdateList(&project.Lists[i]); err != nil {
						c.JSON(500, gin.H{"error": "Internal error"})
						return
					}
				}
			}
		}
		list.Position = req.Position
	}

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

	// update positions
	project, err := h.db.GetProject(list.ProjectID)

	if err != nil {
		c.JSON(404, gin.H{"error": "Not found"})
		return
	}

	for i := 0; i < len(project.Lists); i++ {
		if project.Lists[i].Position > list.Position {
			project.Lists[i].Position--
			if err := h.db.UpdateList(&project.Lists[i]); err != nil {
				c.JSON(500, gin.H{"error": "Internal error"})
				return
			}
		}
	}

	c.JSON(200, list)
}
