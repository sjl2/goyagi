package application

import (
    "github.com/go-pg/pg"
    "github.com/pkg/errors"
    "github.com/sjl2/goyagi/pkg/config"
    "github.com/sjl2/goyagi/pkg/database"
    "github.com/sjl2/goyagi/pkg/sentry"
)

// App contains necessary references that will be persisted throughout the
// application's lifecycle.
type App struct {
    Config config.Config
    DB     *pg.DB
    Sentry sentry.Sentry
}

// New creates a new instance of App with a Config and DB connection.
func New() (App, error) {
    cfg := config.New()

    db, err := database.New(cfg)
    if err != nil {
        return App{}, errors.Wrap(err, "application")
    }

	sentry, err := sentry.New(cfg)
    if err != nil {
        return App{}, errors.Wrap(err, "application")
    }

    return App{cfg, db, sentry}, nil
}
