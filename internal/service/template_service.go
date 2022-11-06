package service

import (
	"text/template"

	"github.com/theterminalguy/writeonce/internal/entity"
	twrap "github.com/theterminalguy/writeonce/internal/template"

	repo "github.com/theterminalguy/writeonce/internal/repository"
)

type TemplateService struct {
	TeplateRepo *repo.TemplateRepository
}

func NewTemplateService() *TemplateService {
	return &TemplateService{
		TeplateRepo: repo.NewTemplateRepository(),
	}
}

func (s *TemplateService) CreateTemplate(p repo.TemplateParams) (*entity.Template, error) {
	tmpl, err := template.New(p.Name).Parse(p.Template)
	if err != nil {
		return nil, err
	}
	wrp := twrap.NewWrapper(tmpl)
	p.Placeholder = wrp.ExtractPlaceholders()
	return s.TeplateRepo.Create(p)
}
