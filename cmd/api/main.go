package api

import (
	"log"
	"subscription-api/internal/handlers"
	"subscription-api/internal/repositories"
	"subscription-api/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=root dbname=loopay port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Falha ao conectar ao banco de dados: %v", err)
	}

	userRepo := repositories.NewUserRepository(db)
	planRepo := repositories.NewPlanRepository(db)
	paymentService := services.NewPaymentService()

	router := gin.Default()

	handlers.RegisterRoutes(router, userRepo, planRepo, paymentService)

	log.Println("Server rodando na porta 8080")
	router.Run(":8080")
}
