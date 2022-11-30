package mongofile

import (
	"context"
	"fmt"

	"github.com/CrowhopTech/shinysorter/backend/pkg/filedb"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	databaseName            = "shiny_sorter"
	filesCollectionName     = "files"
	tagsCollectionName      = "tags"
	questionsCollectionName = "questions"
)

var _ filedb.FileMetadataService = new(mongoConnection)

type mongoConnection struct {
	client              *mongo.Client
	filesCollection     *mongo.Collection
	tagsCollection      *mongo.Collection
	questionsCollection *mongo.Collection
}

func (mc *mongoConnection) setUpIndices(ctx context.Context) error {
	unique := true
	_, err := mc.filesCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{
			{
				Key:   "name",
				Value: 1,
			},
		},
		Options: &options.IndexOptions{
			Unique: &unique,
		},
	})
	if err != nil {
		return err
	}

	_, err = mc.tagsCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{
			{
				Key:   "userFriendlyName",
				Value: 1,
			},
		},
		Options: &options.IndexOptions{
			Unique: &unique,
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func New(ctx context.Context, connectionURI string, purge bool) (*mongoConnection, func() error, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionURI))
	if err != nil {
		return nil, nil, err
	}

	cleanupFunc := func() error {
		err := client.Disconnect(ctx)
		if err != nil {
			logrus.WithError(err).Error("Failed to clean up mongodb connection")
		}
		return err
	}

	mc := &mongoConnection{
		client:              client,
		filesCollection:     client.Database(databaseName).Collection(filesCollectionName),
		tagsCollection:      client.Database(databaseName).Collection(tagsCollectionName),
		questionsCollection: client.Database(databaseName).Collection(questionsCollectionName),
	}

	if purge {
		err = mc.filesCollection.Drop(ctx)
		if err != nil {
			cleanupFunc()
			return nil, nil, err
		}
		err = mc.tagsCollection.Drop(ctx)
		if err != nil {
			cleanupFunc()
			return nil, nil, err
		}
		err = mc.questionsCollection.Drop(ctx)
		if err != nil {
			cleanupFunc()
			return nil, nil, err
		}
	}

	err = mc.setUpIndices(ctx)
	if err != nil {
		cleanupFunc()
		return nil, nil, fmt.Errorf("failed to set up tag indices: %v", err)
	}

	return mc, cleanupFunc, nil
}
