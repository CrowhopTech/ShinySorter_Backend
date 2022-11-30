package main

import (
	"context"
	"fmt"

	"github.com/CrowhopTech/shinysorter/backend/pkg/filedb/supafile"
)

func main() {
	sup, err := supafile.New(context.Background(), "", false)
	if err != nil {
		panic(err)
	}

	res, err := sup.GetFileByName(context.Background(), "test-file.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(res)

	res, err = sup.GetFileByID(context.Background(), "1")
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
