package main

import (
	"fmt"
	"github.com/ehcp10/consulta/internal/domain"
	"github.com/ehcp10/consulta/internal/http"
	"github.com/ehcp10/consulta/internal/repository"
	"github.com/ehcp10/consulta/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	//os.Setenv("DB_HOST", "localhost")
	//os.Setenv("DB_USER", "postgres")
	//os.Setenv("DB_PASSWORD", "secret")
	//os.Setenv("DB_NAME", "psych_db")
	//os.Setenv("DB_PORT", "5432")

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	apiPort := os.Getenv("API_PORT")

	// Exemplo: Conectando ao banco de dados
	fmt.Printf("Conectando ao banco de dados:\n")
	fmt.Printf("Host: %s, User: %s, Password: %s, Database: %s\n", dbHost, dbUser, dbPassword, dbName)

	db, err := repository.NewDB()
	if err != nil {
		log.Fatal("Erro ao conectar ao DB:", err)
	}

	if err := db.AutoMigrate(
		&domain.Client{},
		&domain.ContactClient{},
		&domain.EmergencyContact{},
		&domain.Appointment{},
	); err != nil {
		log.Fatal("Erro ao migrar o BD:", err)
	}

	r := gin.Default()

	//repositories
	clientRepo := repository.NewClientRepository(db)

	//usecases
	clientUC := usecase.NewClientUseCase(clientRepo)

	//handlers
	http.NewClientHandler(r, clientUC)

	log.Printf("Serving on port %s", apiPort)
	if err := r.Run(":" + apiPort); err != nil {
		log.Fatal("Erro ao rodar servidor:", err)
	}

}
