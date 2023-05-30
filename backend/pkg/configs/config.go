package configs

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/expose443/forum/backend/pkg/logger"
)

func NewConfig(logger *logger.LogLevel) *Config {
	logger.Warning("setting defaults variables for .env")
	setDefaults()
	logger.Warning("setting variables from .env")
	err := setEnv()
	if err != nil {
		logger.Error(fmt.Sprintf("error when setting values from .env %v", err))
	}
	return &Config{
		logger: logger,
	}
}

type Config struct {
	logger *logger.LogLevel
}

func (cfg *Config) GetString(key string) string {
	return os.Getenv(key)
}
func (cfg *Config) GetInt(key string) int {
	value := os.Getenv(key)
	num, err := strconv.Atoi(value)
	if err != nil {
		cfg.logger.Error(err.Error())
		return -1
	}
	return num
}

func setDefaults() {
	os.Setenv("SERVER_PORT", "8080")
	os.Setenv("SERVER_ADDRESS", "http://localhost:8080")
	os.Setenv("DB_NAME", "database.db")
	os.Setenv("DB_PASSWORD", "secret")
	os.Setenv("DB_PORT", "8191")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("SERVER_READ_TIMEOUT", "10")
	os.Setenv("SERVER_WRITE_TIMEOUT", "10")
	os.Setenv("SERVER_IDLE_TIMEOUT", "120")
}

func setEnv() error {
	path := filepath.Join(".env")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return err
	}
	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	envs := make(map[string]string)

	vars := strings.Split(string(file), "\n")
	for i := range vars {
		e := strings.Split(vars[i], "=")
		if len(e) != 2 {
			return fmt.Errorf(fmt.Sprintf("invalid format %s", path))
		}
		envs[strings.ReplaceAll(e[0], " ", "")] = strings.ReplaceAll(e[1], " ", "")
	}
	for key, value := range envs {
		err := os.Setenv(key, value)
		if err != nil {
			return err
		}
	}
	return nil
}
