package imagedb

type TagOperation int

const (
	Any TagOperation = 0
	All TagOperation = 1
)

type TagSearch struct {
	TagValues []string // Should always have at least one
	Invert    bool     // (e.g. do we want to see these, or not see these?)
}

type ImageFilter struct {
	Name   string
	Md5Sum *string

	// RequireTags is a map of required tags on the image
	// If you pass a key with a nil value, it will just
	// require the key to exist on the item. If you pass
	// a search, the tag must exist and pass the search.
	RequireTags map[string]*TagSearch

	// RequireTagOperation dictates, if RequireTags is set,
	// how multiple tags will be treated.
	RequireTagOperation TagOperation

	// ExcludeTags is a map of tags on the image that will
	// exclude it from the result. If you pass a key with a
	// nil value, any image with that key in its tags will
	// be excluded. If a search is passed, the image will
	// be excluded if it has the key and passes the search.
	ExcludeTags map[string]*TagSearch

	// ExcludeTagOperation dictates, if ExcludeTags is set,
	// how multiple tags will be treated.
	ExcludeTagOperation TagOperation
}
