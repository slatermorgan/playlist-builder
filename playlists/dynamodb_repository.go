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

// Get a playlist
func (r *DynamoDBRepository) Get(ctx context.Context, id string) (*Playlist, error) {
	playlist := &Playlist{}
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

	if err := dynamodbattribute.UnmarshalMap(result.Item, &playlist); err != nil {
		return nil, err
	}

	return playlist, nil
}

// GetAll playlists
func (r *DynamoDBRepository) GetAll(ctx context.Context) ([]*Playlist, error) {
	playlists := make([]*Playlist, 0)
	result, err := r.session.ScanWithContext(ctx, &dynamodb.ScanInput{
		TableName: aws.String(r.tableName),
	})
	if err != nil {
		return nil, err
	}

	if err := dynamodbattribute.UnmarshalListOfMaps(result.Items, &playlists); err != nil {
		return nil, err
	}

	return playlists, nil
}

type updatePlaylist struct {
	TimeStart  int    `json:":ts"`
	TimeEnd    int    `json:":te"`
	ClientID   string `json:":c"`
	AssignedTo string `json:":a"`
}

type playlistKey struct {
	ID string `json:":id"`
}

// Update a playlist
func (r *DynamoDBRepository) Update(ctx context.Context, id string, playlist *UpdatePlaylist) error {
	log.Println("id", id)
	update, err := dynamodbattribute.MarshalMap(&updatePlaylist{
		TimeStart:  playlist.TimeStart,
		TimeEnd:    playlist.TimeEnd,
		ClientID:   playlist.ClientID,
		AssignedTo: playlist.AssignedTo,
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

// Create a playlist
func (r *DynamoDBRepository) Create(ctx context.Context, playlist *Playlist) error {
	item, err := dynamodbattribute.MarshalMap(playlist)
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

// Delete a playlist
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
