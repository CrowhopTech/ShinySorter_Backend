package main

import (
	"fmt"

	"github.com/CrowhopTech/shinysorter/backend/pkg/filedb"
	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/server/models"
	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/server/restapi/operations/files"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func translateDBFileToREST(img *filedb.File) *models.FileEntry {
	if img == nil {
		return nil
	}
	tags := []int64{}
	if img.Tags != nil {
		tags = *img.Tags
	}
	id := img.ID
	return &models.FileEntry{
		ID:            &id,
		Name:          &img.Name,
		Md5sum:        &img.Md5Sum,
		Tags:          tags,
		HasBeenTagged: *img.HasBeenTagged,
		MimeType:      &img.MIMEType,
	}
}

// ListFiles gets images matching the given query parameters
func ListFiles(params files.ListFilesParams) middleware.Responder {
	requestCtx := rootCtx

	requireTagged := true
	filter := filedb.FileFilter{}

	if params.HasBeenTagged != nil && !*params.HasBeenTagged {
		requireTagged = false
	}
	filter.Tagged = &requireTagged

	if len(params.IncludeTags) > 0 {
		filter.RequireTagOperation = filedb.All
		if params.IncludeOperator != nil {
			switch *params.IncludeOperator {
			case "any":
				filter.RequireTagOperation = filedb.Any
			case "all":
				filter.RequireTagOperation = filedb.All
			default:
				return files.NewListFilesBadRequest().WithPayload(fmt.Sprintf("failed to parse tag operator '%s'", *params.IncludeOperator))
			}
		}
		filter.RequireTags = params.IncludeTags
	}

	if len(params.ExcludeTags) > 0 {
		filter.ExcludeTagOperation = filedb.All
		if params.ExcludeOperator != nil {
			switch *params.ExcludeOperator {
			case "any":
				filter.ExcludeTagOperation = filedb.Any
			case "all":
				filter.ExcludeTagOperation = filedb.All
			default:
				return files.NewListFilesBadRequest().WithPayload(fmt.Sprintf("failed to parse tag operator '%s'", *params.ExcludeOperator))
			}
		}
		filter.ExcludeTags = params.ExcludeTags
	}

	if params.Limit != nil {
		filter.Limit = *params.Limit
	}

	if params.Continue != nil {
		parsedCont, err := primitive.ObjectIDFromHex(*params.Continue)
		if err != nil {
			return files.NewListFilesBadRequest().WithPayload(fmt.Sprintf("invalid object ID for continue '%s': %v", *params.Continue, err))
		}

		filter.Continue = parsedCont
	}

	count := int64(-1)
	var err error

	if params.Continue == nil || *params.Continue == "" {
		logrus.WithField("filter", filter).Debug("First page, running count query")
		count, err = imageMetadataConnection.CountFiles(requestCtx, filter)
		if err != nil {
			return files.NewListFilesInternalServerError().WithPayload(fmt.Sprintf("failed to count images: %v", err))
		}

	}

	logrus.WithField("filter", filter).Info("Running file query")

	results, err := imageMetadataConnection.ListFiles(requestCtx, &filter)
	if err != nil {
		return files.NewListFilesInternalServerError().WithPayload(fmt.Sprintf("failed to list images: %v", err))
	}

	output := []*models.FileEntry{}

	for _, img := range results {
		converted := translateDBFileToREST(img)
		output = append(output, converted)
	}

	return files.NewListFilesOK().WithPayload(output).WithXFileCount(count)
}

func GetFileByID(params files.GetFileByIDParams) middleware.Responder {
	requestCtx := rootCtx

	results, err := imageMetadataConnection.ListFiles(requestCtx, &filedb.FileFilter{
		ID: params.ID,
	})
	if err != nil {
		return files.NewGetFileByIDInternalServerError().WithPayload(fmt.Sprintf("failed to list images with name filter: %v", err))
	}

	if len(results) == 0 {
		return files.NewGetFileByIDNotFound()
	}

	if len(results) > 1 {
		return files.NewGetFileByIDInternalServerError().WithPayload(fmt.Sprintf("image list for ID %s returned %d results, expected exactly 1", params.ID, len(results)))
	}

	output := translateDBFileToREST(results[0])

	return files.NewGetFileByIDOK().WithPayload(output)
}

func CreateFile(params files.CreateFileParams) middleware.Responder {
	requestCtx := rootCtx

	f := false
	createdID, err := imageMetadataConnection.CreateFileEntry(requestCtx, &filedb.File{
		FileMetadata: filedb.FileMetadata{
			Name: params.ID,
		},
		// TODO: validate that tags actually exist
		HasContent:    &f,
		HasBeenTagged: &f,
	})
	if err != nil {
		return files.NewCreateFileInternalServerError().WithPayload(fmt.Sprintf("failed to insert image: %v", err))
	}

	createdFile, err := imageMetadataConnection.GetFileByID(requestCtx, createdID)
	if err != nil {
		return files.NewCreateFileInternalServerError().WithPayload(fmt.Sprintf("failed to get created image: %v", err))
	}

	output := translateDBFileToREST(createdFile)

	return files.NewCreateFileCreated().WithPayload(output)
}

func PatchFileByID(params files.PatchFileByIDParams) middleware.Responder {
	requestCtx := rootCtx

	logrus.WithFields(logrus.Fields{
		"image_id": params.ID,
		"patch":    params.Patch,
	}).Info("Patching image")

	img := filedb.File{
		FileMetadata: filedb.FileMetadata{
			ID: params.ID,
		},
	}

	// TODO: validate that tags actually exist
	if params.Patch.Tags != nil {
		img.Tags = &params.Patch.Tags
	}

	if params.Patch.HasBeenTagged != nil {
		img.HasBeenTagged = params.Patch.HasBeenTagged
	}

	newImg, err := imageMetadataConnection.ModifyFileEntry(requestCtx, &img)
	if err != nil {
		return files.NewPatchFileByIDInternalServerError().WithPayload(fmt.Sprintf("failed to modify image entry %s: %v", params.ID, err))
	}

	output := translateDBFileToREST(newImg)

	return files.NewPatchFileByIDOK().WithPayload(output)
}
