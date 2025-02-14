package http

import (
	"github.com/ehcp10/consulta/internal/domain"
	"github.com/ehcp10/consulta/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ClientHandler struct {
	clientUC usecase.ClientUseCase
}

func NewClientHandler(r *gin.Engine, uc usecase.ClientUseCase) {
	handler := &ClientHandler{
		clientUC: uc,
	}

	r.POST("/clients", handler.CreateClient)
	r.GET("/client/:id", handler.GetClientByID)
	r.PATCH("/client/:id", handler.UpdateClientByID)
	r.DELETE("/client/:id", handler.DeleteClientByID)
}

func (h *ClientHandler) CreateClient(c *gin.Context) {
	var client domain.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.clientUC.CreateNewClient(c.Request.Context(), &client); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, client)
}

func (h *ClientHandler) GetClientByID(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client, err := h.clientUC.GetClientByID(c.Request.Context(), intID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, client)
}

func (h *ClientHandler) UpdateClientByID(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var client domain.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.clientUC.EditClientByID(c.Request.Context(), intID, &client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, client)
}

func (h *ClientHandler) DeleteClientByID(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err = h.clientUC.DeleteClientByID(c.Request.Context(), intID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusNoContent, gin.H{})
	return
}
