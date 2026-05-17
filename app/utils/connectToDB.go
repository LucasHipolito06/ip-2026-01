package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectToDB() *sql.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Erro ao testar conexão com banco:", err)
	}

	fmt.Println("Conectado ao banco de dados com sucesso!")
	return db
}

func CreatePacientesTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS pacientes (
		id SERIAL PRIMARY KEY,
		nome VARCHAR(100) NOT NULL,
		cpf VARCHAR(14) UNIQUE NOT NULL,
		data_nascimento DATE NOT NULL,
		telefone VARCHAR(20),
		endereco VARCHAR(150),
		convenio VARCHAR(100),
		diagnostico VARCHAR(150),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Erro ao criar tabela pacientes:", err)
	}

	fmt.Println("Tabela pacientes verificada/criada com sucesso!")
}