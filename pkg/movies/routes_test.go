package movies

import (
    "testing"

    "github.com/labstack/echo"
    "github.com/sjl2/goyagi/internal/test"
    "github.com/sjl2/goyagi/pkg/application"
    "github.com/stretchr/testify/assert"
)

func TestRegisterRoutes(t *testing.T) {
    e := echo.New()
    app := application.App{}

    RegisterRoutes(e, app)

    routes := test.FilterRoutes(e.Routes())
    assert.Len(t, routes, 3)
    assert.Contains(t, routes, "GET /movies")
    assert.Contains(t, routes, "GET /movies/:id")
    assert.Contains(t, routes, "POST /movies")
}
