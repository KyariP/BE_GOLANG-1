package routes

import (
	"dumbmerch/handler"
	"dumbmerch/pkg/middleware"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repositories"

	"github.com/labstack/echo/v4"
)

func TripRoutes(e *echo.Group) {
	TripRepository := repositories.RepositoryTrip(mysql.DB)
	h := handler.HandlerTrip(TripRepository)
	e.GET("/trip", h.FindTrip)
	e.GET("/trip/:id", h.FindTripId)
	e.DELETE("/trip/:id", h.DeleteTrip)
	e.POST("/trip", middleware.UploadFile(h.CreateTrip))
	e.PATCH("/trip/:id", h.UpdateTrip)
}
