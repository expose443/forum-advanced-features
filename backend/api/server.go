package api

import (
	"crypto/tls"
	"net/http"
	"path/filepath"
	"time"

	"github.com/expose443/forum/backend/pkg/configs"
	"github.com/expose443/forum/backend/pkg/logger"
)

func NewServer(cfg configs.Config, logger logger.LogLevel) *http.Server {
	pathCrt := filepath.Join("./certificats/", "certificate.crt")
	pathKey := filepath.Join("./certificats/", "private.key")
	cert, err := tls.LoadX509KeyPair(pathCrt, pathKey)
	if err != nil {
		logger.Error(err.Error())
	}

	return &http.Server{
		Addr:           cfg.GetString("SERVER_ADDRESS"),
		ReadTimeout:    time.Duration(cfg.GetInt("SERVER_READ_TIMEOUT") * int(time.Second)),
		WriteTimeout:   time.Duration(cfg.GetInt("SERVER_WRITE_TIMEOUT") * int(time.Second)),
		IdleTimeout:    time.Duration(cfg.GetInt("SERVER_IDLE_TIMEOUT") * int(time.Second)),
		Handler:        routes(),
		MaxHeaderBytes: 1 << 20,
		ErrorLog:       logger.ErrorLog,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{
				cert,
			},
		},
	}

}

func routes() http.Handler {
	r := http.NewServeMux()
	r.Handle("/", Home())
	return r
}

func Home() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the home page!"))

	})
}
