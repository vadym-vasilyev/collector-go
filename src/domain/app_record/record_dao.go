package app_record

import (
	"context"
	"fmt"
	"github.com/vadym-vasilyev/collector-go/src/client/mongo_client"
	"github.com/vadym-vasilyev/collector-go/src/utils/logger"
	"github.com/vadym-vasilyev/collector-go/src/utils/rest_errors"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	mongoCollectionName = "record"
)

func (recordBatch *RecordsBatch) Save() rest_errors.RestErr {
	var ptr []interface{}
	for _, t := range recordBatch.Records {
		ptr = append(ptr, &t)
	}
	if err := mongo_client.Client.SaveMany(mongoCollectionName, ptr); err != nil {
		return rest_errors.NewInternalServerError("Can't fetch records", nil)
	}
	return nil
}

func (record *Record) Save() rest_errors.RestErr {
	if err := mongo_client.Client.Save(mongoCollectionName, record); err != nil {
		return rest_errors.NewInternalServerError("Can't fetch records", nil)
	}
	return nil
}

func (record *Record) FindByTokenAndSession() ([]Record, rest_errors.RestErr) {
	filter := bson.M{"AppToken": record.AppToken, "SessionId": record.SessionId}
	if cursor, err := mongo_client.Client.Fetch(mongoCollectionName, filter); err != nil {
		return nil, rest_errors.NewInternalServerError("Can't fetch records", err)
	} else {
		var fetchedRecords []Record
		if err = cursor.All(context.Background(), &fetchedRecords); err != nil {
			logger.Error(fmt.Sprintf("Get an error during retrieving all finding documents for %#v", record), err)
			return nil, rest_errors.NewInternalServerError("Cat't fetch records", err)
		}
		return fetchedRecords, nil
	}
}
