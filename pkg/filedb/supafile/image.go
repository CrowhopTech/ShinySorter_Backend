package supafile

import (
	"context"
	"fmt"
	"math"
	"strconv"

	"github.com/CrowhopTech/shinysorter/backend/pkg/filedb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *supaConnection) convertResultToStruct(results []map[string]interface{}) (*filedb.File, error) {
	var (
		ok bool

		id            int64
		name          string
		md5sum        string
		mimeType      string
		hasBeenTagged bool
		hasContent    bool
	)

	if len(results) == 0 {
		return nil, fmt.Errorf("no results provided to convertResultToStruct")
	}

	fileResult := results[0]

	idFloat, ok := fileResult["id"].(float64)
	if !ok {
		return nil, fmt.Errorf("failed to convert ID value '%v' to float64", fileResult["id"])
	}
	id = int64(math.Floor(idFloat))

	name, ok = fileResult["filename"].(string)
	if !ok {
		return nil, fmt.Errorf("failed to convert filename value '%v' to string", fileResult["filename"])
	}

	md5sum, ok = fileResult["md5sum"].(string)
	if !ok {
		return nil, fmt.Errorf("failed to convert md5sum value '%v' to string", fileResult["md5sum"])
	}

	mimeType, ok = fileResult["mimeType"].(string)
	if !ok {
		return nil, fmt.Errorf("failed to convert mime type value '%v' to string", fileResult["mimeType"])
	}

	hasContent, ok = fileResult["hasContent"].(bool)
	if !ok {
		return nil, fmt.Errorf("failed to convert 'has content' value '%v' to boolean", fileResult["hasContent"])
	}

	hasBeenTagged, ok = fileResult["hasBeenTagged"].(bool)
	if !ok {
		return nil, fmt.Errorf("failed to convert 'has been tagged' value '%v' to boolean", fileResult["hasBeenTagged"])
	}

	tags := []int64{}

	// Loop through the rest of the results and set the tags
	for _, r := range results {
		tagIDFloat, ok := r["tagid"].(float64)
		if !ok {
			return nil, fmt.Errorf("failed to convert tagID value '%v' to int64", fileResult["tagid"])
		}
		tagID := int64(math.Floor(tagIDFloat))

		tags = append(tags, tagID)
	}

	return &filedb.File{
		FileMetadata: filedb.FileMetadata{
			ID:       strconv.FormatInt(id, 10),
			Name:     name,
			Md5Sum:   md5sum,
			MIMEType: mimeType,
		},
		HasBeenTagged: &hasBeenTagged,
		HasContent:    &hasContent,
		Tags:          &tags,
	}, nil
}

// GetFileByName will get the file with the given name.
// If not found, will return nil, not an error.
func (s *supaConnection) GetFileByName(ctx context.Context, name string) (*filedb.File, error) {
	// TODO: sanitize!
	var result []map[string]interface{} // Will be a list of duplicate file entries, one for each tag
	_, err := s.client.DB.From("fileswithtags").Select("*", "", false).Eq("filename", name).ExecuteTo(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to get file by name '%s': %v", name, err)
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("no file found with filename '%s'", name)
	}

	return s.convertResultToStruct(result)
}

// GetFileByID will get the file with the given ID.
// If not found, will return nil, not an error.
func (s *supaConnection) GetFileByID(ctx context.Context, id string) (*filedb.File, error) {
	_, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid file ID '%s': %v", id, err)
	}
	// Even though we don't use the return value,
	// this will let us catch improperly formatted IDs a little easier

	var result []map[string]interface{} // Will be a list of duplicate file entries, one for each tag
	_, err = s.client.DB.From("fileswithtags").Select("*", "", false).Eq("id", id).ExecuteTo(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to get file by id '%s': %v", id, err)
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("no file found with id '%s'", id)
	}

	return s.convertResultToStruct(result)
}

// CountFiles will return the count of how many file entries match
// the given query.
func (s *supaConnection) CountFiles(ctx context.Context, filter filedb.FileFilter) (int64, error) {
	req := s.client.DB.From("files").Select("*", "exact", false)

	getQueriesForFilter(req, &filter)

	return -1, fmt.Errorf("not implemented")
}

// ListFiles will list files that match the given filter, if provided.
// If no filter is provided, all results will be returned (oh no).
// If no files match the filter, err will be nil and an empty slice will be returned.
// TODO: Implement a pagination strategy
func (s *supaConnection) ListFiles(ctx context.Context, filter *filedb.FileFilter) ([]*filedb.File, error) {
	return nil, fmt.Errorf("not implemented")
}

// CreateFileEntry will create the given entry in the database.
// If one already exists with the given name, this will check for conflicts
// using ConflictsWith. If there is a conflict, an error will be returned.
// If not, no action will be taken.
func (s *supaConnection) CreateFileEntry(ctx context.Context, i *filedb.File) (primitive.ObjectID, error) {
	return primitive.NilObjectID, fmt.Errorf("not implemented")
}

// ModifyFileEntry will modify the given file. ID/name is required,
// If any of the others are set they will overwrite. This includes tags,
// so make sure to provide the whole tag array, not just modifications.
// TODO: change this API to have better tag modifications (add and remove tags)
func (s *supaConnection) ModifyFileEntry(ctx context.Context, i *filedb.File) (*filedb.File, error) {
	return nil, fmt.Errorf("not implemented")
}
