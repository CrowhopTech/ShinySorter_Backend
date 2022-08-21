package main

import (
	"context"
	"encoding/json"
	"os"
	"path"
	"time"

	"github.com/CrowhopTech/shinysorter/backend/pkg/filedb"
	"github.com/sirupsen/logrus"
)

type DatabaseDump struct {
	Files     []*filedb.File
	Tags      []*filedb.Tag
	Questions []*filedb.Question
}

func databaseDumpLoop(interval time.Duration, outputDir string) {
	t := time.NewTicker(interval)

	for {
		<-t.C

		databaseDump(outputDir)
	}
}

func databaseDump(outputDir string) {
	// TODO: expose permissions as a parameter
	err := os.MkdirAll(outputDir, os.FileMode(0755))
	if err != nil {
		logrus.WithError(err).Error("Failed to create dump directory: skipping this round")
	}

	// TODO: support pagination in this dump
	files, err := imageMetadataConnection.ListFiles(context.Background(), nil)
	if err != nil {
		logrus.WithError(err).Error("Failed to list files from database for regular dump: skipping this round")
		return
	}
	tags, err := imageMetadataConnection.ListTags(context.Background())
	if err != nil {
		logrus.WithError(err).Error("Failed to list tags from database for regular dump: skipping this round")
		return
	}
	questions, err := imageMetadataConnection.ListQuestions(context.Background())
	if err != nil {
		logrus.WithError(err).Error("Failed to list questions from database for regular dump: skipping this round")
		return
	}

	dump := DatabaseDump{
		Files:     files,
		Tags:      tags,
		Questions: questions,
	}
	contents, err := json.MarshalIndent(dump, "", "  ")
	if err != nil {
		logrus.WithError(err).Error("Failed to marshal dump into JSON: skipping this round")
		return
	}

	dumpPath := path.Join(outputDir, time.Now().Format(time.RFC822))
	err = os.WriteFile(dumpPath, contents, os.FileMode(0444)) // Write file as read-only so it's immutable
	if err != nil {
		logrus.WithError(err).Errorf("Failed to write dump to file '%s': skipping this round", dumpPath)
		return
	}

	logrus.WithField("path", dumpPath).Info("Successfully wrote database dump to storage directory")
}
