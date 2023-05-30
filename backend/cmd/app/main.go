package main

import (
	"fmt"

	"github.com/expose443/forum/backend/pkg/configs"
	"github.com/expose443/forum/backend/pkg/logger"
)

func main() {
	logger := logger.New()
	cfg := configs.NewConfig(logger)
	fmt.Println(cfg.GetString("DB_PORT"))

}
