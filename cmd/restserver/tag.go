package main

import (
	"fmt"

	"github.com/CrowhopTech/shinysorter/backend/pkg/filedb"
	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/server/models"
	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/server/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

func translateDBTagToREST(tag *filedb.Tag) *models.Tag {
	if tag == nil {
		return nil
	}
	return &models.Tag{
		Description:      tag.Description,
		ID:               tag.ID,
		Name:             tag.Name,
		UserFriendlyName: tag.UserFriendlyName,
	}
}

func tagOptionToFiledb(to *models.QuestionTagOptionsItems0) filedb.TagOption {
	// TODO: possible NPE
	return filedb.TagOption{
		TagID:      *to.TagID,
		OptionText: *to.OptionText,
	}
}

func tagOptionArrayTofiledb(input []*models.QuestionTagOptionsItems0) []filedb.TagOption {
	toReturn := make([]filedb.TagOption, len(input))
	for i, t := range input {
		toReturn[i] = tagOptionToFiledb(t)
	}
	return toReturn
}

func tagOptionToSwagger(to filedb.TagOption) *models.QuestionTagOptionsItems0 {
	return &models.QuestionTagOptionsItems0{
		OptionText: &to.OptionText,
		TagID:      &to.TagID,
	}
}

func tagOptionArrayToSwagger(input []filedb.TagOption) []*models.QuestionTagOptionsItems0 {
	toReturn := make([]*models.QuestionTagOptionsItems0, len(input))
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

	output := []*models.Tag{}

	for _, tag := range results {
		converted := translateDBTagToREST(tag)
		output = append(output, converted)
	}

	return operations.NewListTagsOK().WithPayload(output)
}

func CreateTag(params operations.CreateTagParams) middleware.Responder {
	requestCtx := rootCtx

	createdTag, err := imageMetadataConnection.CreateTag(requestCtx, &filedb.Tag{
		Name:             params.NewTag.Name,
		UserFriendlyName: params.NewTag.UserFriendlyName,
		Description:      params.NewTag.Description,
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

	if len(params.Patch.Name) > 0 {
		tag.Name = params.Patch.Name
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
