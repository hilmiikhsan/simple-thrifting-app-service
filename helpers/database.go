package helpers

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	DB          *sqlx.DB
	RedisClient *redis.Client
	RedisCtx    = context.Background()
)

func SetupPostgres() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		GetEnv("DB_HOST", "127.0.0.1"),
		GetEnv("DB_USER", ""),
		GetEnv("DB_PASSWORD", ""),
		GetEnv("DB_NAME", ""),
		GetEnv("DB_PORT", "5432"),
	)

	DB, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		Logger.Fatal("failed to connect to database: ", err)
	}

	err = SeedUserAccount(DB)
	if err != nil {
		Logger.Fatalf("Error seeding admin account: %v\n", err)
	}

	Logger.Info("Successfully connected to PostgreSQL database...")
}

func SetupRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", GetEnv("REDIS_HOST", "127.0.0.1"), GetEnv("REDIS_PORT", "6379")),
		Password: GetEnv("REDIS_PASSWORD", ""),
		DB:       GetEnvInt("REDIS_DB", 0),
	})

	_, err := RedisClient.Ping(RedisCtx).Result()
	if err != nil {
		Logger.Fatal("failed to connect to Redis: ", err)
	}

	Logger.Info("Successfully connected to Redis...")
}
