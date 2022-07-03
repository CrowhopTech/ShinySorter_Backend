package main

import (
	"fmt"

	"github.com/CrowhopTech/shinysorter/backend/pkg/filedb"
	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/server/models"
	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/server/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

func translateDBFileToREST(img *filedb.File) *models.File {
	if img == nil {
		return nil
	}
	tags := []int64{}
	if img.Tags != nil {
		tags = *img.Tags
	}
	return &models.File{
		ID:            img.Name,
		Md5sum:        img.Md5Sum,
		Tags:          tags,
		HasBeenTagged: img.HasBeenTagged,
		MimeType:      img.MIMEType,
	}
}

//ListFiles gets images matching the given query parameters
func ListFiles(params operations.ListFilesParams) middleware.Responder {
	requestCtx := rootCtx

	filter := filedb.FileFilter{
		Tagged: params.HasBeenTagged,
	}

	if len(params.IncludeTags) > 0 {
		filter.RequireTagOperation = filedb.All
		if params.IncludeOperator != nil {
			switch *params.IncludeOperator {
			case "any":
				filter.RequireTagOperation = filedb.Any
			case "all":
				filter.RequireTagOperation = filedb.All
			default:
				return operations.NewListFilesBadRequest().WithPayload(fmt.Sprintf("failed to parse tag operator '%s'", *params.IncludeOperator))
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
				return operations.NewListFilesBadRequest().WithPayload(fmt.Sprintf("failed to parse tag operator '%s'", *params.ExcludeOperator))
			}
		}
		filter.ExcludeTags = params.ExcludeTags
	}

	results, err := imageMetadataConnection.ListFiles(requestCtx, &filter)
	if err != nil {
		return operations.NewListFilesInternalServerError().WithPayload(fmt.Sprintf("failed to list images: %v", err))
	}

	if len(results) == 0 {
		return operations.NewListFilesNotFound().WithPayload([]*models.File{})
	}

	output := []*models.File{}

	for _, img := range results {
		converted := translateDBFileToREST(img)
		output = append(output, converted)
	}

	return operations.NewListFilesOK().WithPayload(output)
}

func GetFileByID(params operations.GetFileByIDParams) middleware.Responder {
	requestCtx := rootCtx

	results, err := imageMetadataConnection.ListFiles(requestCtx, &filedb.FileFilter{
		Name: params.ID,
	})
	if err != nil {
		return operations.NewGetFileByIDInternalServerError().WithPayload(fmt.Sprintf("failed to list images with name filter: %v", err))
	}

	if len(results) == 0 {
		return operations.NewGetFileByIDNotFound()
	}

	if len(results) > 1 {
		return operations.NewGetFileByIDInternalServerError().WithPayload(fmt.Sprintf("image list for ID %s returned %d results, expected exactly 1", params.ID, len(results)))
	}

	output := translateDBFileToREST(results[0])

	return operations.NewGetFileByIDOK().WithPayload(output)
}

func CreateFile(params operations.CreateFileParams) middleware.Responder {
	requestCtx := rootCtx

	f := false
	err := imageMetadataConnection.CreateFileEntry(requestCtx, &filedb.File{
		FileMetadata: filedb.FileMetadata{
			Name:   params.NewFile.ID,
			Md5Sum: params.NewFile.Md5sum,
		},
		// TODO: validate that tags actually exist
		Tags:          &params.NewFile.Tags,
		HasContent:    &f,
		HasBeenTagged: &f,
	})
	if err != nil {
		return operations.NewCreateFileInternalServerError().WithPayload(fmt.Sprintf("failed to insert image: %v", err))
	}

	createdFile, err := imageMetadataConnection.GetFile(requestCtx, params.NewFile.ID)
	if err != nil {
		return operations.NewCreateFileInternalServerError().WithPayload(fmt.Sprintf("failed to get created image: %v", err))
	}

	output := translateDBFileToREST(createdFile)

	return operations.NewCreateFileCreated().WithPayload(output)
}

func PatchFileByID(params operations.PatchFileByIDParams) middleware.Responder {
	requestCtx := rootCtx

	logrus.WithFields(logrus.Fields{
		"image_id": params.ID,
		"patch":    params.Patch,
	}).Info("Patching image")

	img := filedb.File{
		FileMetadata: filedb.FileMetadata{
			Name: params.ID,
		},
	}

	if len(params.Patch.Md5sum) > 0 {
		img.FileMetadata.Md5Sum = params.Patch.Md5sum
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
		return operations.NewPatchFileByIDInternalServerError().WithPayload(fmt.Sprintf("failed to modify image entry %s: %v", params.ID, err))
	}

	output := translateDBFileToREST(newImg)

	return operations.NewPatchFileByIDOK().WithPayload(output)
}
