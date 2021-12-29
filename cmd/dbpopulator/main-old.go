// package main

// import (
// 	"context"
// 	"flag"
// 	"time"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"

// 	"github.com/CrowhopTech/shinysorter/backend/pkg/filescan"
// 	"github.com/sirupsen/logrus"
// )

// func ImageFromFileEntry(f *filescan.FileEntry) ImageOld {
// 	return ImageOld{
// 		Path:       f.Path,
// 		Inode:      f.Inode,
// 		AccessedAt: f.ADate,
// 	}
// }

// type ImageOld struct {
// 	Id         primitive.ObjectID `bson:"_id,omitempty"`
// 	Path       string             `bson:"path"`
// 	Inode      string             `bson:"inode"` // Basically the real "primary key"
// 	AccessedAt time.Time          `bson:"changedAt"`
// }

// type PopulatorConfig struct {
// 	ID                     primitive.ObjectID `bson:"_id,omitempty"`                    // should always be set to some sort of "one" value
// 	LastScannedFileModDate time.Time          `bson:"lastScannedFileModDate,omitempty"` // mdate of the last file we scanned, use this as the "since" for any search we do
// }

// var (
// 	populatorConfigID, _ = primitive.ObjectIDFromHex("000000000000000000000001")
// 	populatorConfig      PopulatorConfig

// 	defaultConfig = PopulatorConfig{
// 		ID:                     populatorConfigID,
// 		LastScannedFileModDate: time.Unix(0, 0),
// 	}
// )

// func mainold() {
// 	// TODO: NEW PLAN!
// 	// There's really no good way for us to do this in-place without assuming full control
// 	// over the intake process. Since I'd still like to touch these files in Explorer, I'd
// 	// instead like to make an "intake" folder where we can drop files and they'll be copied
// 	// over, inserted into the DB, and removed only once the DB insert is successful.

// 	logrus.Fatal("If you're seeing this error, check the comment I left - past Adam")

// 	logrus.SetLevel(logrus.DebugLevel)

// 	//fileDirectory := flag.String("file-directory", "C:\\Users\\creep\\Desktop\\testdata", "What directory the populator should (recursively) scan for files")
// 	fileRescanIntervalFlag := flag.Duration("file-rescan-interval", time.Second*30, "How often the db-populator should rescan files added after the last scan")
// 	// dbRescanIntervalFlag := flag.Duration("db-rescan-interval", time.Hour*24, "How often the db-populator should recheck all files in the db to scan for missing files")
// 	mongodbConectionURI := flag.String("mongodb-connection-uri", "mongodb://localhost:27017", "The connection URI for the MongoDB metadata database")

// 	flag.Parse()

// 	logrus.Infof("Starting image-tagging dbpopulator")

// 	logrus.Debugf("Connecting to DB")
// 	cleanup, err := initDB(context.Background(), *mongodbConectionURI)
// 	if cleanup != nil {
// 		defer cleanup()
// 	}
// 	if err != nil {
// 		logrus.WithError(err).Fatal("Failed to connect to DB")
// 		// Stops here, fatal
// 	}

// 	searchPath := "/mnt/c/Users/creep/Desktop/testdata"
// 	fileScanner, err := filescan.NewADateScanner(*fileRescanIntervalFlag, searchPath, nil, func(t time.Time) error {
// 		populatorConfig.LastScannedFileModDate = t
// 		return updateDBConfig(context.Background())
// 	})
// 	if err != nil {
// 		logrus.WithError(err).Fatal("Failed to initialize file scanner")
// 		// Stops here, fatal
// 	}

// 	fileScanner.RegisterCallback(processFile)
// 	_ = fileScanner.Watch(context.Background(), false, false) // Will keep scanning forever, no reason to check for errors
// }

// func processFile(ctx context.Context, f *filescan.FileEntry) error {
// 	// Get entry for the given inode
// 	// If exists
// 	//   Update it, check for conflicts if any
// 	// Else
// 	//   Insert new entry
// 	//   Check for other entries with same path, freak out if exists?
// 	result := imagesCollection.FindOne(ctx, bson.M{
// 		"inode": bson.M{
// 			"$in": bson.A{f.Inode},
// 		},
// 	})
// 	if result.Err() != nil && result.Err() != mongo.ErrNoDocuments {
// 		return result.Err()
// 	}

