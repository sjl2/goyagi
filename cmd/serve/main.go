package main

import (
    "net/http"

    "github.com/sjl2/goyagi/pkg/application"
    "github.com/sjl2/goyagi/pkg/server"
    "github.com/lob/logger-go"
)

func main() {
    log := logger.New()

    app := application.New()

    srv := server.New(app)

    log.Info("server started", logger.Data{"port": app.Config.Port})

    err := srv.ListenAndServe()
    if err != nil && err != http.ErrServerClosed {
        log.Err(err).Fatal("server stopped")
    }

    log.Info("server stopped")
}
