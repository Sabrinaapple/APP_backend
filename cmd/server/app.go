package server

import (
	"context"
	"errors"
	"freefrom.space/cms-service/internal/app"
	"freefrom.space/cms-service/internal/model"
	"freefrom.space/cms-service/internal/pkg/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"net/http"
	"time"
)

func Run() {
	app.New(fx.Invoke(run)).Run()
}

func run(lc fx.Lifecycle, app *gin.Engine, config model.Env) {
	srv := &http.Server{
		Addr:    config.Basic.HttpAddress,
		Handler: app,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Fatalf("listen: %s\n", err)
				} else {
					log.Info("server run success")
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("gracefully shutting down server")
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := srv.Shutdown(ctx); err != nil {
				log.Error("error occurred while gracefully shutting down server")
				return err
			}
			log.Info("graceful server shut down completed")
			return nil
		},
	})
}
