package main

import (
	"database/sql"
	"log"

	"go-user-api/config"
	db "go-user-api/db/sqlc"
	"go-user-api/internal/handler"
	"go-user-api/internal/logger"
	"go-user-api/internal/routes"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// 1. Initialize Logger
	logger.InitLogger()
	defer logger.Log.Sync()

	// 2. Load config from .env
	cfg := config.LoadConfig()

	// 3. Connect to Database using config
	dbConn, err := sql.Open(cfg.DBDriver, cfg.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	// 3. Setup Fiber and Validator
	app := fiber.New()
	validate := validator.New()

	// 4. Initialize Handler
	userHandler := &handler.UserHandler{
		DBQueries: queries,
		Validator: validate,
	}

	// 6. Setup Routes
	routes.SetupRoutes(app, userHandler)

	// 7. Start Server
	logger.Log.Info("Server is starting on port " + cfg.Port + "...")
	log.Fatal(app.Listen(cfg.Port))
}
