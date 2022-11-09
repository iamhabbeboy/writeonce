package service

import (
	"bytes"
	"text/template"

	"github.com/google/uuid"
	"github.com/theterminalguy/writeonce/internal/entity"
	twrap "github.com/theterminalguy/writeonce/internal/template"

	repo "github.com/theterminalguy/writeonce/internal/repository"
)

type TemplateService struct {
	TemplateRepo *repo.TemplateRepository
}

func NewTemplateService() *TemplateService {
	return &TemplateService{
		TemplateRepo: repo.NewTemplateRepository(),
	}
}

func (s *TemplateService) CreateTemplate(p repo.TemplateParams) (*entity.Template, error) {
	tmpl, err := template.New(p.Name).Parse(p.Template)
	if err != nil {
		return nil, err
	}
	wrp := twrap.NewWrapper(tmpl)
	p.Placeholders = wrp.ExtractPlaceholders()
	return s.TemplateRepo.Create(p)
}

type GenerateParams struct {
	ProjectID  uuid.UUID         `json:"project_id"`
	TemplateID uuid.UUID         `json:"template_id"`
	Variables  map[string]string `json:"variables"`
}

func (s *TemplateService) Generate(p GenerateParams) (*entity.ParsedTemplate, error) {
	t, err := s.TemplateRepo.Get(p.TemplateID)
	if err != nil {
		return nil, err
	}
	tmpl, err := template.New(t.Name).Parse(t.Template)
	if err != nil {
		return nil, err
	}
	wrp := twrap.NewWrapper(tmpl)
	var dst bytes.Buffer
	if err := wrp.Execute(&dst, p.Variables); err != nil {
		return nil, err
	}
	return &entity.ParsedTemplate{
		Body: dst.String(),
		MetaData: &entity.ParsedTemplateMeta{
			Name: t.Name,
		},
	}, nil
}
