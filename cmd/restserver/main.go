package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/CrowhopTech/shinysorter/backend/pkg/imagedb"
	"github.com/CrowhopTech/shinysorter/backend/pkg/imagedb/inmem"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var (
	imageMetadataConnection imagedb.ImageMetadata
)

// GET images with query
// PATCH image (update image)
// DELETE image (query option to delete original file too?) (move to trash bin?)

func imagesEndpoint(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getImages(w, r)
		return
	case http.MethodPatch:
		patchImage(w, r)
		return
	case http.MethodDelete:
		deleteImage(w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func getImages(w http.ResponseWriter, r *http.Request) {
	images, err := imageMetadataConnection.ListImages(context.Background(), nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("failed to list images: %v", err)))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(images)
}

func patchImage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("PATCH image"))
}

func deleteImage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("DELETE images"))
}

// Existing code from above
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/image", imagesEndpoint)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func parseFlags() {
	// TODO: add logic to validate flag values here
	flag.Parse()
}

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	parseFlags()

	// Initialize database connection
	inmemDB := inmem.New(true)

	imageMetadataConnection = inmemDB

	handleRequests()
}
