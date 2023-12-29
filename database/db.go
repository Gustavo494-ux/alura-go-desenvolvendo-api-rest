package database

import (
	"log"

	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConectaComBancoDeDados() (*gorm.DB, error) {
	// Configurar informações de conexão
	stringConexao := "host=localhost user=root password=root dbname=root port=5433 sslmode=disable"

	// Conectar ao banco de dados PostgreSQL
	db, err := gorm.Open(postgres.Open(stringConexao), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados", err)
	}

	return db, nil
}
