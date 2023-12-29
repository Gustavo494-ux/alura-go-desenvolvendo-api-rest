package database

import (
	"log"

	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	Err error
)

func ConectaComBancoDeDados() {
	// Configurar informações de conexão
	stringConexao := "host=localhost user=root password=root dbname=root port=5433 sslmode=disable"

	// Conectar ao banco de dados PostgreSQL
	DB, Err = gorm.Open(postgres.Open(stringConexao), &gorm.Config{})
	if Err != nil {
		log.Panic("Erro ao conectar com o banco de dados", Err)
	}
}
