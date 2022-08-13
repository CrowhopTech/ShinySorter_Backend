package mongofile

import (
	"context"
	"fmt"
	"os"
	"strconv"

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

	maxFilesEnv     = "MAX_FILE_COUNT"
	maxTagsEnv      = "MAX_TAG_COUNT"
	maxQuestionsEnv = "MAX_QUESTION_COUNT"

	defaultMaxFilesCount     = 5
	defaultMaxTagsCount      = 5
	defaultMaxQuestionsCount = 5
)

var _ filedb.FileMetadataService = new(mongoConnection)

type mongoConnection struct {
	client              *mongo.Client
	filesCollection     *mongo.Collection
	tagsCollection      *mongo.Collection
	questionsCollection *mongo.Collection

	maxFiles     int64
	maxTags      int64
	maxQuestions int64
}

func getCollectionLimits() (int64, int64, int64) {
	filesVar, _ := os.LookupEnv(maxFilesEnv)
	tagsVar, _ := os.LookupEnv(maxTagsEnv)
	questionsVar, _ := os.LookupEnv(maxQuestionsEnv)

	var (
		maxFiles     = int64(defaultMaxFilesCount)
		maxTags      = int64(defaultMaxTagsCount)
		maxQuestions = int64(defaultMaxQuestionsCount)
		err          error
	)

	if filesVar != "" {
		maxFiles, err = strconv.ParseInt(filesVar, 10, 64)
		if err != nil {
			logrus.Panicf("Unable to parse max files count '%s' as an int", filesVar)
		}
	}
	if tagsVar != "" {
		maxTags, err = strconv.ParseInt(tagsVar, 10, 64)
		if err != nil {
			logrus.Panicf("Unable to parse max tags count '%s' as an int", tagsVar)
		}
	}
	if questionsVar != "" {
		maxQuestions, err = strconv.ParseInt(questionsVar, 10, 64)
		if err != nil {
			logrus.Panicf("Unable to parse max questions count '%s' as an int", questionsVar)
		}
	}

	return maxFiles, maxTags, maxQuestions
}

func (mc *mongoConnection) setUpIndices(ctx context.Context) error {
	unique := true
	_, err := mc.tagsCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
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

	maxFiles, maxTags, maxQuestions := getCollectionLimits()

	mc := &mongoConnection{
		client:              client,
		filesCollection:     client.Database(databaseName).Collection(filesCollectionName),
		tagsCollection:      client.Database(databaseName).Collection(tagsCollectionName),
		questionsCollection: client.Database(databaseName).Collection(questionsCollectionName),
		maxFiles:            maxFiles,
		maxTags:             maxTags,
		maxQuestions:        maxQuestions,
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
