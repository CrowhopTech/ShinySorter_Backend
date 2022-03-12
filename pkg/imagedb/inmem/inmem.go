package inmem

import (
	"context"
	"sync"

	"github.com/CrowhopTech/shinysorter/backend/pkg/imagedb"
)

var _ imagedb.ImageMetadata = new(imageDB)

var dummyImageData []imagedb.Image = []imagedb.Image{
	{
		FileMetadata: imagedb.FileMetadata{
			Name:   "dummy1.jpeg",
			Md5Sum: "notahash",
		},
		Tags: map[string]string{
			"test-tag": "testval",
		},
	},
	{
		FileMetadata: imagedb.FileMetadata{
			Name:   "dummy2.jpeg",
			Md5Sum: "stillnotahash",
		},
		Tags: map[string]string{
			"test-tag":       "otherval",
			"test-other-tag": "",
		},
	},
	{
		FileMetadata: imagedb.FileMetadata{
			Name:   "dummy3.jpeg",
			Md5Sum: "mostdefinitelynotahash",
		},
		Tags: map[string]string{
			"test-other-tag": "testval",
		},
	},
}

type imageDB struct {
	images     []imagedb.Image
	imagesLock sync.RWMutex
}

func New(dummyData bool) *imageDB {
	tr := imageDB{}
	if dummyData {
		tr.images = dummyImageData
	}
	return &tr
}

// GetImage will get the image with the given name.
// If not found, will return nil, not an error.
func (im *imageDB) GetImage(ctx context.Context, name string) (*imagedb.Image, error) {
	im.imagesLock.RLock()
	defer im.imagesLock.RUnlock()

	for _, img := range im.images {
		// Check if context is done, if so, return context cancelled
		select {
		case <-ctx.Done():
			return nil, context.Canceled
		default:
			if img.Name == name {
				return img.Clone(), nil
			}
		}
	}
	return nil, nil
}

// ListImages will list images that match the given filter, if provided.
// If no filter is provided, all results will be returned (oh no).
// If no images match the filter, err will be nil and an empty slice will be returned.
func (im *imageDB) ListImages(ctx context.Context, filter *imagedb.ImageFilter) ([]*imagedb.Image, error) {
	im.imagesLock.RLock()
	defer im.imagesLock.RUnlock()

	matchingImages := []*imagedb.Image{}

	for _, img := range im.images {
		// Check if context is done, if so, return context cancelled
		select {
		case <-ctx.Done():
			return nil, context.Canceled
		default:
			if filter != nil {
				if filter.Name != nil && img.Name != *filter.Name {
					continue
				}
				if filter.Md5Sum != nil && img.Md5Sum != *filter.Md5Sum {
					continue
				}
			}
			matchingImages = append(matchingImages, img.Clone())
		}
	}
	return matchingImages, nil
}

// CreateImageEntry will create the given entry in the database.
// If one already exists with the given name, this will check for conflicts
// using ConflictsWith. If there is a conflict, an error will be returned.
// If not, no action will be taken.
func (im *imageDB) CreateImageEntry(ctx context.Context, i *imagedb.Image) error {
	existingImg, err := im.GetImage(ctx, i.Name)
	if err != nil {
		return err
	}

	if existingImg != nil {
		// Check for conflicts
		return existingImg.ConflictsWith(i)
	}

	im.imagesLock.Lock()
	defer im.imagesLock.Unlock()

	im.images = append(im.images, *i.Clone())
	return nil
}
