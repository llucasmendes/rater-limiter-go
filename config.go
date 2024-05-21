package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	RedisAddr      string
	RedisPassword  string
	RateLimitIP    int
	RateLimitToken int
	BlockDuration  int
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	rateLimitIP, err := strconv.Atoi(os.Getenv("RATE_LIMIT_IP"))
	if err != nil {
		log.Fatal("Invalid RATE_LIMIT_IP value")
	}

	rateLimitToken, err := strconv.Atoi(os.Getenv("RATE_LIMIT_TOKEN"))
	if err != nil {
		log.Fatal("Invalid RATE_LIMIT_TOKEN value")
	}

	blockDuration, err := strconv.Atoi(os.Getenv("BLOCK_DURATION"))
	if err != nil {
		log.Fatal("Invalid BLOCK_DURATION value")
	}

	return Config{
		RedisAddr:      os.Getenv("REDIS_ADDR"),
		RedisPassword:  os.Getenv("REDIS_PASSWORD"),
		RateLimitIP:    rateLimitIP,
		RateLimitToken: rateLimitToken,
		BlockDuration:  blockDuration,
	}
}
