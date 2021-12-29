package mongoimg

import (
	"context"

	"github.com/CrowhopTech/shinysorter/backend/pkg/imagedb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ imagedb.ImageMetadata = new(mongoConnection)

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

// CreateImageEntry will create the given entry in the database.
// If one already exists with the given name, this will check for conflicts
// using ConflictsWith. If there is a conflict, an error will be returned.
// If not, no action will be taken.
func (mc *mongoConnection) CreateImageEntry(ctx context.Context, i *imagedb.Image) error {
	existingImg, err := mc.GetImage(ctx, i.Name)
	if err != nil {
		return err
	}

	if existingImg == nil {
		// Doesn't exist, let's just create it
		_, err = mc.imagesCollection.InsertOne(ctx, i)
		return err
	}

	// Already exists, success depends on if the existing image conflicts with the new one
	return i.ConflictsWith(existingImg)
}
