package router

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"note_service/app/internal/app_context"
	"note_service/app/internal/auth"
	"note_service/app/pkg/logging"
	"note_service/app/pkg/metric"
	"note_service/app/pkg/middleware/jwt"
	"note_service/app/pkg/shutdown"
	"os"
	"path"
	"path/filepath"
	"syscall"
	"time"

	"github.com/julienschmidt/httprouter"
)

func Init() {
	logger := logging.Getlogger()
	logger.Println("Router init..")

	router := httprouter.New()

	router.HandlerFunc("POST", auth.URL, auth.Auth)

	//metrics
	router.HandlerFunc("GET", metric.HEARTBEAT_URL, jwt.JWTMiddleware(metric.Heartbeat))
	router.HandlerFunc("GET", metric.TEST_URL, metric.Test)

	ctx := app_context.GetInstance()
	cfg := ctx.Config

	var server *http.Server
	var listener net.Listener

	if cfg.Listen.Type == "sock" {
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}

		socketPath := path.Join(appDir, "app.sock")
		logger.Infof("socket path: %s", socketPath)

		listener, err = net.Listen("unix", socketPath)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		logger.Infof("bind application to host %s and port %s", cfg.Listen.BindIP, cfg.Listen.Port)

		var err error

		listener, err = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
		if err != nil {
			logger.Fatal(err)
		}
	}

	server = &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go shutdown.Graceful([]os.Signal{syscall.SIGABRT, syscall.SIGHUP, syscall.SIGQUIT, os.Interrupt, syscall.SIGTERM}, server)

	if err := server.Serve(listener); err != nil {
		switch {
		case errors.Is(err, http.ErrAbortHandler):
			logger.Warn("server shutdown")
		default:
			logger.Fatal(err)
		}
	}

}
