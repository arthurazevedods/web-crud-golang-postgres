package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Definições de constantes para modos de SSL
const (
	DB_SSLMODE_DISABLE     = "disable"
	DB_SSLMODE_ALLOW       = "allow"
	DB_SSLMODE_PREFER      = "prefer"
	DB_SSLMODE_REQUIRE     = "require"
	DB_SSLMODE_VERIFY_CA   = "verify-ca"
	DB_SSLMODE_VERIFY_FULL = "verify-full"
)

// Estrutura para armazenar as variáveis de ambiente
type DbSSLMode string

type Config struct {
	ApiEnv           string    `env:"API_ENV,required"`
	ApiHost          string    `env:"API_HOST,required"`
	ApiPort          string    `env:"API_PORT,required"`
	DatabaseHost     string    `env:"DATABASE_HOST,required"`
	DatabasePort     string    `env:"DATABASE_PORT,required"`
	DatabaseUsername string    `env:"DATABASE_USERNAME,required"`
	DatabasePassword string    `env:"DATABASE_PASSWORD,required"`
	DatabaseDBName   string    `env:"DATABASE_DBNAME,required"`
	DatabaseSSLMode  DbSSLMode `env:"DATABASE_SSL_MODE,required"`
}

// Função para carregar o arquivo .env, se ele existir
func loadEnvFile() error {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		return nil
	}
	return godotenv.Load(".env")
}

// Função para obter as configurações do ambiente
func GetConfig() (Config, error) {
	if err := loadEnvFile(); err != nil {
		return Config{}, err
	}

	var config Config
	if err := env.Parse(&config); err != nil {
		return Config{}, err
	}

	// Validação de SSL Mode
	validSSLModes := []string{
		DB_SSLMODE_DISABLE, DB_SSLMODE_ALLOW, DB_SSLMODE_PREFER,
		DB_SSLMODE_REQUIRE, DB_SSLMODE_VERIFY_CA, DB_SSLMODE_VERIFY_FULL,
	}
	valid := false
	for _, sslmode := range validSSLModes {
		if string(config.DatabaseSSLMode) == sslmode {
			valid = true
			break
		}
	}
	if !valid {
		return Config{}, fmt.Errorf("invalid DATABASE_SSL_MODE: %s", config.DatabaseSSLMode)
	}

	return config, nil
}

// Função para conectar ao banco de dados usando a configuração
func ConectaComBancoDeDados() (*sql.DB, error) {
	config, err := GetConfig()
	if err != nil {
		return nil, err
	}

	// Construir string de conexão
	conexao := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		config.DatabaseUsername, config.DatabasePassword, config.DatabaseDBName,
		config.DatabaseHost, config.DatabasePort, config.DatabaseSSLMode,
	)

	// Conectar ao banco
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar com o banco: %v", err)
	}

	return db, nil
}
