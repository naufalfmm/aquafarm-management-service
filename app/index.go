package app

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/naufalfmm/aquafarm-management-service/consts"
	"github.com/naufalfmm/aquafarm-management-service/hooks"
	"github.com/naufalfmm/aquafarm-management-service/infrastructures"
	"github.com/naufalfmm/aquafarm-management-service/middlewares"
	"github.com/naufalfmm/aquafarm-management-service/persistents"
	"github.com/naufalfmm/aquafarm-management-service/resources"
	"github.com/naufalfmm/aquafarm-management-service/usecases"
	"github.com/naufalfmm/aquafarm-management-service/utils/generateResp"
)

type App struct {
	Ec          *echo.Echo
	Resources   resources.Resources
	Middlewares middlewares.Middlewares
	Hooks       hooks.Hooks
}

func Init() App {
	ec := echo.New()

	res, err := resources.Init()
	if err != nil {
		panic(err)
	}

	persist, err := persistents.Init(res)
	if err != nil {
		panic(err)
	}

	usec, err := usecases.Init(persist, res)
	if err != nil {
		panic(err)
	}

	middl, err := middlewares.Init(res, usec)
	if err != nil {
		panic(err)
	}

	hookImp, err := hooks.Init(usec, res)
	if err != nil {
		panic(err)
	}

	infr, err := infrastructures.Init(usec, res, middl)
	if err != nil {
		panic(err)
	}

	infr.Register(ec)

	return App{
		Ec:          ec,
		Resources:   res,
		Middlewares: middl,
		Hooks:       hookImp,
	}
}

func (app App) Run() {
	app.Ec.Use(app.Middlewares.PanicRecover(), app.Middlewares.ImplementCors(), middleware.Logger())
	app.Ec.Validator = app.Resources.Validator

	app.Ec.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, generateResp.Default{
			Ok:      true,
			Message: fmt.Sprintf("Welcome to %s", app.Resources.Config.ServiceName),
		})
	})

	echo.NotFoundHandler = func(c echo.Context) error {
		return generateResp.NewJSONResponse(c, http.StatusNotFound, "", consts.ErrPathNotFound)
	}

	go app.Hooks.BulkUpsertEndpoints(app.Ec.AcquireContext())

	if err := app.Ec.Start(fmt.Sprintf(":%d", app.Resources.Config.ServicePort)); err != nil {
		panic(err)
	}
}
