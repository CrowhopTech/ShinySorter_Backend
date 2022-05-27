package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"

	"github.com/rs/cors"

	"github.com/CrowhopTech/shinysorter/backend/pkg/imagedb"
	"github.com/CrowhopTech/shinysorter/backend/pkg/imagedb/mongoimg"
	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/server/models"
	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/server/restapi"
	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/server/restapi/operations"
	"github.com/sirupsen/logrus"
)

var rootCtx context.Context

var (
	imageMetadataConnection imagedb.ImageMetadata

	mongodbConectionURI = flag.String("mongodb-connection-uri", "mongodb://localhost:27017", "The connection URI for the MongoDB metadata database")
)

func parseFlags() {
	flag.Parse()
}

func CheckHealth(params operations.CheckHealthParams) middleware.Responder {
	// TODO: implement a "startup routine" for liveness vs. readiness
	return operations.NewCheckHealthServiceUnavailable()
}

func translateDBImageToREST(img *imagedb.Image) *models.Image {
	if img == nil {
		return nil
	}
	return &models.Image{
		ID:            img.Name,
		Md5sum:        img.Md5Sum,
		Tags:          *img.Tags,
		HasBeenTagged: img.HasBeenTagged,
	}
}

func translateDBTagToREST(tag *imagedb.Tag) *models.Tag {
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

func tagOptionToImagedb(to *models.QuestionTagOptionsItems0) imagedb.TagOption {
	// TODO: possible NPE
	return imagedb.TagOption{
		TagID:      *to.TagID,
		OptionText: *to.OptionText,
	}
}

func tagOptionArrayToImagedb(input []*models.QuestionTagOptionsItems0) []imagedb.TagOption {
	toReturn := make([]imagedb.TagOption, len(input))
	for i, t := range input {
		toReturn[i] = tagOptionToImagedb(t)
	}
	return toReturn
}

func tagOptionToSwagger(to imagedb.TagOption) *models.QuestionTagOptionsItems0 {
	return &models.QuestionTagOptionsItems0{
		OptionText: &to.OptionText,
		TagID:      &to.TagID,
	}
}

func tagOptionArrayToSwagger(input []imagedb.TagOption) []*models.QuestionTagOptionsItems0 {
	toReturn := make([]*models.QuestionTagOptionsItems0, len(input))
	for i, t := range input {
		toReturn[i] = tagOptionToSwagger(t)
	}
	return toReturn
}

func translateDBQuestionToREST(question *imagedb.Question) *models.Question {
	return &models.Question{
		OrderingID:       question.OrderingID,
		QuestionID:       question.ID,
		QuestionText:     question.QuestionText,
		RequiresQuestion: *question.RequiresQuestion,
		TagOptions:       tagOptionArrayToSwagger(question.TagOptions),
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
		return operations.NewListImagesNotFound().WithPayload("[]")
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

	err := imageMetadataConnection.CreateImageEntry(requestCtx, &imagedb.Image{
		FileMetadata: imagedb.FileMetadata{
			Name:   params.NewImage.ID,
			Md5Sum: params.NewImage.Md5sum,
		},
		// TODO: validate that tags actually exist
		Tags: &params.NewImage.Tags,
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

func GetImageContent(params operations.GetImageContentParams) middleware.Responder {
	return operations.NewGetImageContentInternalServerError().WithPayload("not implemented")
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

	createdTag, err := imageMetadataConnection.CreateTag(requestCtx, &imagedb.Tag{
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

	tag := imagedb.Tag{
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

//ListQuestions lists all registered questions
func ListQuestions(params operations.ListQuestionsParams) middleware.Responder {
	requestCtx := rootCtx

	results, err := imageMetadataConnection.ListQuestions(requestCtx)
	if err != nil {
		return operations.NewListQuestionsInternalServerError().WithPayload(fmt.Sprintf("failed to list questions: %v", err))
	}

	output := []*models.Question{}

	for _, question := range results {
		converted := translateDBQuestionToREST(question)
		output = append(output, converted)
	}

	return operations.NewListQuestionsOK().WithPayload(output)
}

func CreateQuestion(params operations.CreateQuestionParams) middleware.Responder {
	requestCtx := rootCtx

	createdQuestion, err := imageMetadataConnection.CreateQuestion(requestCtx, &imagedb.Question{
		ID:               params.NewQuestion.QuestionID,
		OrderingID:       params.NewQuestion.OrderingID,
		RequiresQuestion: &params.NewQuestion.RequiresQuestion,
		QuestionText:     params.NewQuestion.QuestionText,
		TagOptions:       tagOptionArrayToImagedb(params.NewQuestion.TagOptions),
	})
	if err != nil {
		return operations.NewCreateQuestionInternalServerError().WithPayload(fmt.Sprintf("failed to insert question: %v", err))
	}

	output := translateDBQuestionToREST(createdQuestion)

	return operations.NewCreateQuestionCreated().WithPayload(output)
}

func PatchQuestionByID(params operations.PatchQuestionByIDParams) middleware.Responder {
	requestCtx := rootCtx

	question := imagedb.Question{
		ID: params.ID,
	}

	if len(params.Patch.QuestionText) > 0 {
		question.QuestionText = params.Patch.QuestionText
	}

	// TODO: handle these three parameters better
	if len(params.Patch.TagOptions) > 0 {
		question.TagOptions = tagOptionArrayToImagedb(params.Patch.TagOptions)
	}

	if params.Patch.RequiresQuestion > 0 {
		question.RequiresQuestion = &params.Patch.RequiresQuestion
	}

	if params.Patch.OrderingID > 0 {
		question.OrderingID = params.Patch.OrderingID
	}

	newQuestion, err := imageMetadataConnection.ModifyQuestion(requestCtx, &question)
	if err != nil {
		return operations.NewPatchQuestionByIDInternalServerError().WithPayload(fmt.Sprintf("failed to modify question entry %d: %v", params.ID, err))
	}

	output := translateDBQuestionToREST(newQuestion)

	return operations.NewPatchQuestionByIDOK().WithPayload(output)
}

func main() {
	rootCtx = context.Background()

	logrus.SetLevel(logrus.DebugLevel)

	parseFlags()

	// Initialize database connection
	mongoConn, cleanupFunc, err := mongoimg.New(rootCtx, *mongodbConectionURI, false)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to initialize database connection")
	}
	defer cleanupFunc()

	imageMetadataConnection = mongoConn

	// Initialize Swagger
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewShinySorterAPI(swaggerSpec)

	api.CheckHealthHandler = operations.CheckHealthHandlerFunc(CheckHealth)

	api.ListImagesHandler = operations.ListImagesHandlerFunc(ListImages)
	api.GetImageByIDHandler = operations.GetImageByIDHandlerFunc(GetImageByID)
	api.CreateImageHandler = operations.CreateImageHandlerFunc(CreateImage)
	api.PatchImageByIDHandler = operations.PatchImageByIDHandlerFunc(PatchImageByID)

	api.GetImageContentHandler = operations.GetImageContentHandlerFunc(GetImageContent)

	api.ListTagsHandler = operations.ListTagsHandlerFunc(ListTags)
	api.CreateTagHandler = operations.CreateTagHandlerFunc(CreateTag)
	api.PatchTagByIDHandler = operations.PatchTagByIDHandlerFunc(PatchTagByID)

	api.ListQuestionsHandler = operations.ListQuestionsHandlerFunc(ListQuestions)
	api.CreateQuestionHandler = operations.CreateQuestionHandlerFunc(CreateQuestion)
	api.PatchQuestionByIDHandler = operations.PatchQuestionByIDHandlerFunc(PatchQuestionByID)

	// Start listening using having the handlers and port
	// already set up.
	// Add the CORS AllowAll policy since the web UI is running on a different port
	// on the same address, so technically cross-origin.
	if err := http.ListenAndServe(":10000", cors.AllowAll().Handler(api.Serve(nil))); err != nil {
		log.Fatalln(err)
	}
}
