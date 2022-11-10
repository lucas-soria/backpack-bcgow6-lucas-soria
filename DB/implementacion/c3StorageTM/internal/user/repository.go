package user

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lucas-soria/backpack-bcgow6-lucas-soria/internal/domain"
)

const (
	TableName = "Users"
)

type Repository interface {
	Get(ctx *gin.Context, id string) (*domain.User, error)
	GetAll(ctx *gin.Context) (*domain.User, error)
	Store(ctx *gin.Context, user *domain.User) error
	Update(ctx *gin.Context, user *domain.User, id string) error
	Delete(ctx *gin.Context, id string) error
}

type dynamoRepository struct {
	db    *dynamodb.DynamoDB
	table string
}

func NewDynamoRepository(db *dynamodb.DynamoDB) Repository {
	return &dynamoRepository{
		db:    db,
		table: TableName,
	}
}

func (dynamoRepository *dynamoRepository) Get(ctx *gin.Context, id string) (*domain.User, error) {
	result, errGetWithCtx := dynamoRepository.db.GetItemWithContext(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(dynamoRepository.table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})
	if errGetWithCtx != nil {
		return nil, errGetWithCtx
	}
	if result.Item == nil {
		return nil, nil
	}
	return domain.ItemToUser(result.Item)
}

func (dynamoRepository *dynamoRepository) GetAll(ctx *gin.Context) (*domain.User, error) {
	return nil, nil
}

func (dynamoRepository *dynamoRepository) Store(ctx *gin.Context, user *domain.User) error {
	user.ID = uuid.New().String()
	av, errMarshal := dynamodbattribute.MarshalMap(user)
	if errMarshal != nil {
		return errMarshal
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(dynamoRepository.table),
	}
	_, errPutWithCtx := dynamoRepository.db.PutItemWithContext(ctx, input)
	if errPutWithCtx != nil {
		return errPutWithCtx
	}
	return nil
}

func (dynamoRepository *dynamoRepository) Update(ctx *gin.Context, user *domain.User, id string) error {
	return nil
}

func (dynamoRepository *dynamoRepository) Delete(ctx *gin.Context, id string) error {
	_, errDeleteWithCtx := dynamoRepository.db.DeleteItemWithContext(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(dynamoRepository.table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})
	if errDeleteWithCtx != nil {
		return errDeleteWithCtx
	}
	return nil
}
