package routes

import (
	"project/controllers"
	"project/database"
	"project/repositories"

	"github.com/labstack/echo/v4"
)

func MemberRoutes(e *echo.Group) {
	memberRepository := repositories.RepositoryMember(database.ConnDB)
	h := controllers.ControlMember(memberRepository)

	e.GET("/members", h.FindMember)
	e.GET("/member/:id", h.GetMember)
	e.POST("/member", h.CreateMember)
	e.PATCH("/member/:id", h.UpdateMember)
	e.DELETE("/member/:id", h.DeleteMember)
}
