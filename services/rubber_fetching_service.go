package services

import "github.com/YaroslavDev/table-tennis-db-crawler/model"

type RubberFetchingService interface {
	FetchRubbers() ([]*model.Rubber, error)
}

