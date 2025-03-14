package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	retry2 "github.com/SlimeMutation/btc-wallet-go/common/retry"
	"github.com/SlimeMutation/btc-wallet-go/config"
	_ "github.com/SlimeMutation/btc-wallet-go/database/utils/serializers"
)

type DB struct {
	gorm *gorm.DB

	CreateTable CreateTableDB
	Business    BusinessDB
}

func NewDB(ctx context.Context, dbConfig config.DBConfig) (*DB, error) {
	dsn := fmt.Sprintf("host=%s dbname=%s sslmode=disable", dbConfig.Host, dbConfig.Name)
	if dbConfig.Port != 0 {
		dsn += fmt.Sprintf(" port=%d", dbConfig.Port)
	}
	if dbConfig.User != "" {
		dsn += fmt.Sprintf(" user=%s", dbConfig.User)
	}
	if dbConfig.Password != "" {
		dsn += fmt.Sprintf(" password=%s", dbConfig.Password)
	}

	newLogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)

	gormConfig := gorm.Config{
		SkipDefaultTransaction: true,
		CreateBatchSize:        3_000,
		Logger:                 newLogger, // Set the logger
	}

	retryStrategy := &retry2.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
	gormDbBox, err := retry2.Do[*gorm.DB](context.Background(), 10, retryStrategy, func() (*gorm.DB, error) {
		gormDb, err := gorm.Open(postgres.Open(dsn), &gormConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to database: %w", err)
		}
		return gormDb, nil
	})

	if err != nil {
		return nil, err
	}

	db := &DB{
		gorm:        gormDbBox,
		CreateTable: NewCreateTableDB(gormDbBox),
		Business:    NewBusinessDB(gormDbBox),
	}
	return db, nil
}

func (db *DB) Transaction(fn func(db *DB) error) error {
	return db.gorm.Transaction(func(tx *gorm.DB) error {
		txDB := &DB{
			gorm:     tx,
			Business: NewBusinessDB(tx),
		}
		return fn(txDB)
	})
}

func (db *DB) Close() error {
	sql, err := db.gorm.DB()
	if err != nil {
		return err
	}
	return sql.Close()
}

func (db *DB) ExecuteSQLMigration(migrationsFolder string) error {
	err := filepath.Walk(migrationsFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("Failed to process migration file: %s", path))
		}
		if info.IsDir() {
			return nil
		}
		fileContent, readErr := os.ReadFile(path)
		if readErr != nil {
			return errors.Wrap(readErr, fmt.Sprintf("Error reading SQL file: %s", path))
		}

		execErr := db.gorm.Exec(string(fileContent)).Error
		if execErr != nil {
			return errors.Wrap(execErr, fmt.Sprintf("Error executing SQL script: %s", path))
		}
		return nil
	})
	return err
}
