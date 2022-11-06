package service

import (
	"text/template"

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
