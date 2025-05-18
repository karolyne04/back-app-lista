package persistence

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq" // Driver PostgreSQL
)

type PostgresDB struct {
	DB *sql.DB
}

func NewPostgresDB(host, port, user, password, dbname string) (*PostgresDB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao PostgreSQL: %w", err)
	}

	// Configurar pool de conexões
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Verificar conexão
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("erro ao verificar conexão com PostgreSQL: %w", err)
	}

	log.Println("Conexão com PostgreSQL estabelecida com sucesso")
	return &PostgresDB{DB: db}, nil
}

// Close fecha a conexão com o banco de dados
func (p *PostgresDB) Close() error {
	return p.DB.Close()
}