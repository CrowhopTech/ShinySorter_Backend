package main

import (
	"fmt"

	"github.com/CrowhopTech/shinysorter/backend/pkg/imagedb"
	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/server/models"
	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/server/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

func translateDBImageToREST(img *imagedb.Image) *models.Image {
	if img == nil {
		return nil
	}
	tags := []int64{}
	if img.Tags != nil {
		tags = *img.Tags
	}
	return &models.Image{
		ID:            img.Name,
		Md5sum:        img.Md5Sum,
		Tags:          tags,
		HasBeenTagged: img.HasBeenTagged,
		MimeType:      img.MIMEType,
	}
}

//ListImages gets images matching the given query parameters
func ListImages(params operations.ListImagesParams) middleware.Responder {
	requestCtx := rootCtx

	filter := imagedb.ImageFilter{
		Tagged: params.HasBeenTagged,
	}

	if len(params.IncludeTags) > 0 {
		filter.RequireTagOperation = imagedb.All
		if params.IncludeOperator != nil {
			switch *params.IncludeOperator {
			case "any":
				filter.RequireTagOperation = imagedb.Any
			case "all":
				filter.RequireTagOperation = imagedb.All
			default:
				return operations.NewListImagesBadRequest().WithPayload(fmt.Sprintf("failed to parse tag operator '%s'", *params.IncludeOperator))
			}
		}
		filter.RequireTags = params.IncludeTags
	}

	if len(params.ExcludeTags) > 0 {
		filter.ExcludeTagOperation = imagedb.All
		if params.ExcludeOperator != nil {
			switch *params.ExcludeOperator {
			case "any":
				filter.ExcludeTagOperation = imagedb.Any
			case "all":
				filter.ExcludeTagOperation = imagedb.All
			default:
				return operations.NewListImagesBadRequest().WithPayload(fmt.Sprintf("failed to parse tag operator '%s'", *params.ExcludeOperator))
			}
		}
		filter.ExcludeTags = params.ExcludeTags
	}

	results, err := imageMetadataConnection.ListImages(requestCtx, &filter)
	if err != nil {
		return operations.NewListImagesInternalServerError().WithPayload(fmt.Sprintf("failed to list images: %v", err))
	}

	if len(results) == 0 {
		return operations.NewListImagesNotFound().WithPayload([]*models.Image{})
	}

	output := []*models.Image{}

	for _, img := range results {
		converted := translateDBImageToREST(img)
		output = append(output, converted)
	}

	return operations.NewListImagesOK().WithPayload(output)
}

func GetImageByID(params operations.GetImageByIDParams) middleware.Responder {
	requestCtx := rootCtx

	results, err := imageMetadataConnection.ListImages(requestCtx, &imagedb.ImageFilter{
		Name: params.ID,
	})
	if err != nil {
		return operations.NewGetImageByIDInternalServerError().WithPayload(fmt.Sprintf("failed to list images with name filter: %v", err))
	}

	if len(results) == 0 {
		return operations.NewGetImageByIDNotFound()
	}

	if len(results) > 1 {
		return operations.NewGetImageByIDInternalServerError().WithPayload(fmt.Sprintf("image list for ID %s returned %d results, expected exactly 1", params.ID, len(results)))
	}

	output := translateDBImageToREST(results[0])

	return operations.NewGetImageByIDOK().WithPayload(output)
}

func CreateImage(params operations.CreateImageParams) middleware.Responder {
	requestCtx := rootCtx

	f := false
	err := imageMetadataConnection.CreateImageEntry(requestCtx, &imagedb.Image{
		FileMetadata: imagedb.FileMetadata{
			Name:   params.NewImage.ID,
			Md5Sum: params.NewImage.Md5sum,
		},
		// TODO: validate that tags actually exist
		Tags:          &params.NewImage.Tags,
		HasContent:    &f,
		HasBeenTagged: &f,
	})
	if err != nil {
		return operations.NewCreateImageInternalServerError().WithPayload(fmt.Sprintf("failed to insert image: %v", err))
	}

	createdImage, err := imageMetadataConnection.GetImage(requestCtx, params.NewImage.ID)
	if err != nil {
		return operations.NewCreateImageInternalServerError().WithPayload(fmt.Sprintf("failed to get created image: %v", err))
	}

	output := translateDBImageToREST(createdImage)

	return operations.NewCreateImageCreated().WithPayload(output)
}

func PatchImageByID(params operations.PatchImageByIDParams) middleware.Responder {
	requestCtx := rootCtx

	logrus.WithFields(logrus.Fields{
		"image_id": params.ID,
		"patch":    params.Patch,
	}).Info("Patching image")

	img := imagedb.Image{
		FileMetadata: imagedb.FileMetadata{
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

	newImg, err := imageMetadataConnection.ModifyImageEntry(requestCtx, &img)
	if err != nil {
		return operations.NewPatchImageByIDInternalServerError().WithPayload(fmt.Sprintf("failed to modify image entry %s: %v", params.ID, err))
	}

	output := translateDBImageToREST(newImg)

	return operations.NewPatchImageByIDOK().WithPayload(output)
}
