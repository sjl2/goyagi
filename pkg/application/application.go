package application

import (
    "github.com/sjl2/goyagi/pkg/config"
)

// App contains necessary references that will be persisted throughout the
// application's lifecycle.
type App struct {
    Config config.Config
}

// New creates a new instance of App
func New() App {
    cfg := config.New()

    return App{cfg}
}
