package filedb

import "go.mongodb.org/mongo-driver/bson/primitive"

type TagOperation int

const (
	Any TagOperation = 0
	All TagOperation = 1
)

type TagSearch struct {
	TagValues []string // Should always have at least one
	Invert    bool     // (e.g. do we want to see these, or not see these?)
}

type FileFilter struct {
	ID     primitive.ObjectID
	Name   string
	Md5Sum *string

	// RequireTags is a list of required tags that must be
	// present on the file to return it in the search.
	RequireTags []int64

	// RequireTagOperation dictates, if RequireTags is set,
	// how multiple tags will be treated.
	RequireTagOperation TagOperation

	// ExcludeTags is a list of tags that will exclude a given
	// file from the search.
	ExcludeTags []int64

	// ExcludeTagOperation dictates, if ExcludeTags is set,
	// how multiple tags will be treated.
	ExcludeTagOperation TagOperation

	// Tagged dictates, if set, whether to return only
	// files that have been tagged or have not been tagged.
	// Note that this is not tied to the actual value of tags, it
	// is more intended for "A human has looked at this and verified"
	Tagged *bool

	// MissingContent can be used to filter for files that only have
	// metadata but no content (indicating a failed or in progress upload)
	// Defaults to false
	MissingContent bool

	// Limit is the number of results to return from a query
	Limit int64

	// Continue is the ObjectID of the last file on the previous page.
	// In other words, it returns all files with an ID greater than this
	Continue primitive.ObjectID
}