// 	if result.Err() != nil {
// 		// Not found error
// 		// This inode didn't exist before: *should* be a new file, but check for conflict paths elsewhere in the DB
// 		// TODO: just add a unique index instead, that will prevent dupes
// 		insertResult, err := imagesCollection.InsertOne(ctx, ImageFromFileEntry(f))
// 		if err != nil {
// 			return err
// 		}
// 		f.LogFields(false).WithFields(logrus.Fields{
// 			"new_id": insertResult.InsertedID,
// 		}).Debug("Inserted new image document")
// 	} else {
// 		// Exists, update
// 		existingImage := ImageOld{}
// 		err := result.Decode(&existingImage)
// 		if err != nil {
// 			return err
// 		}

// 		// TODO: just add a unique index instead, that will prevent dupes
// 		// TODO: Update file hash
// 		updateResult, err := imagesCollection.UpdateByID(ctx, existingImage.Id, bson.M{
// 			"$set": bson.M{
// 				"path":      f.Path,
// 				"changedAt": f.ADate,
// 			},
// 		})
// 		// UpdateByID doesn't throw an "error" for no matches, so let's make it one
// 		if err == nil && updateResult.MatchedCount == 0 {
// 			err = mongo.ErrNoDocuments
// 		}
// 		if err != nil {
// 			return err
// 		}

// 		f.LogFields(false).Debug("Updated image")
// 	}

// 	return nil
// }

// func initDBOld(ctx context.Context, uri string) (func() error, error) {
// 	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
// 	if err != nil {
// 		return nil, err
// 	}

// 	cleanupFunc := func() error {
// 		return client.Disconnect(context.TODO())
// 	}

// 	// OPTIONAL: drop everything!
// 	err = client.Database(databaseName).Drop(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	imagesCollection = client.Database(databaseName).Collection(imagesCollectionName)
// 	configCollection = client.Database(databaseName).Collection(configCollectionName)

// 	result := configCollection.FindOne(ctx, PopulatorConfig{
// 		ID: populatorConfigID,
// 	})
// 	found := true
// 	if result.Err() != nil {
// 		if result.Err() != mongo.ErrNoDocuments {
// 			return cleanupFunc, result.Err()
// 		}
// 		found = false
// 	}

// 	if err := result.Decode(&populatorConfig); err != nil {
// 		if err != mongo.ErrNoDocuments {
// 			return cleanupFunc, err
// 		}
// 		found = false
// 	}

// 	if !found {
// 		// Default
// 		populatorConfig = defaultConfig
// 	}

// 	err = updateDBConfig(ctx)
// 	if err != nil {
// 		return cleanupFunc, err
// 	}

// 	return cleanupFunc, nil
// }

// func updateDBConfig(ctx context.Context) error {
// 	upsert := true
// 	_, err := configCollection.ReplaceOne(ctx,
// 		bson.M{
// 			"_id": populatorConfigID,
// 		},
// 		populatorConfig,
// 		&options.ReplaceOptions{Upsert: &upsert})
// 	return err
// }

// // fileRescanIntervalFlag will run fileScan() every scanInterval.
// // func backgroundScan(scanInterval time.Duration) {

// // }

// // fileScan() will recheck for any new files: any files that have a cdate or mdate later than the last scan.
// // func fileScan(ctx context.Context) error {

// // 	searchTime := populatorConfig.LastScannedFileModDate.Add(time.Second)

// // 	// TODO: increment time AFTER processing files, not before
// // 	//       This will help avoid issues if we panic/error on one file in a way we don't recover from
// // 	paths, err := findNewFiles(searchPath, &searchTime)
// // 	if err != nil {
// // 		return err
// // 	}

// // 	for _, file := range paths {

// // 		}
// // 	}

// // 	return nil
// // }
package main
