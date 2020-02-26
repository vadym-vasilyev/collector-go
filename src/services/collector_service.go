package services

import (
	"github.com/vadym-vasilyev/collector-go/src/domain/app_record"
	"github.com/vadym-vasilyev/collector-go/src/utils/rest_errors"
)

var (
	CollectorService collectorServiceInterface = &collectorService{}
)

type collectorServiceInterface interface {
	Save(recordsBatch *app_record.RecordsBatch) (*app_record.RecordsBatch, rest_errors.RestErr)
	Get(AppToken string, sessionId string) (*app_record.RecordsBatch, rest_errors.RestErr)
}

type collectorService struct{}

func (c *collectorService) Get(AppToken string, sessionId string) (*app_record.RecordsBatch, rest_errors.RestErr) {
	panic("implement me")
}

func (c *collectorService) Save(recordsBatch *app_record.RecordsBatch) (*app_record.RecordsBatch, rest_errors.RestErr) {
	panic("implement me")
}
