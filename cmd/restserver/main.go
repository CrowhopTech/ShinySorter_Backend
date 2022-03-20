package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"

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
	tags := map[string]string{}
	for k, v := range img.Tags {
		tags[k] = v
	}

	return &models.Image{
		ID:     &img.Name,
		Md5sum: &img.Md5Sum,
		Tags:   tags,
	}
}

// tagexists,tagemptyval:,tag:val,tag:val1|val2,tag:!notval1|notval2
// Split by comma
//   Split by colon
//   If no colon, just an exists query
//   If colon, check for exclamation (save as inverse)
//   Split second part by comma

func parseTagQueryString(tags []string) (map[string]*imagedb.TagSearch, error) {
	const (
		kvDelimeter     = ":"
		inverseMarker   = "!"
		valuesDelimeter = "|"
	)

	toReturn := map[string]*imagedb.TagSearch{}

	for _, t := range tags {
		tag := strings.TrimSpace(t)
		if len(tag) == 0 {
			continue
		}
		kvComps := strings.Split(tag, kvDelimeter)
		if len(kvComps) == 1 {
			// Just an exists query
			toReturn[kvComps[0]] = nil
			continue
		}
		if len(kvComps) > 2 {
			return nil, fmt.Errorf("tag '%s' is malformatted", tag)
		}
		inverseQuery := strings.HasPrefix(kvComps[1], inverseMarker)
		valuesString := strings.TrimPrefix(kvComps[1], inverseMarker)
		values := strings.Split(valuesString, valuesDelimeter)
		toReturn[kvComps[0]] = &imagedb.TagSearch{
			TagValues: values,
			Invert:    inverseQuery,
		}
	}

	return toReturn, nil
}

//GetImages gets images matching the given query parameters
func GetImages(params operations.GetImagesParams) middleware.Responder {
	requestCtx := rootCtx

	filter := imagedb.ImageFilter{}

	if len(params.IncludeTags) > 0 {
		filter.RequireTagOperation = imagedb.All
		if params.IncludeOperator != nil {
			switch *params.IncludeOperator {
			case "any":
				filter.RequireTagOperation = imagedb.Any
			case "all":
				filter.RequireTagOperation = imagedb.All
			default:
				return operations.NewGetImagesBadRequest().WithPayload(fmt.Sprintf("failed to parse tag operator '%s'", *params.IncludeOperator))
			}
		}

		requireTags, err := parseTagQueryString(params.IncludeTags)
		if err != nil {
			return operations.NewGetImagesBadRequest().WithPayload(fmt.Sprintf("failed to parse tag query: %v", err))
		}
		filter.RequireTags = requireTags
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
				return operations.NewGetImagesBadRequest().WithPayload(fmt.Sprintf("failed to parse tag operator '%s'", *params.ExcludeOperator))
			}
		}

		excludeTags, err := parseTagQueryString(params.ExcludeTags)
		if err != nil {
			return operations.NewGetImagesBadRequest().WithPayload(fmt.Sprintf("failed to parse tag query: %v", err))
		}
		filter.ExcludeTags = excludeTags
	}

	results, err := imageMetadataConnection.ListImages(requestCtx, &filter)
	if err != nil {
		return operations.NewGetImagesInternalServerError().WithPayload(fmt.Sprintf("failed to list images: %v", err))
	}

	if len(results) == 0 {
		return operations.NewGetImagesNotFound().WithPayload("[]")
	}

	output := []*models.Image{}

	for _, img := range results {
		converted := translateDBImageToREST(img)
		output = append(output, converted)
	}

	return operations.NewGetImagesOK().WithPayload(output)
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

func PatchImageByID(params operations.PatchImageByIDParams) middleware.Responder {
	return operations.NewPatchImageByIDBadRequest()
}

func GetImageContent(params operations.GetImageContentParams) middleware.Responder {
	return operations.NewGetImageContentNotFound()
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
	server := restapi.NewServer(api)

	api.GetImagesHandler = operations.GetImagesHandlerFunc(GetImages)
	api.GetImageByIDHandler = operations.GetImageByIDHandlerFunc(GetImageByID)
	api.PatchImageByIDHandler = operations.PatchImageByIDHandlerFunc(PatchImageByID)
	api.GetImageContentHandler = operations.GetImageContentHandlerFunc(GetImageContent)
	api.CheckHealthHandler = operations.CheckHealthHandlerFunc(CheckHealth)
	defer server.Shutdown()

	server.Port = 10000

	// Start listening using having the handlers and port
	// already set up.
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
