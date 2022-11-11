package main

import (
	"fmt"

	"github.com/CrowhopTech/shinysorter/backend/pkg/filedb"
	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/server/models"
	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/server/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

func translateDBTagToREST(tag *filedb.Tag) *models.TagEntry {
	if tag == nil {
		return nil
	}
	return &models.TagEntry{
		Description:      &tag.Description,
		ID:               &tag.ID,
		UserFriendlyName: &tag.UserFriendlyName,
	}
}

func tagOptionToFiledb(to *models.TagOption) filedb.TagOption {
	// TODO: possible NPE
	return filedb.TagOption{
		TagID:      *to.TagID,
		OptionText: *to.OptionText,
	}
}

func tagOptionArrayTofiledb(input []*models.TagOption) []filedb.TagOption {
	toReturn := make([]filedb.TagOption, len(input))
	for i, t := range input {
		toReturn[i] = tagOptionToFiledb(t)
	}
	return toReturn
}

func tagOptionToSwagger(to filedb.TagOption) *models.TagOption {
	return &models.TagOption{
		OptionText: &to.OptionText,
		TagID:      &to.TagID,
	}
}

func tagOptionArrayToSwagger(input []filedb.TagOption) []*models.TagOption {
	toReturn := make([]*models.TagOption, len(input))
	for i, t := range input {
		toReturn[i] = tagOptionToSwagger(t)
	}
	return toReturn
}

//ListTags lists all registered tags
func ListTags(params operations.ListTagsParams) middleware.Responder {
	requestCtx := rootCtx

	results, err := imageMetadataConnection.ListTags(requestCtx)
	if err != nil {
		return operations.NewListTagsInternalServerError().WithPayload(fmt.Sprintf("failed to list tags: %v", err))
	}

	output := []*models.TagEntry{}

	for _, tag := range results {
		converted := translateDBTagToREST(tag)
		output = append(output, converted)
	}

	return operations.NewListTagsOK().WithPayload(output)
}

func CreateTag(params operations.CreateTagParams) middleware.Responder {
	requestCtx := rootCtx

	createdTag, err := imageMetadataConnection.CreateTag(requestCtx, &filedb.Tag{
		UserFriendlyName: *params.NewTag.UserFriendlyName,
		Description:      *params.NewTag.Description,
	})
	if err != nil {
		return operations.NewCreateTagInternalServerError().WithPayload(fmt.Sprintf("failed to insert tag: %v", err))
	}

	output := translateDBTagToREST(createdTag)

	return operations.NewCreateTagCreated().WithPayload(output)
}

func PatchTagByID(params operations.PatchTagByIDParams) middleware.Responder {
	requestCtx := rootCtx

	tag := filedb.Tag{
		ID: params.ID,
	}

	if len(params.Patch.UserFriendlyName) > 0 {
		tag.UserFriendlyName = params.Patch.UserFriendlyName
	}

	if len(params.Patch.Description) > 0 {
		tag.Description = params.Patch.Description
	}

	newTag, err := imageMetadataConnection.ModifyTag(requestCtx, &tag)
	if err != nil {
		return operations.NewPatchTagByIDInternalServerError().WithPayload(fmt.Sprintf("failed to modify tag entry %d: %v", params.ID, err))
	}

	output := translateDBTagToREST(newTag)

	return operations.NewPatchTagByIDOK().WithPayload(output)
}

func DeleteTag(params operations.DeleteTagParams) middleware.Responder {
	requestCtx := rootCtx

	err := imageMetadataConnection.DeleteTag(requestCtx, params.ID)
	if err != nil {
		return operations.NewDeleteTagInternalServerError().WithPayload(fmt.Sprintf("failed to delete tag: %v", err))
	}

	return operations.NewDeleteTagOK()
}
