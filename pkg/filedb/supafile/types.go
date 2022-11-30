package supafile

import (
	"context"
	"os"

	supa "github.com/dominictwlee/supabase-go"
)

const (
	databaseName            = "shiny_sorter"
	filesCollectionName     = "files"
	tagsCollectionName      = "tags"
	questionsCollectionName = "questions"
)

// var _ filedb.FileMetadataService = new(supaConnection)

type supaConnection struct {
	client *supa.Client
}

func New(ctx context.Context, connectionURI string, purge bool) (*supaConnection, error) {
	supabaseUrl := "http://db.pmtnfgtrjfxzgtspqnmk.supabase.co"
	supabaseKey := os.Getenv("SUPABASE_KEY")
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)

	return &supaConnection{
		client: supabase,
	}, nil
}
