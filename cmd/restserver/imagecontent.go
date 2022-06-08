package main

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/h2non/filetype"
	"github.com/sirupsen/logrus"

	"github.com/CrowhopTech/shinysorter/backend/pkg/image"
	"github.com/CrowhopTech/shinysorter/backend/pkg/imagedb"
	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/server/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

const (
	thumbnailSubdirName = "thumbs"
	thumbnailExtension  = ".png"
)

func getThumbnailPath(imageID string) string {
	return path.Join(*storageDirFlag, thumbnailSubdirName, imageID+thumbnailExtension)
}

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
	// If the file is a thumbnail, we'll return the thumbnail instead
	var filePath string
	if params.Thumb {
		filePath = getThumbnailPath(params.ID)
	} else {
		filePath = path.Join(*storageDirFlag, params.ID)
	}

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

func SetImageContent(params operations.SetImageContentParams) middleware.Responder {
	requestCtx := rootCtx

	if params.FileContents == nil {
		return operations.NewSetImageContentBadRequest().WithPayload("no file contents provided")
	}

	// Verify that the image exists
	img, err := imageMetadataConnection.GetImage(requestCtx, params.ID)
	if err != nil {
		return operations.NewSetImageContentInternalServerError().WithPayload(fmt.Sprintf("failed to list images with name filter: %v", err))
	}
	if img == nil {
		return operations.NewSetImageContentNotFound()
	}

	// Now that we know the image exists in the DB, let's set the contents (also prevents against file traversal!)
	filePath := path.Join(*storageDirFlag, params.ID)
	thumbPath := getThumbnailPath(params.ID)

	// TODO: expose the file permissions as a parameter
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return operations.NewSetImageContentInternalServerError().WithPayload(fmt.Sprintf("failed to open file '%s': %v", filePath, err))
	}

	writtenBytes, err := io.Copy(file, params.FileContents)
	if err != nil {
		return operations.NewSetImageContentInternalServerError().WithPayload(fmt.Sprintf("failed to set image content for file '%s': %v", filePath, err))
	}

	// TODO: compare md5sum of original to written?
	// TODO: allow passing md5sum of image as parameter to ensure it gets written correctly?

	logrus.WithFields(logrus.Fields{
		"bytes_written": writtenBytes,
		"file_path":     filePath,
	}).Debug("Wrote file contents, calculating file md5sum and MIME type")

	md5Sum, err := image.GetFileMd5Sum(filePath)
	if err != nil {
		return operations.NewSetImageContentInternalServerError().WithPayload(fmt.Sprintf("failed to get md5sum for file '%s': %v", filePath, err))
	}

	fileType, err := filetype.MatchFile(filePath)
	if err != nil {
		return operations.NewSetImageContentInternalServerError().WithPayload(fmt.Sprintf("failed to determine MIME type for file '%s': %v", filePath, err))
	}

	logrus.WithFields(logrus.Fields{
		"mime_type": fileType,
		"md5sum":    md5Sum,
	}).Debug("Calculated md5sum and MIME type, writing thumbnail")

	// Create the thumbnails directory if it doesn't exist
	if err := os.MkdirAll(path.Join(*storageDirFlag, "thumbs"), 0755); err != nil {
		return operations.NewSetImageContentInternalServerError().WithPayload(fmt.Sprintf("failed to create thumbnail directory '%s': %v", thumbPath, err))
	}

	// This will end up resulting in ".mp4.png" or ".png.png" or ".jpg.png"
	// TODO: rewrite this to have just the png file extension
	// TODO: store this in the DB so if we change the name format later we can still find things?
	err = image.WriteFileThumbnail(filePath, fileType, thumbPath)
	if err != nil {
		return operations.NewSetImageContentInternalServerError().WithPayload(fmt.Sprintf("failed to write thumbnail for file '%s': %v", thumbPath, err))
	}

	// Patch to set file as having contents and set the md5sum
	t := true
	_, err = imageMetadataConnection.ModifyImageEntry(requestCtx, &imagedb.Image{
		FileMetadata: imagedb.FileMetadata{
			Name:     params.ID,
			Md5Sum:   md5Sum,
			MIMEType: fileType.MIME.Value,
		},
		HasContent: &t,
	})
	if err != nil {
		return operations.NewSetImageContentInternalServerError().WithPayload(fmt.Sprintf("failed to mark file '%s' as having content: %v", filePath, err))
	}

	return operations.NewSetImageContentOK()
}
