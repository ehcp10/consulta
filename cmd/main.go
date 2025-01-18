package main

import (
	"github.com/ehcp10/consulta/internal/domain"
	"github.com/ehcp10/consulta/internal/repository"
	"log"
	"os"
)

func main() {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "secret")
	os.Setenv("DB_NAME", "psych_db")
	os.Setenv("DB_PORT", "5432")

	// 1. Conectar ao BD
	db, err := repository.NewDB()
	if err != nil {
		log.Fatal("Erro ao conectar ao DB:", err)
	}

	// 2. AutoMigrate (apenas para desenvolvimento / protótipo)
	if err := db.AutoMigrate(
		&domain.Client{},
		&domain.ContactClient{},
		&domain.EmergencyContact{},
		&domain.Appointment{},
	); err != nil {
		log.Fatal("Erro ao migrar o BD:", err)
	}
}
