package main

import (
	"fmt"
	"os"
	"path"

	"github.com/h2non/filetype"

	"github.com/CrowhopTech/shinysorter/backend/pkg/imagedb"
	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/server/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

func GetImageContent(params operations.GetImageContentParams) middleware.Responder {
	requestCtx := rootCtx

	// Verify that the image exists
	results, err := imageMetadataConnection.ListImages(requestCtx, &imagedb.ImageFilter{
		Name: params.ID,
	})
	if err != nil {
		return operations.NewGetImageContentInternalServerError().WithPayload(fmt.Sprintf("failed to list images with name filter: %v", err))
	}

	if len(results) == 0 {
		return operations.NewGetImageContentNotFound()
	}

	if len(results) > 1 {
		return operations.NewGetImageContentInternalServerError().WithPayload(fmt.Sprintf("image list for ID %s returned %d results, expected exactly 1", params.ID, len(results)))
	}

	// Now that we know the image exists in the DB, let's read the contents (also prevents against file traversal!)
	filePath := path.Join(*storageDirFlag, params.ID)

	// TODO: do this whenever we set the content and save it in the DB!
	fileType, err := filetype.MatchFile(filePath)
	if err != nil {
		return operations.NewGetImageContentInternalServerError().WithPayload(fmt.Sprintf("failed to identify filetype for path '%s': %v", filePath, err))
	}

	fileMimeType := "application/octet-stream"
	if fileType != filetype.Unknown && fileType.MIME.Value != "" {
		fileMimeType = fileType.MIME.Value
	}

	file, err := os.Open(filePath)
	if err != nil {
		return operations.NewGetImageContentInternalServerError().WithPayload(fmt.Sprintf("failed to open file '%s': %v", filePath, err))
	}

	return operations.NewGetImageContentOK().WithContentType(fileMimeType).WithPayload(file)
}
