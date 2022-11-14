package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"

	"github.com/rs/cors"

	"github.com/CrowhopTech/shinysorter/backend/pkg/filedb"
	"github.com/CrowhopTech/shinysorter/backend/pkg/filedb/mongofile"
	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/server/restapi"
	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/server/restapi/operations"
	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/server/restapi/operations/files"
	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/server/restapi/operations/questions"
	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/server/restapi/operations/tags"
	"github.com/sirupsen/logrus"
)

var rootCtx context.Context

var (
	imageMetadataConnection filedb.FileMetadataService

	mongodbConectionURI   = flag.String("mongodb-connection-uri", "mongodb://localhost:27017", "The connection URI for the MongoDB metadata database")
	storageDirFlag        = flag.String("storage-dir", "./storage", "The directory to store files in")
	databaseDumpFrequency = flag.Duration("dump-frequency", time.Hour*24, "How often to dump a JSON copy of the database into the storage dir")
	// TODO: implement a database retention policy (short-term clear out after x days, but keep one from every y up to z)

	listenPort = flag.Int("listen-port", 10000, "The port for the server to listen on")
	logLevel   = flag.String("log-level", "info", "The log level to print at")
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

	parsedLevel, err := logrus.ParseLevel(*logLevel)
	if err != nil {
		logrus.Panicf("Failed to parse log level %s", *logLevel)
	}
	logrus.SetLevel(parsedLevel)
}

func CheckHealth(params operations.CheckHealthParams) middleware.Responder {
	// TODO: implement a "startup routine" for liveness vs. readiness
	// TODO: flip the order of initialization so that we set up the REST server first, and then add a proper
	//       health check here to validate the DB connection
	return operations.NewCheckHealthOK() // If we got here, the server is clearly up so let's just return OK
}

func main() {
	rootCtx = context.Background()

	parseFlags()

	logrus.Info("Initializing database connection...")

	// Initialize database connection
	mongoConn, cleanupFunc, err := mongofile.New(rootCtx, *mongodbConectionURI, false)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to initialize database connection")
	}
	if cleanupFunc != nil {
		defer cleanupFunc()
	}

	logrus.Info("Successfully connected to database")

	imageMetadataConnection = mongoConn

	logrus.Info("Initializing Swagger spec...")

	// Initialize Swagger
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		logrus.WithError(err).Fatal("Failed to load Swagger spec")
	}

	api := operations.NewShinySorterAPI(swaggerSpec)

	api.CheckHealthHandler = operations.CheckHealthHandlerFunc(CheckHealth)

	api.FilesListFilesHandler = files.ListFilesHandlerFunc(ListFiles)
	api.FilesGetFileByIDHandler = files.GetFileByIDHandlerFunc(GetFileByID)
	api.FilesCreateFileHandler = files.CreateFileHandlerFunc(CreateFile)
	api.FilesPatchFileByIDHandler = files.PatchFileByIDHandlerFunc(PatchFileByID)

	api.FilesGetFileContentHandler = files.GetFileContentHandlerFunc(GetFileContent)
	api.FilesSetFileContentHandler = files.SetFileContentHandlerFunc(SetFileContent)

	api.TagsListTagsHandler = tags.ListTagsHandlerFunc(ListTags)
	api.TagsCreateTagHandler = tags.CreateTagHandlerFunc(CreateTag)
	api.TagsPatchTagByIDHandler = tags.PatchTagByIDHandlerFunc(PatchTagByID)
	api.TagsDeleteTagHandler = tags.DeleteTagHandlerFunc(DeleteTag)

	api.QuestionsListQuestionsHandler = questions.ListQuestionsHandlerFunc(ListQuestions)
	api.QuestionsCreateQuestionHandler = questions.CreateQuestionHandlerFunc(CreateQuestion)
	api.QuestionsPatchQuestionByIDHandler = questions.PatchQuestionByIDHandlerFunc(PatchQuestionByID)
	api.QuestionsDeleteQuestionHandler = questions.DeleteQuestionHandlerFunc(DeleteQuestion)
	api.QuestionsReorderQuestionsHandler = questions.ReorderQuestionsHandlerFunc(ReorderQuestions)

	logrus.Info("Swagger spec and handlers initialized, starting to listen for requests")

	go databaseDumpLoop(*databaseDumpFrequency, path.Join(*storageDirFlag, "dumps"))

	// Start listening using having the handlers and port already set up.
	// Add the CORS AllowAll policy since the web UI is running on a different port
	// on the same address, so technically cross-origin.
	exposeCors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
		ExposedHeaders:   []string{"X-Filecount"},
	})

	if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", *listenPort), exposeCors.Handler(api.Serve(nil))); err != nil {
		log.Fatalln(err)
	}
}
