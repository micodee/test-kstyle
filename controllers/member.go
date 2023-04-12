package controllers

import (
	"net/http"
	"project/dto"
	"project/models"
	"project/repositories"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type memberControl struct {
	MemberRepository repositories.MemberRepository
}

func ControlMember(MemberRepository repositories.MemberRepository) *memberControl {
	return &memberControl{MemberRepository}
}

func (h *memberControl) FindMember(c echo.Context) error {
	members, err := h.MemberRepository.FindMember()
	if err != nil {
		return c.JSON(http.StatusBadRequest, "error Bad Request")
	}
	return c.JSON(http.StatusOK, members)
}

func (h *memberControl) GetMember(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	member, err := h.MemberRepository.GetMember(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "error Bad Request")
	}
	return c.JSON(http.StatusOK, member)
}

func (h *memberControl) CreateMember(c echo.Context) error {
	request := new(dto.MemberRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, "error Bad Request")
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "error Bad Request Validation")
	}

	member := models.Member{
		Username:  request.Username,
		Gender:    request.Gender,
		SkinType:  request.SkinType,
		SkinColor: request.SkinColor,
	}

	data, err := h.MemberRepository.CreateMember(member)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "error Internal Server")
	}
	return c.JSON(http.StatusCreated, respMember(data))
}

func (h *memberControl) UpdateMember(c echo.Context) error {
	request := new(dto.MemberRequestUpdate)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, "error Bad Request")
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "error Bad Request Validation")
	}

	id, _ := strconv.Atoi(c.Param("id"))
	member, err := h.MemberRepository.GetMember(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "error Bad Request Get ID")
	}

	if request.Username != "" {
		member.Username = request.Username
	}
	if request.Gender != "" {
		member.Gender = request.Gender
	}
	if request.SkinType != "" {
		member.SkinType = request.SkinType
	}
	if request.SkinColor != "" {
		member.SkinColor = request.SkinColor
	}

	data, err := h.MemberRepository.UpdateMember(member, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "error Internal Server")
	}
	return c.JSON(http.StatusCreated, respMember(data))
}

func (h *memberControl) DeleteMember(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	member, err := h.MemberRepository.GetMember(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "error Bad Request")
	}

	data, err := h.MemberRepository.DeleteMember(member, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "error Internal Server")
	}
	return c.JSON(http.StatusOK, respMember(data))
}

func respMember(u models.Member) dto.MemberResponse {
	return dto.MemberResponse{
		IDMember:  u.IDMember,
		Username:  u.Username,
		Gender:    u.Gender,
		SkinType:  u.SkinType,
		SkinColor: u.SkinColor,
	}
}
