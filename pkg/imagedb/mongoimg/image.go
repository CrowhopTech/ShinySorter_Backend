package mongoimg

import (
	"context"
	"fmt"

	"github.com/CrowhopTech/shinysorter/backend/pkg/imagedb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetImage will get the image with the given name.
// If not found, will return nil, not an error.
func (mc *mongoConnection) GetImage(ctx context.Context, name string) (*imagedb.Image, error) {
	res := mc.imagesCollection.FindOne(ctx, bson.M{
		"_id": name,
	})
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, res.Err()
	}

	img := imagedb.Image{}
	err := res.Decode(&img)
	if err != nil {
		return nil, err
	}

	return &img, nil
}

// ListImages will list images that match the given filter, if provided.
// If no filter is provided, all results will be returned (oh no).
// If no images match the filter, err will be nil and an empty slice will be returned.
func (mc *mongoConnection) ListImages(ctx context.Context, filter *imagedb.ImageFilter) ([]*imagedb.Image, error) {
	compiledFilter := getQueriesForFilter(filter)

	cursor, err := mc.imagesCollection.Find(ctx, compiledFilter)
	if err != nil {
		return nil, fmt.Errorf("error while running Find: %v", err)
	}

	results := []*imagedb.Image{}

	for cursor.Next(ctx) {
		var result imagedb.Image
		err := cursor.Decode(&result)
		if err != nil {
			return nil, fmt.Errorf("error while running Decode: %v", err)
		}
		results = append(results, &result)
	}

	return results, nil
}

// CreateImageEntry will create the given entry in the database.
// If one already exists with the given name, this will check for conflicts
// using ConflictsWith. If there is a conflict, an error will be returned.
// If not, no action will be taken.
func (mc *mongoConnection) CreateImageEntry(ctx context.Context, i *imagedb.Image) error {
	// TODO: filter for valid name characters here! (mainly need to restrict colons (:) and pipes (|) for tagging query purposes)
	existingImg, err := mc.GetImage(ctx, i.Name)
	if err != nil {
		return err
	}

	// TODO: validate length and characters of md5sum, and enforce case

	if existingImg == nil {
		// Doesn't exist, let's just create it
		_, err = mc.imagesCollection.InsertOne(ctx, i)
		return err
	}

	// Already exists, success depends on if the existing image conflicts with the new one
	return i.ConflictsWith(existingImg)
}

func (mc *mongoConnection) ModifyImageEntry(ctx context.Context, i *imagedb.Image) (*imagedb.Image, error) {
	// Check name length
	if len(i.Name) == 0 {
		return nil, fmt.Errorf("name empty in provided image")
	}

	setParams := bson.M{}

	if len(i.Md5Sum) > 0 {
		// TODO: validate length and characters of md5sum, and enforce case
		setParams["md5sum"] = i.Md5Sum
	}

	if i.Tags != nil {
		setParams["tags"] = i.Tags
	}

	_, err := mc.imagesCollection.UpdateByID(ctx, i.Name, bson.M{
		"$set": setParams,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update image %s: %v", i.Name, err)
	}

	modifiedImage, err := mc.GetImage(ctx, i.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to get modified image %s: %v", i.Name, err)
	}

	return modifiedImage, nil
}
