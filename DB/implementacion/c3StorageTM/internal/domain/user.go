package domain

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type User struct {
	ID         string `json:"id"`
	Firstname  string `json:"first_name"`
	Lastname   string `json:"last_name"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	IP         string `json:"ip"`
	MacAddress string `json:"mac_address"`
	Website    string `json:"website"`
	Image      string `json:"image"`
}

func ItemToUser(input map[string]*dynamodb.AttributeValue) (*User, error) {
	var item User
	err := dynamodbattribute.UnmarshalMap(input, &item)
	if err != nil {
		return nil, err
	}
	return &item, nil
}
