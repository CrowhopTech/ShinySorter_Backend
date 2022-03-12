package main

import (
	"flag"
	"log"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"

	"github.com/CrowhopTech/shinysorter/backend/pkg/imagedb"
	"github.com/CrowhopTech/shinysorter/backend/pkg/imagedb/inmem"
	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/server/restapi"
	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/server/restapi/operations"
	"github.com/sirupsen/logrus"
)

var (
	imageMetadataConnection imagedb.ImageMetadata
)

func parseFlags() {
	// TODO: add logic to validate flag values here
	flag.Parse()
}

func CheckHealth(params operations.CheckHealthParams) middleware.Responder {
	return operations.NewCheckHealthServiceUnavailable()
}

//GetImages gets images matching the given query parameters
func GetImages(params operations.GetImagesParams) middleware.Responder {
	return operations.NewGetImagesBadRequest()
}

func GetImageByID(params operations.GetImageByIDParams) middleware.Responder {
	return operations.NewGetImageByIDNotFound()
}

func PatchImageByID(params operations.PatchImageByIDParams) middleware.Responder {
	return operations.NewPatchImageByIDBadRequest()
}

func GetImageContent(params operations.GetImageContentParams) middleware.Responder {
	return operations.NewGetImageContentNotFound()
}

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	parseFlags()

	// Initialize database connection
	inmemDB := inmem.New(true)

	imageMetadataConnection = inmemDB

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
