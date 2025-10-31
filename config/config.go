package config

import (
	"log"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Postgres PostgresConfig
	DB       DBPoolConfig
}

type PostgresConfig struct {
	Host     string `env:"POSTGRES_HOST,required"`
	Port     string `env:"POSTGRES_PORT" envDefault:"5432"`
	User     string `env:"POSTGRES_USER" envDefault:"postgres"`
	Password string `env:"POSTGRES_PASSWORD,required"`
	DBName   string `env:"POSTGRES_DB" envDefault:"appdb"`
}

type DBPoolConfig struct {
	MaxOpenConns    int           `env:"DB_MAX_OPEN_CONNS" envDefault:"60"`
	MaxIdleConns    int           `env:"DB_MAX_IDLE_CONNS" envDefault:"30"`
	ConnMaxLifetime time.Duration `env:"DB_CONN_MAX_LIFETIME" envDefault:"30m"`
	ConnMaxIdleTime time.Duration `env:"DB_CONN_MAX_IDLE_TIME" envDefault:"20m"`
}

func Load() (*Config, error) {
	// Загружаем .env (игнорируем, если нет)
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file: %v", err)
	}

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	log.Printf("Config loaded: DB=%s, Host=%s, Pool=%d/%d",
		cfg.Postgres.DBName, cfg.Postgres.Host,
		cfg.DB.MaxOpenConns, cfg.DB.MaxIdleConns)

	return cfg, nil
}
