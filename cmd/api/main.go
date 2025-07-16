package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"karol/internal/application/usecase"
	"karol/internal/domain/service"
	"karol/internal/infrastructure/persistence"
	"karol/internal/interfaces/api"
	"karol/internal/interfaces/api/handler"
	"karol/internal/interfaces/api/middleware"
)

func main() {
	// Configurações do banco de dados
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "postgres")
	dbName := getEnv("DB_NAME", "app_db")

	// Configurações da aplicação
	jwtSecret := getEnv("JWT_SECRET", "seu_segredo_jwt_super_secreto")
	jwtIssuer := getEnv("JWT_ISSUER", "api-go-ddd")
	serverPort := getEnv("SERVER_PORT", "8080")

	// Inicializar conexão com o banco de dados
	db, err := persistence.NewPostgresDB(dbHost, dbPort, dbUser, dbPassword, dbName)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	// Inicializar repositórios
	userRepository := persistence.NewPostgresUserRepository(db)
	shoppingListRepository := persistence.NewPostgresShoppingListRepository(db)
	listItemRepository := persistence.NewPostgresListItemRepository(db)
	categoryRepository := persistence.NewPostgresCategoryRepository(db)
	tipRepository := persistence.NewPostgresTipRepository(db)

	// Inicializar serviços
	jwtService := service.NewJWTService(jwtSecret, jwtIssuer)

	// Inicializar casos de uso
	authUseCase := usecase.NewAuthUseCase(userRepository, jwtService)
	userUseCase := usecase.NewUserUseCase(userRepository)
	shoppingListUseCase := usecase.NewShoppingListUseCase(shoppingListRepository)
	listItemUseCase := usecase.NewListItemUseCase(listItemRepository)
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepository)
	tipUseCase := usecase.NewTipUseCase(tipRepository)

	// Inicializar handlers
	authHandler := handler.NewAuthHandler(authUseCase)
	userHandler := handler.NewUserHandler(userUseCase)
	shoppingListHandler := handler.NewShoppingListHandler(shoppingListUseCase)
	listItemHandler := handler.NewListItemHandler(listItemUseCase)
	categoryHandler := handler.NewCategoryHandler(categoryUseCase)
	tipHandler := handler.NewTipHandler(tipUseCase)

	// Inicializar middlewares
	authMiddleware := middleware.NewAuthMiddleware(jwtService)

	// Configurar router
	router := api.NewRouter(
		authHandler,
		userHandler,
		authMiddleware,
		shoppingListHandler,
		listItemHandler,
		categoryHandler,
		tipHandler,
	)
	handler := router.Setup()

	// Configurar servidor HTTP
	server := &http.Server{
		Addr:         ":" + serverPort,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Iniciar servidor
	log.Printf("Servidor iniciado na porta %s", serverPort)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}

// getEnv retorna o valor da variável de ambiente ou um valor padrão
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
