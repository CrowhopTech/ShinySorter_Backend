package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"

	"github.com/rs/cors"

	"github.com/CrowhopTech/shinysorter/backend/pkg/imagedb"
	"github.com/CrowhopTech/shinysorter/backend/pkg/imagedb/mongoimg"
	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/server/restapi"
	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/server/restapi/operations"
	"github.com/sirupsen/logrus"
)

var rootCtx context.Context

var (
	imageMetadataConnection imagedb.ImageMetadata

	mongodbConectionURI = flag.String("mongodb-connection-uri", "mongodb://localhost:27017", "The connection URI for the MongoDB metadata database")
	storageDirFlag      = flag.String("storage-dir", "./storage", "The directory to store files in")
)

func parseFlags() {
	flag.Parse()

	if result, err := os.Stat(*storageDirFlag); err != nil {
		if os.IsNotExist(err) {
			logrus.Fatalf("Storage directory '%s' does not exist: please create it and try again", *storageDirFlag)
		} else {
			logrus.Fatalf("Error while checking info for storage directory '%s'", *storageDirFlag)
		}
	} else if !result.IsDir() {
		logrus.Fatalf("Storage path '%s' exists but is not a directory", *storageDirFlag)
	}
}

func CheckHealth(params operations.CheckHealthParams) middleware.Responder {
	// TODO: implement a "startup routine" for liveness vs. readiness
	return operations.NewCheckHealthServiceUnavailable()
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
