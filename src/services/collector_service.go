package services

import (
	"github.com/vadym-vasilyev/collector-go/src/domain/app_record"
	"github.com/vadym-vasilyev/collector-go/src/utils/rest_errors"
)

var (
	CollectorService collectorServiceInterface = &collectorService{}
)

type collectorServiceInterface interface {
	Save(recordsBatch *app_record.RecordsBatch) rest_errors.RestErr
	Get(appToken string, sessionId string) (*app_record.RecordsBatch, rest_errors.RestErr)
}

type collectorService struct{}

func (c *collectorService) Get(appToken string, sessionId string) (*app_record.RecordsBatch, rest_errors.RestErr) {
	record := app_record.Record{
		AppToken:  appToken,
		SessionId: sessionId,
	}
	if records, err := record.FindByExample(); err != nil {
		return nil, err
	} else {
		rb := app_record.RecordsBatch{Records: records}
		return &rb, nil
	}
}

func (c *collectorService) Save(recordsBatch *app_record.RecordsBatch) rest_errors.RestErr {
	return recordsBatch.Save()
}
