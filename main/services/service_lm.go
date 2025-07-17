package services

import (
	"context"
	"main/models"
	"main/repository"

	"gorm.io/gorm"
)

type ServiceLME interface {
	GetLME(ctx context.Context, id string) (models.LicenciaDTO, error)
}
type serviceLME struct {
	repo repository.RepoLm
}

func NewServiceLME(db *gorm.DB) ServiceLME {
	return serviceLME{repository.NewRepoLM(db)}
}


func (s serviceLME) GetLME(ctx context.Context, id string) (models.LicenciaDTO, error) {
	// resp, err := s.repo.GetLME(ctx, id)
	var resp models.LicenciaDTO

	lmedb, err := s.repo.GetLME(ctx, id)

	if err != nil {
		return resp, err
	}

	resp.Id = lmedb.Id

	return resp, nil
}
