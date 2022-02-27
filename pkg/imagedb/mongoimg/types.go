package mongoimg

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	databaseName         = "shiny_sorter"
	imagesCollectionName = "images"
)

type mongoConnection struct {
	imagesCollection *mongo.Collection
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

	if purge {
		err = client.Database(databaseName).Drop(ctx)
		if err != nil {
			cleanupFunc()
			return nil, nil, err
		}
	}

	imagesCollection := client.Database(databaseName).Collection(imagesCollectionName)

	return &mongoConnection{
		imagesCollection: imagesCollection,
	}, cleanupFunc, nil
}
