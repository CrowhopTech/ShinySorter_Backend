package imagedb

import (
	"context"
	"fmt"
)

// FileMetadata contains information that comes from the file itself:
// the name of the file, its checksum, etc.
type FileMetadata struct {
	Name   string `bson:"_id"`
	Md5Sum string `bson:"md5sum"`
}

// Image represents all data about an image: the file metadata, as well
// as any user-added metadata (tags) or any other useful caching info
type Image struct {
	FileMetadata `bson:",inline"`
	Tags         map[string]string `bson:"tags"`
}

// ConflictsWith returns if the provided image has unresolvable conflicts
// with this image. This includes:
//  * Mismatched Md5sums
//  * Mismatched names
// But does not include:
//  * Tags in any way
func (i *Image) ConflictsWith(other *Image) error {
	if other == nil {
		return nil // No conflict, other doesn't exist
	}

	if i.Name != other.Name {
		return fmt.Errorf("images have conflicting names '%s' and '%s", i.Name, other.Name)
	}

	if i.Md5Sum != other.Md5Sum {
		return fmt.Errorf("images have conflicting md5sums")
	}

	return nil
}

func (i *Image) Clone() *Image {
	copiedTags := map[string]string{}
	for k, v := range i.Tags {
		copiedTags[k] = v
	}
	return &Image{
		FileMetadata: FileMetadata{
			Name:   i.Name,
			Md5Sum: i.Md5Sum,
		},
		Tags: copiedTags,
	}
}

// ImageMetadata represents a service to access image metadata from a
// given backing database.
type ImageMetadata interface {
	// GetImage will get the image with the given name.
	// If not found, will return nil, not an error.
	GetImage(ctx context.Context, name string) (*Image, error)

	// ListImages will list images that match the given filter, if provided.
	// If no filter is provided, all results will be returned (oh no).
	// If no images match the filter, err will be nil and an empty slice will be returned.
	// TODO: Implement a pagination strategy
	ListImages(ctx context.Context, filter *ImageFilter) ([]*Image, error)

	// CreateImageEntry will create the given entry in the database.
	// If one already exists with the given name, this will check for conflicts
	// using ConflictsWith. If there is a conflict, an error will be returned.
	// If not, no action will be taken.
	CreateImageEntry(ctx context.Context, i *Image) error
}
