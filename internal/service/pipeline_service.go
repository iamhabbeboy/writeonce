package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
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
		Name string   `json:"name"`
		Tags []string `json:"tags"`
	} `json:"meta_data"`
}

func (s *PipelineService) Run(pipeline *entity.Pipeline) {
	t, err := s.ts.Generate(GenerateParams{
		TemplateID: pipeline.Template.ID,
		Variables:  pipeline.Template.Variables,
	})
	if err != nil {
		return
	}
	var errs []entity.PipeStatus

	for _, pipe := range pipeline.Pipes {
		record, err := s.PipeRepo.Get(pipe.PipeID)
		if err != nil {
			errs = append(errs, entity.PipeStatus{ID: pipe.PipeID, Err: err})
			continue
		}
		// set request body
		// this is a combination of the parsed template the pipes input parameters
		reqBody := map[string]interface{}{
			"template":   t,
			"parameters": pipe.Parameters,
		}
		bodyBytes, err := json.Marshal(reqBody)
		if err != nil {
			errs = append(errs, entity.PipeStatus{ID: pipe.PipeID, Err: err})
			continue
		}
		req, err := http.NewRequest(http.MethodPost, record.Endpoint, bytes.NewBuffer(bodyBytes))
		if err != nil {
			errs = append(errs, entity.PipeStatus{ID: pipe.PipeID, Err: err})
			continue
		}
		req.Header.Set("Content-Type", "application/json")
		// set additional headers
		for k, v := range pipe.Headers {
			req.Header.Set(k, v)
		}
		resp, err := http.DefaultClient.Do(req)

		// first check if the remote service is available
		if resp == nil {
			errs = append(errs, entity.PipeStatus{ID: pipe.PipeID, Err: fmt.Errorf("remote service is not available")})
			continue
		}
		if err != nil {
			errs = append(errs, entity.PipeStatus{ID: pipe.PipeID, Err: err, HTTPStatus: resp.StatusCode})
			continue
		}
		defer resp.Body.Close()
		// check response status code no in 200 range
		if resp.StatusCode < 200 || resp.StatusCode > 299 {
			errs = append(errs,
				entity.PipeStatus{
					ID:         pipe.PipeID,
					Err:        fmt.Errorf("pipe %s returned status code %d", record.Name, resp.StatusCode),
					HTTPStatus: resp.StatusCode,
				})
			continue
		}
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			errs = append(errs, entity.PipeStatus{ID: pipe.PipeID, Err: err})
			continue
		}
		fmt.Println("Response Body:", string(respBody))
	}
	log.Println(errs)
}
