package service

import "github.com/theterminalguy/writeonce/internal/entity"

type GenerateService struct {
}

func NewGenerateService() *GenerateService {
	return &GenerateService{}
}

func (s *GenerateService) GenerateTemplate(t *entity.Template, _ interface{}) (string, error) {
	return "", nil
}
