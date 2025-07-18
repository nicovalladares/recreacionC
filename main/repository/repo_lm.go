package repository

import (
	"context"
	"main/models"

	"gorm.io/gorm"
)

type RepoLm interface{
	GetLME(ctx context.Context, id string) (models.LmLicencia, error)
}

type repoLm struct {
	db *gorm.DB
}

func NewRepoLM(db *gorm.DB) RepoLm {
	return repoLm{db}
}

func (r repoLm) GetLME(ctx context.Context, id string) (models.LmLicencia, error) {
	var resp models.LmLicencia
	query := `SELECT id, numero FROM LM.LM_LICENCIA WHERE id = ?`


	db := r.db.WithContext(ctx)

	err := db.Raw(query, id).Find(&resp).Error
	if err != nil {
		return resp, err
	}
	return resp, nil

}