package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/theterminalguy/writeonce/internal/entity"
	"github.com/theterminalguy/writeonce/internal/repository"
)

type PipelineService struct {
	ts *TemplateService

	PipeRepo *repository.PipeRepository
}

func NewPipelineService() *PipelineService {
	return &PipelineService{
		ts:       NewTemplateService(),
		PipeRepo: repository.NewPipeRepository(),
	}
}

type TemplatePipeParams struct {
	Body     string `json:"body"`
	MetaData struct {
		Name    string   `json:"name"`
		Tags    []string `json:"tags"`
		Version string   `json:"version"`
	} `json:"meta_data"`
}

func (s *PipelineService) Run(pipeline *entity.Pipeline) error {
	t, err := s.ts.Generate(GenerateParams{
		TemplateID: pipeline.Template.ID,
		Variables:  pipeline.Template.Variables,
	})
	if err != nil {
		return err
	}
	for _, pipeInput := range pipeline.Pipes {
		record, err := s.PipeRepo.Get(pipeInput.PipeID)
		if err != nil {
			return err
		}
		reqBody := map[string]interface{}{
			"template":   t,
			"parameters": pipeInput.Parameters,
		}
		bodyBytes, err := json.Marshal(reqBody)
		if err != nil {
			return err
		}
		req, err := http.NewRequest(http.MethodPost, record.Endpoint, bytes.NewBuffer(bodyBytes))
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		fmt.Println(resp)
	}
	return nil
}
