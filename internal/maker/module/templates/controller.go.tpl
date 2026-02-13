package {{ .Name }}

import "github.com/labstack/echo/v4"

type Controller struct {
	service *Service
}

func NewController(e *echo.Echo, service *Service) *Controller {
	controller := &Controller{service: service}
	api.RegisterRoutes(e, controller)
	return controller
}
