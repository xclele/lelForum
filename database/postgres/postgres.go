package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"

	"lelForum/settings"
)

var db *sqlx.DB

// Init initializes the PostgreSQL database connection
func Init(cfg *settings.PostgreSQLConfig) (err error) {
	if cfg == nil {
		return fmt.Errorf("postgres config is nil")
	}

	// Build DSN, only include password if it's not empty
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.DB, cfg.SSLMode,
	)
	if cfg.Password != "" {
		dsn = fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DB, cfg.SSLMode,
		)
	}

	// Use sqlx.Open instead of sqlx.Connect to avoid premature connection
	db, err = sqlx.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("open postgres failed: %w", err)
	}

	// Set connection pool parameters before pinging
	if cfg.MaxOpenConns > 0 {
		db.SetMaxOpenConns(cfg.MaxOpenConns)
	}
	if cfg.MaxIdleConns > 0 {
		db.SetMaxIdleConns(cfg.MaxIdleConns)
	}

	// Verify the connection
	if err = db.Ping(); err != nil {
		return fmt.Errorf("ping postgres failed: %w", err)
	}

	zap.L().Info("postgres connected successfully")
	return nil
}

func Close() {
	if db == nil {
		return
	}
	if err := db.Close(); err != nil {
		zap.L().Error("close postgres failed", zap.Error(err))
		return
	}
	zap.L().Info("postgres connection closed")
}

func DB() *sqlx.DB {
	return db
}
