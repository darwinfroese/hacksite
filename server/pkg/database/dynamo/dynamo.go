package dynamo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/darwinfroese/hacksite/server/models"
	"github.com/darwinfroese/hacksite/server/pkg/database"
)

// table names
const (
	accountsTable = "hacksite-accounts"
	projectsTable = "hacksite-projects"
	sessionsTable = "hacksite-sessions"
	emailsTable   = "hacksite-emails"
)

type emailEntry struct {
	Email, Username string
}

type dynamoDB struct {
	db *dynamodb.DynamoDB
}

// New creates a new dynamoDB connection
func New(config *aws.Config) database.Database {
	return &dynamoDB{
		db: dynamodb.New(session.New(), config),
	}
}

func (d *dynamoDB) AddProject(project models.Project) error {
	return putItem(d.db, project, projectsTable)
}

func (d *dynamoDB) GetProject(id string) (models.Project, error) {
	var project models.Project

	input := &dynamodb.GetItemInput{
		TableName: aws.String(projectsTable),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(id),
			},
		},
	}

	result, err := d.db.GetItem(input)
	if err != nil {
		return project, err
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &project)

	return project, err
}

func (d *dynamoDB) UpdateProject(project models.Project) error {
	return putItem(d.db, project, projectsTable)
}

func (d *dynamoDB) RemoveProject(id string) error {
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(id),
			},
		},
		TableName: aws.String(projectsTable),
	}

	_, err := d.db.DeleteItem(input)

	return err
}

func (d *dynamoDB) CreateAccount(account models.Account) error {
	err := putItem(d.db, account, accountsTable)
	if err != nil {
		return err
	}

	entry := emailEntry{
		Username: account.Username,
		Email:    account.Email,
	}
	return putItem(d.db, entry, emailsTable)
}

func (d *dynamoDB) GetAccountByUsername(username string) (models.Account, error) {
	var account models.Account

	input := &dynamodb.GetItemInput{
		TableName: aws.String(accountsTable),
		Key: map[string]*dynamodb.AttributeValue{
			"Username": {
				S: aws.String(username),
			},
		},
	}

	result, err := d.db.GetItem(input)
	if err != nil {
		return account, err
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &account)

	return account, nil
}

func (d *dynamoDB) GetAccountByEmail(email string) (models.Account, error) {
	var entry emailEntry
	var account models.Account

	input := &dynamodb.GetItemInput{
		TableName: aws.String(emailsTable),
		Key: map[string]*dynamodb.AttributeValue{
			"Email": {
				S: aws.String(email),
			},
		},
	}

	result, err := d.db.GetItem(input)
	if err != nil {
		return account, err
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &entry)
	if err != nil {
		return account, err
	}

	// if we don't have a username, we don't have an account
	if entry.Username == "" {
		return account, nil
	}

	return d.GetAccountByUsername(entry.Username)
}

func (d *dynamoDB) UpdateAccount(account models.Account) error {
	return putItem(d.db, account, accountsTable)
}

func (d *dynamoDB) StoreSession(session models.Session) error {
	return putItem(d.db, session, sessionsTable)
}

func (d *dynamoDB) GetSession(sessionToken string) (models.Session, error) {
	var session models.Session

	input := &dynamodb.GetItemInput{
		TableName: aws.String(sessionsTable),
		Key: map[string]*dynamodb.AttributeValue{
			"Token": {
				S: aws.String(sessionToken),
			},
		},
	}

	result, err := d.db.GetItem(input)
	if err != nil {
		return session, err
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &session)

	return session, nil
}

func (d *dynamoDB) GetAllSessions() ([]models.Session, error) {
	return nil, nil
}

func (d *dynamoDB) RemoveSession(sessionToken string) error {
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Token": {
				S: aws.String(sessionToken),
			},
		},
		TableName: aws.String(sessionsTable),
	}

	_, err := d.db.DeleteItem(input)

	return err
}

func putItem(db *dynamodb.DynamoDB, item interface{}, tableName string) error {
	attr, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      attr,
		TableName: aws.String(tableName),
	}

	_, err = db.PutItem(input)

	return err
}
