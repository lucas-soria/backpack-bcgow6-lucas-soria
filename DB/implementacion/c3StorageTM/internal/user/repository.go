package user

import (
	"errors"
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

var (
	ErrNotFound = errors.New("user could not be found")
	ErrNoUsers  = errors.New("no users found in database")
)

type Repository interface {
	Get(ctx *gin.Context, id string) (*domain.User, error)
	GetAll(ctx *gin.Context) ([]*domain.User, error)
	Store(ctx *gin.Context, user *domain.User) error
	Update(ctx *gin.Context, user *domain.User) error
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
		return nil, ErrNotFound
	}
	return domain.ItemToUser(result.Item)
}

func (dynamoRepository *dynamoRepository) GetAll(ctx *gin.Context) ([]*domain.User, error) {
	result, errScanWithCtx := dynamoRepository.db.ScanWithContext(ctx, &dynamodb.ScanInput{
		TableName: aws.String(dynamoRepository.table),
	})
	if errScanWithCtx != nil {
		return nil, errScanWithCtx
	}
	if result.Items == nil {
		return nil, ErrNoUsers
	}
	var users []*domain.User
	for _, user := range result.Items {
		userMapped, _ := domain.ItemToUser(user)
		users = append(users, userMapped)
	}
	return users, nil
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

func (dynamoRepository *dynamoRepository) Update(ctx *gin.Context, user *domain.User) error {
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(dynamoRepository.table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(user.ID),
			},
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":f": {
				S: aws.String(user.Firstname),
			},
			":l": {
				S: aws.String(user.Lastname),
			},
			":u": {
				S: aws.String(user.Username),
			},
			":p": {
				S: aws.String(user.Password),
			},
			":e": {
				S: aws.String(user.Email),
			},
			":i": {
				S: aws.String(user.IP),
			},
			":m": {
				S: aws.String(user.MacAddress),
			},
			":w": {
				S: aws.String(user.Website),
			},
			":im": {
				S: aws.String(user.Image),
			},
		},
		UpdateExpression: aws.String("set first_name = :f, last_name = :l, username = :u, password = :p, email = :e, ip = :i, mac_address = :m, website = :w, image = :im"),
	}
	_, errUpdateWithCtx := dynamoRepository.db.UpdateItemWithContext(ctx, input)
	if errUpdateWithCtx != nil {
		return errUpdateWithCtx
	}
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
