package handler

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	repo "github.com/theterminalguy/writeonce/internal/repository"
	"github.com/theterminalguy/writeonce/internal/service"
)

type TemplateHandler struct {
	TemplateRepo    *repo.TemplateRepository
	TemplateService *service.TemplateService
}

func NewTemplateHandler() *TemplateHandler {
	return &TemplateHandler{
		TemplateRepo: repo.NewTemplateRepository(),
	}
}

func (h *TemplateHandler) Search(c echo.Context) error {
	return nil
}

func (h *TemplateHandler) ReadAll(c echo.Context) error {
	return c.JSON(http.StatusOK, h.TemplateRepo.GetAll())
}

func (h *TemplateHandler) ReadByID(c echo.Context) error {
	id, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	record, err := h.TemplateRepo.Get(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if record == nil {
		return c.String(http.StatusNotFound, "record not found")
	}
	return c.JSON(http.StatusOK, record)
}

func (h *TemplateHandler) CreateOne(c echo.Context) error {
	ct := c.Request().Header.Get(echo.HeaderContentType)

	var params repo.TemplateParams
	if strings.HasPrefix(ct, echo.MIMETextPlain) {
		params.ProjectID = uuid.MustParse(c.FormValue("project_id"))
		params.Name = c.FormValue("name")
		params.Description = c.FormValue("description")

		body := c.Request().Body
		defer body.Close()
		b, err := ioutil.ReadAll(body)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		params.Template = string(b)
	} else if err := c.Bind(&params); err != nil {
		return err
	}
	record, err := h.TemplateService.CreateTemplate(params)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, record)
}

func (h *TemplateHandler) UpdateByID(c echo.Context) error {
	id, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	params := new(repo.TemplateParams)
	if err := c.Bind(params); err != nil {
		return err
	}
	record, err := h.TemplateRepo.Update(id, *params)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, record)
}

func (h *TemplateHandler) DeleteOne(c echo.Context) error {
	id, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	err = h.TemplateRepo.Delete(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusNoContent, nil)
}
