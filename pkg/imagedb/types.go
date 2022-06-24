package imagedb

import (
	"context"
	"fmt"
)

// FileMetadata contains information that comes from the file itself:
// the name of the file, its checksum, etc.
type FileMetadata struct {
	Name     string `bson:"_id"`
	Md5Sum   string `bson:"md5sum"`
	MIMEType string `bson:"mimeType"`
}

// Image represents all data about an image: the file metadata, as well
// as any user-added metadata (tags) or any other useful caching info
type Image struct {
	FileMetadata  `bson:",inline"`
	Tags          *[]int64 `bson:"tags,omitempty"`
	HasBeenTagged *bool    `bson:"hasBeenTagged,omitempty"`
	HasContent    *bool    `bson:"hasContent,omitempty"`
}

type Tag struct {
	ID               int64  `bson:"_id"`
	Name             string `bson:"name"`
	UserFriendlyName string `bson:"userFriendlyName"`
	Description      string `bson:"description"`
}

type TagOption struct {
	TagID      int64  `bson:"tagID"`
	OptionText string `bson:"optionText"`
}

type Question struct {
	ID                int64       `bson:"_id"`
	OrderingID        int64       `bson:"orderingID"`
	QuestionText      string      `bson:"questionText"`
	TagOptions        []TagOption `bson:"tagOptions"`
	MutuallyExclusive *bool       `bson:"mutuallyExclusive"`
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
		return fmt.Errorf("images have conflicting md5sums '%s' and '%s'", i.Md5Sum, other.Md5Sum)
	}

	return nil
}

// Clone returns a copy of the image, but entirely detached from the original
// object. Modifications to the copied object will not affect the original in any way.
func (i *Image) Clone() *Image {
	copiedTags := make([]int64, len(*i.Tags))
	copy(copiedTags, *i.Tags)
	return &Image{
		FileMetadata: FileMetadata{
			Name:     i.Name,
			Md5Sum:   i.Md5Sum,
			MIMEType: i.MIMEType,
		},
		Tags:          &copiedTags,
		HasBeenTagged: i.HasBeenTagged,
		HasContent:    i.HasContent,
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

	// ModifyImageEntry will modify the given image. ID/name is required,
	// If any of the others are set they will overwrite. This includes tags,
	// so make sure to provide the whole tag array, not just modifications.
	// TODO: change this API to have better tag modifications (add and remove tags)
	ModifyImageEntry(ctx context.Context, i *Image) (*Image, error)

	// ListTags will return the list of all tags. There are no filter options as this
	// list will never be extremely large.
	ListTags(ctx context.Context) ([]*Tag, error)

	// CreateTag will define a new tag for use on images. ID will be auto-set, if a value
	// is provided it will not be used.
	CreateTag(ctx context.Context, t *Tag) (*Tag, error)

	// ModifyTag will modify the given tag with the new given information.
	ModifyTag(ctx context.Context, t *Tag) (*Tag, error)

	// DeleteTag will delete the given tag, as well as remove it from all images that reference it.
	DeleteTag(ctx context.Context, id int64) error

	ListQuestions(ctx context.Context) ([]*Question, error)

	CreateQuestion(ctx context.Context, q *Question) (*Question, error)

	ModifyQuestion(ctx context.Context, q *Question) (*Question, error)

	DeleteQuestion(ctx context.Context, id int64) error
}
