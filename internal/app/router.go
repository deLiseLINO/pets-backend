package app

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"pets-backend/internal/handler/login"
	"pets-backend/internal/services"
	"pets-backend/pkg/httpserver"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (app *App) initRouter(
	optSvc *services.OtpService,
	userSvc *services.UserService,
	ssoSvc *services.SSOService,
) http.Handler {
	router := gin.Default()

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.WithFields(log.Fields{
			"method":       httpMethod,
			"absolutePath": absolutePath,
			"handlerName":  handlerName,
			"nuHandlers":   nuHandlers,
		}).Debug()
	}

	router.POST("/login/otp_generate", login.HandleGenerateOtp(optSvc))
	router.POST("/login/by_code", login.HandleByCode(optSvc, userSvc, ssoSvc))
	router.POST("/login/signup", login.HandleSignUp(userSvc, optSvc))

	return router
}

func (app *App) runServer(handler http.Handler) {
	readTimeout := time.Second * time.Duration(app.cfg.HTTPServer.ReadTimeout)
	writeTimeout := time.Second * time.Duration(app.cfg.HTTPServer.WriteTimeout)
	shutdownDuration := time.Second * time.Duration(app.cfg.HTTPServer.ShutdownTimeout)

	httpServer := httpserver.New(
		handler,
		httpserver.Addr(app.cfg.HTTPServer.ListenAddress()),
		httpserver.ReadTimeout(readTimeout),
		httpserver.WriteTimeout(writeTimeout),
		httpserver.ShutdownTimeout(shutdownDuration),
	)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		log.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err := httpServer.Shutdown()
	if err != nil {
		log.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

	log.Info("server shutdown completed")
}
