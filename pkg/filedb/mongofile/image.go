package mongofile

import (
	"context"
	"fmt"

	"github.com/CrowhopTech/shinysorter/backend/pkg/filedb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetFileByID will get the file with the given ID.
// If not found, will return nil, not an error.
func (mc *mongoConnection) GetFileByID(ctx context.Context, id primitive.ObjectID) (*filedb.File, error) {
	return mc.getFile(ctx, bson.M{
		"_id": id,
	})
}

// GetFileByName will get the file with the given name.
// If not found, will return nil, not an error.
func (mc *mongoConnection) GetFileByName(ctx context.Context, name string) (*filedb.File, error) {
	return mc.getFile(ctx, bson.M{
		"name": name,
	})
}

func (mc *mongoConnection) getFile(ctx context.Context, query bson.M) (*filedb.File, error) {
	res := mc.filesCollection.FindOne(ctx, query)
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, res.Err()
	}

	img := filedb.File{}
	err := res.Decode(&img)
	if err != nil {
		return nil, err
	}

	return &img, nil
}

// ListFiles will list files that match the given filter, if provided.
// If no filter is provided, all results will be returned (oh no).
// If no files match the filter, err will be nil and an empty slice will be returned.
func (mc *mongoConnection) ListFiles(ctx context.Context, filter *filedb.FileFilter) ([]*filedb.File, error) {
	compiledFilter := getQueriesForFilter(filter)

	cursor, err := mc.filesCollection.Find(ctx, compiledFilter, &options.FindOptions{
		Limit: &filter.Limit,
	})
	if err != nil {
		return nil, fmt.Errorf("error while running Find: %v", err)
	}

	results := []*filedb.File{}

	for cursor.Next(ctx) {
		var result filedb.File
		err := cursor.Decode(&result)
		if err != nil {
			return nil, fmt.Errorf("error while running Decode: %v", err)
		}
		results = append(results, &result)
	}

	return results, nil
}

// CreateFileEntry will create the given entry in the database.
// If one already exists with the given name, this will check for conflicts
// using ConflictsWith. If there is a conflict, an error will be returned.
// If not, no action will be taken.
func (mc *mongoConnection) CreateFileEntry(ctx context.Context, i *filedb.File) (primitive.ObjectID, error) {
	// TODO: filter for valid name characters here! (mainly need to restrict colons (:) and pipes (|) for tagging query purposes)
	existingImg, err := mc.GetFileByName(ctx, i.Name)
	if err != nil {
		return primitive.NilObjectID, err
	}

	// TODO: validate length and characters of md5sum, and enforce case

	if existingImg == nil {
		// Doesn't exist, let's just create it
		count, err := mc.filesCollection.CountDocuments(ctx, bson.M{})
		if err != nil {
			return primitive.NilObjectID, fmt.Errorf("failed to get document count: %v", err)
		}
		if count >= mc.maxFiles {
			return primitive.NilObjectID, fmt.Errorf("the maximum number of files (%d) have been inserted", mc.maxFiles)
		}
		i.ID = primitive.NewObjectID()

		res, err := mc.filesCollection.InsertOne(ctx, i)
		if err != nil {
			return primitive.NilObjectID, fmt.Errorf("failed to insert file: %v", err)
		}
		insertedID := res.InsertedID.(primitive.ObjectID)
		return insertedID, err
	}

	err = i.ConflictsWith(existingImg)
	if err != nil {
		return primitive.NilObjectID, err
	}

	// Already exists, success depends on if the existing file conflicts with the new one
	return existingImg.ID, nil
}

// getUpdateParameter will return the update parameter for the given file.
// It's extracted here for easier unit testing.
func (mc *mongoConnection) getUpdateParameter(i *filedb.File) (bson.M, error) {
	setParams := bson.M{}

	if i == nil {
		return bson.M{}, nil
	}

	if len(i.Md5Sum) > 0 {
		// TODO: validate length and characters of md5sum, and enforce case
		setParams["md5sum"] = i.Md5Sum
	}

	if i.Tags != nil {
		setParams["tags"] = i.Tags
	}

	if i.HasBeenTagged != nil {
		setParams["hasBeenTagged"] = *i.HasBeenTagged
	}

	if i.HasContent != nil {
		setParams["hasContent"] = *i.HasContent
	}

	if i.MIMEType != "" {
		setParams["mimeType"] = i.MIMEType
	}

	return bson.M{"$set": setParams}, nil
}

func (mc *mongoConnection) ModifyFileEntry(ctx context.Context, i *filedb.File) (*filedb.File, error) {
	// Check name length
	update, err := mc.getUpdateParameter(i)
	if err != nil {
		return nil, fmt.Errorf("invalid file update provided (%v): %v", i, err)
	}

	_, err = mc.filesCollection.UpdateByID(ctx, i.ID, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update file %s: %v", i.Name, err)
	}

	modifiedFile, err := mc.GetFileByID(ctx, i.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get modified file %s: %v", i.Name, err)
	}

	return modifiedFile, nil
}
