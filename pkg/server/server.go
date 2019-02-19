package server

import (
    "fmt"
    "net/http"

    "github.com/labstack/echo"
)

// New returns a new HTTP server with the registered routes.
func New() *http.Server {
    e := echo.New()

    srv := &http.Server{
        Addr:    fmt.Sprintf(":%d", 3000),
        Handler: e,
    }

    return srv
}
