package dotenv

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct{}

func New(env string) (*Config, error) {
	godotenv.Load(fmt.Sprintf(".env.%s", env)) //nolint:errcheck //env vars are stored remotely
	return &Config{}, nil
}

func (c *Config) Get(key string) any {
	return os.Getenv(key)
}

func (c *Config) GetBool(key string) bool {
	v := os.Getenv(key)

	b, err := strconv.ParseBool(v)
	if err != nil {
		panic("failed to parse value to bool")
	}

	return b
}

func (c *Config) GetFloat64(key string) float64 {
	v := os.Getenv(key)

	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		panic("failed to parse value to float64")
	}

	return f
}

func (c *Config) GetInt(key string) int {
	v := os.Getenv(key)

	i, err := strconv.Atoi(v)
	if err != nil {
		panic("failed to parse value to int")
	}

	return i
}

func (c *Config) GetString(key string) string {
	return os.Getenv(key)
}
