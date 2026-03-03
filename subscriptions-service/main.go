package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ZakSlinin/Technostrelka-pobeda-backend/subscriptions-service/handler"
	"github.com/ZakSlinin/Technostrelka-pobeda-backend/subscriptions-service/repository"
	"github.com/ZakSlinin/Technostrelka-pobeda-backend/subscriptions-service/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	pgUser := os.Getenv("DB_USER")
	pgPass := os.Getenv("DB_PASSWORD")
	pgHost := os.Getenv("DB_HOST")
	pgPort := os.Getenv("DB_PORT")
	pgDB := os.Getenv("DB_NAME")

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		pgUser, pgPass, pgHost, pgPort, pgDB,
	)

	m, err := migrate.New(
		"file://migrations",
		dbURL,
	)
	if err != nil {
		log.Fatalf("failed to init migrations: %v", err)
	}
	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Fatalf("failed to apply migrations: %v", err)
	}
	log.Println("Migrations applied successfully")

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatalf("failed to connect to database via GORM: %v", err)
	}

	repo := repository.NewPostgresSubscriptionsRepository(db)
	service := service.NewSubscriptionsService(repo)
	handler := handler.NewSubscriptionsHandler(service)

	r := gin.Default()

	r.Static("/uploads", "/app/uploads")

	api := r.Group("/api/subscriptions")
	{
		api.POST("/create", handler.Create)
		api.GET("/all", handler.GetAllUserByID)
		api.PATCH("/update/:subscription_id", handler.UpdateSubscriptionByID)
		api.DELETE("/delete/:subscription_id", handler.DeleteSubscriptionByID)
	}

	r.Run(":8080")
}
