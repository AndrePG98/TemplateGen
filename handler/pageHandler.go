package handler

import (
	"TemplateGen/views/pages"

	"github.com/labstack/echo/v4"
)

type PageHandler struct {
}

func NewPageHandler() *PageHandler {
	return &PageHandler{}
}

func (ph *PageHandler) HomePage(ctx echo.Context) error {
	return Render(ctx, pages.Index())
}
