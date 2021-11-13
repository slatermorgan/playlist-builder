package playlists

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// DynamoDBRepository -
type DynamoDBRepository struct {
	session   *dynamodb.DynamoDB
	tableName string
}

// NewDynamoDBRepository -
func NewDynamoDBRepository(ddb *dynamodb.DynamoDB, tableName string) *DynamoDBRepository {
	return &DynamoDBRepository{ddb, tableName}
}

// Get a shift
func (r *DynamoDBRepository) Get(ctx context.Context, id string) (*Shift, error) {
	shift := &Shift{}
	input := &dynamodb.GetItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	}

	result, err := r.session.GetItemWithContext(ctx, input)
	if err != nil {
		return nil, err
	}

	if err := dynamodbattribute.UnmarshalMap(result.Item, &shift); err != nil {
		return nil, err
	}

	return shift, nil
}

// GetAll shifts
func (r *DynamoDBRepository) GetAll(ctx context.Context) ([]*Shift, error) {
	shifts := make([]*Shift, 0)
	result, err := r.session.ScanWithContext(ctx, &dynamodb.ScanInput{
		TableName: aws.String(r.tableName),
	})
	if err != nil {
		return nil, err
	}

	if err := dynamodbattribute.UnmarshalListOfMaps(result.Items, &shifts); err != nil {
		return nil, err
	}

	return shifts, nil
}

type updateShift struct {
	TimeStart  int    `json:":ts"`
	TimeEnd    int    `json:":te"`
	ClientID   string `json:":c"`
	AssignedTo string `json:":a"`
}

type shiftKey struct {
	ID string `json:":id"`
}

// Update a shift
func (r *DynamoDBRepository) Update(ctx context.Context, id string, shift *UpdateShift) error {
	log.Println("id", id)
	update, err := dynamodbattribute.MarshalMap(&updateShift{
		TimeStart:  shift.TimeStart,
		TimeEnd:    shift.TimeEnd,
		ClientID:   shift.ClientID,
		AssignedTo: shift.AssignedTo,
	})
	if err != nil {
		return nil
	}
	input := &dynamodb.UpdateItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		ConditionExpression:       aws.String("attribute_exists(id)"),
		ExpressionAttributeValues: update,
		TableName:                 aws.String(r.tableName),
		UpdateExpression: aws.String(
			"set timeStart = :ts, timeEnd = :te, clientID = :c, assignedTo = :a",
		),
		ReturnValues: aws.String("UPDATED_NEW"),
	}
	_, err = r.session.UpdateItemWithContext(ctx, input)
	return err
}

// Create a shift
func (r *DynamoDBRepository) Create(ctx context.Context, shift *Shift) error {
	item, err := dynamodbattribute.MarshalMap(shift)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(r.tableName),
	}
	_, err = r.session.PutItemWithContext(ctx, input)
	return err
}

// Delete a shift
func (r *DynamoDBRepository) Delete(ctx context.Context, id string) error {
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	}
	_, err := r.session.DeleteItemWithContext(ctx, input)
	return err
}
