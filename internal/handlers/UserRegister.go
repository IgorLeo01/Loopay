package handlers

import (
	"Loopay/internal/models"
	"Loopay/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, userRepo *repositories.UserRepository, planRepo *repositories.PlanRepository, paymentService *services.PaymentService) {
	r.POST("/users", func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := userRepo.CreateUser(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao criar usuário" err})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"mensagem": "Usuário criado com sucesso"})
	})
}
