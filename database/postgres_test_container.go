package database

import (
	"context"
	"time"

	tcpostgres "github.com/testcontainers/testcontainers-go/modules/postgres"
	gormpostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TestDBContainer struct {
	Ctx       context.Context
	Container *tcpostgres.PostgresContainer
	DB        *gorm.DB
}

func SetupPostgresContainer() (*TestDBContainer, error) {
	ctx := context.Background()

	container, err := tcpostgres.Run(
		ctx,
		"postgres:16-alpine",
		tcpostgres.WithDatabase("testdb"),
		tcpostgres.WithUsername("test"),
		tcpostgres.WithPassword("12345"),
		tcpostgres.BasicWaitStrategies(),
	)
	if err != nil {
		return nil, err
	}

	dsn, err := container.ConnectionString(
		ctx,
		"sslmode=disable",
	)
	if err != nil {
		_ = container.Terminate(ctx)
		return nil, err
	}

	db, err := gorm.Open(
		gormpostgres.Open(dsn),
		&gorm.Config{},
	)
	if err != nil {
		_ = container.Terminate(ctx)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err == nil {
		sqlDB.SetMaxOpenConns(5)
		sqlDB.SetConnMaxLifetime(1 * time.Minute)
	}

	return &TestDBContainer{
		Ctx:       ctx,
		Container: container,
		DB:        db,
	}, nil
}
