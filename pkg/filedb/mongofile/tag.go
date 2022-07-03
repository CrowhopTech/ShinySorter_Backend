package mongofile

import (
	"context"
	"fmt"

	"github.com/CrowhopTech/shinysorter/backend/pkg/filedb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ListTags will return the list of all tags. There are no filter options as this
// list will never be extremely large.
func (mc *mongoConnection) ListTags(ctx context.Context) ([]*filedb.Tag, error) {
	cursor, err := mc.tagsCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("error while running Find: %v", err)
	}

	results := []*filedb.Tag{}

	for cursor.Next(ctx) {
		var result filedb.Tag
		err := cursor.Decode(&result)
		if err != nil {
			return nil, fmt.Errorf("error while running Decode: %v", err)
		}
		results = append(results, &result)
	}

	return results, nil
}

func (mc *mongoConnection) getNewTagID(ctx context.Context) (int64, error) {
	// Find the highest existing ID, then set the ID to one higher
	highestResCursor := mc.tagsCollection.FindOne(ctx, bson.M{}, &options.FindOneOptions{
		Sort: bson.M{
			"_id": -1,
		},
	})

	highestID := int64(0)
	highestRes := filedb.Tag{}

	err := highestResCursor.Decode(&highestRes)
	if err != mongo.ErrNoDocuments {
		if err != nil {
			return -1, fmt.Errorf("failed to find highest existing tag ID: %v", err)
		}

		highestID = highestRes.ID
	}

	return highestID + 1, nil
}

// CreateTag will define a new tag for use on files. ID will be auto-set, if a value
// is provided it will not be used.
func (mc *mongoConnection) CreateTag(ctx context.Context, t *filedb.Tag) (*filedb.Tag, error) {
	// TODO: validate tag values

	newTagID, err := mc.getNewTagID(ctx)
	if err != nil {
		return nil, fmt.Errorf("error while getting new tag ID: %v", err)
	}

	t.ID = newTagID

	res, err := mc.tagsCollection.InsertOne(ctx, *t)
	if err != nil {
		return nil, fmt.Errorf("error while inserting tag: %v", err)
	}

	created := mc.tagsCollection.FindOne(ctx, bson.M{"_id": res.InsertedID})
	if created.Err() != nil {
		return nil, fmt.Errorf("error while fetching created tag: %v", created.Err())
	}

	createdTag := filedb.Tag{}
	if err := created.Decode(&createdTag); err != nil {
		return nil, fmt.Errorf("error while decoding created tag: %v", err)
	}

	return &createdTag, nil
}

// ModifyTag will modify the given tag with the new given information.
func (mc *mongoConnection) ModifyTag(ctx context.Context, t *filedb.Tag) (*filedb.Tag, error) {
	trueVal := true

	// TODO: validate that object already exists (we can currently create tags using this call...)

	setParams := bson.M{}

	if len(t.Name) > 0 {
		setParams["name"] = t.Name
	}
	if len(t.UserFriendlyName) > 0 {
		setParams["userFriendlyName"] = t.UserFriendlyName
	}
	if len(t.Description) > 0 {
		setParams["description"] = t.Description
	}

	res, err := mc.tagsCollection.UpdateOne(ctx, bson.M{"_id": t.ID}, bson.M{"$set": setParams}, &options.UpdateOptions{
		Upsert: &trueVal,
	})
	if err != nil {
		return nil, fmt.Errorf("error while updating tag: %v", err)
	}

	// TODO: better distinguish between "ID didn't exist" and "document matched original"
	if res.ModifiedCount == 0 {
		return nil, fmt.Errorf("no documents updated")
	}

	updated := mc.tagsCollection.FindOne(ctx, bson.M{"_id": t.ID})
	if updated.Err() != nil {
		return nil, fmt.Errorf("error while fetching updated tag: %v", updated.Err())
	}

	updatedTag := filedb.Tag{}
	if err := updated.Decode(&updatedTag); err != nil {
		return nil, fmt.Errorf("error while decoding updated tag: %v", err)
	}

	return &updatedTag, nil
}

// DeleteTag will delete the given tag, as well as remove it from all files that reference it.
func (mc *mongoConnection) DeleteTag(ctx context.Context, id int64) error {
	res, err := mc.tagsCollection.DeleteOne(ctx, bson.M{"$_id": id})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("no documents deleted")
	}
	return nil
}
